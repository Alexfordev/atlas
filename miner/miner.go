// Copyright 2014 The go-ethereum Authors
// This file is part of the go-ethereum library.
//
// The go-ethereum library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-ethereum library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-ethereum library. If not, see <http://www.gnu.org/licenses/>.

// Package miner implements Ethereum block creation and mining.
package miner

import (
	"fmt"
	"math/big"
	"time"

	"github.com/Alexfordev/atlas/atlas/downloader"
	"github.com/Alexfordev/atlas/consensus"
	"github.com/Alexfordev/atlas/contracts/random"
	"github.com/Alexfordev/atlas/core/chain"
	"github.com/Alexfordev/atlas/core/rawdb"
	"github.com/Alexfordev/atlas/core/state"
	"github.com/Alexfordev/atlas/core/types"
	params2 "github.com/Alexfordev/atlas/params"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/ethdb"
	"github.com/ethereum/go-ethereum/event"
	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/params"
)

// Backend wraps all methods required for mining.
type Backend interface {
	BlockChain() *chain.BlockChain
	TxPool() *chain.TxPool
}

// Config is the configuration parameters of mining.
type Config struct {
	Etherbase common.Address `toml:",omitempty"` // Public address for block mining rewards (default = first account)
	// Notify     []string       `toml:",omitempty"` // HTTP URL list to be notified of new work packages (only useful in ethash).
	// NotifyFull bool           `toml:",omitempty"` // Notify with pending block headers instead of work packages
	ExtraData hexutil.Bytes `toml:",omitempty"` // Block extra data set by the miner
	GasFloor  uint64        // Target gas floor for mined blocks.
	GasCeil   uint64        // Target gas ceiling for mined blocks.
	GasPrice  *big.Int      // Minimum gas price for mining a transaction
	Recommit  time.Duration // The time interval for miner to re-create mining work.
	// Noverify   bool           // Disable remote mining solution verification(only useful in ethash).
}

// Miner creates blocks and searches for proof-of-work values.
type Miner struct {
	mux      *event.TypeMux
	worker   *worker
	coinbase common.Address
	eth      Backend
	engine   consensus.Engine
	db       ethdb.Database // Needed for randomness

	exitCh  chan struct{}
	startCh chan common.Address
	stopCh  chan struct{}
}

func New(eth Backend, config *Config, chainConfig *params2.ChainConfig, mux *event.TypeMux, engine consensus.Engine, isLocalBlock func(block *types.Block) bool, db ethdb.Database) *Miner {
	miner := &Miner{
		eth:     eth,
		mux:     mux,
		engine:  engine,
		exitCh:  make(chan struct{}),
		startCh: make(chan common.Address),
		stopCh:  make(chan struct{}),
		worker:  newWorker(config, chainConfig, engine, eth, mux, isLocalBlock, true, db),
		db:      db,
	}
	go miner.update()

	return miner
}

func (miner *Miner) recoverRandomness() {
	// If this is using the istanbul consensus engine, then we need to check
	// for the randomness cache for the randomness beacon protocol
	_, isIstanbul := miner.engine.(consensus.Istanbul)
	if isIstanbul {
		// getCurrentBlockAndState
		currentBlock := miner.eth.BlockChain().CurrentBlock()
		currentHeader := currentBlock.Header()
		currentState, err := miner.eth.BlockChain().StateAt(currentBlock.Root())
		if err != nil {
			log.Error("Error in retrieving state", "block hash", currentHeader.Hash(), "error", err)
			return
		}

		if currentHeader.Number.Uint64() > 0 {
			vmRunner := miner.eth.BlockChain().NewEVMRunner(currentHeader, currentState)
			// Check to see if we already have the commitment cache
			lastCommitment, err := random.GetLastCommitment(vmRunner, miner.coinbase)
			if err != nil {
				log.Error("Error in retrieving last commitment", "error", err)
				return
			}

			// If there is a non empty last commitment and if we don't have that commitment's
			// cache entry, then we need to recover it.
			if (lastCommitment != common.Hash{}) && (rawdb.ReadRandomCommitmentCache(miner.db, lastCommitment) == common.Hash{}) {
				err := miner.eth.BlockChain().RecoverRandomnessCache(lastCommitment, currentBlock.Hash())
				if err != nil {
					log.Error("Error in recovering randomness cache", "error", err)
					return
				}
			}
		}
	}

}

// update keeps track of the downloader events. Please be aware that this is a one shot type of update loop.
// It's entered once and as soon as `Done` or `Failed` has been broadcasted the events are unregistered and
// the loop is exited. This to prevent a major security vuln where external parties can DOS you with blocks
// and halt your mining operation for as long as the DOS continues.
func (miner *Miner) update() {
	events := miner.mux.Subscribe(downloader.StartEvent{}, downloader.DoneEvent{}, downloader.FailedEvent{})
	defer func() {
		if !events.Closed() {
			events.Unsubscribe()
		}
	}()

	shouldStart := false
	canStart := true
	dlEventCh := events.Chan()
	for {
		select {
		case ev := <-dlEventCh:
			if ev == nil {
				// Unsubscription done, stop listening
				dlEventCh = nil
				continue
			}
			switch ev.Data.(type) {
			case downloader.StartEvent:
				wasMining := miner.Mining()
				miner.worker.stop()
				canStart = false
				if wasMining {
					// Resume mining after sync was finished
					shouldStart = true
					log.Info("Mining aborted due to sync")
				}
			case downloader.FailedEvent:
				canStart = true
				if shouldStart {
					miner.SetValidator(miner.coinbase)
					miner.worker.start()
				}
			case downloader.DoneEvent:
				miner.recoverRandomness()
				canStart = true
				if shouldStart {
					miner.SetValidator(miner.coinbase)
					miner.worker.start()
				}
				// Stop reacting to downloader events
				events.Unsubscribe()
			}
		case addr := <-miner.startCh:
			miner.SetValidator(addr)
			if canStart {
				miner.worker.start()
			}
			shouldStart = true
		case <-miner.stopCh:
			shouldStart = false
			miner.worker.stop()
		case <-miner.exitCh:
			miner.worker.close()
			return
		}
	}
}

func (miner *Miner) Start(validator common.Address, txFeeRecipient common.Address) {
	miner.SetValidator(validator)
	miner.SetTxFeeRecipient(txFeeRecipient)
	miner.startCh <- validator
}

func (miner *Miner) Stop() {
	miner.stopCh <- struct{}{}
}

func (miner *Miner) Close() {
	close(miner.exitCh)
}

func (miner *Miner) Mining() bool {
	return miner.worker.isRunning()
}

// func (miner *Miner) Hashrate() uint64 {
//	if pow, ok := miner.engine.(consensus.PoW); ok {
//		return uint64(pow.Hashrate())
//	}
//	return 0
// }

func (miner *Miner) SetExtra(extra []byte) error {
	if uint64(len(extra)) > params.MaximumExtraDataSize {
		return fmt.Errorf("extra exceeds max length. %d > %v", len(extra), params.MaximumExtraDataSize)
	}
	miner.worker.setExtra(extra)
	return nil
}

// SetRecommitInterval sets the interval for sealing work resubmitting.
// func (miner *Miner) SetRecommitInterval(interval time.Duration) {
//	miner.worker.setRecommitInterval(interval)
// }

// Pending returns the currently pending block and associated state.
func (miner *Miner) Pending() (*types.Block, *state.StateDB) {
	return miner.worker.pending()
}

// PendingBlock returns the currently pending block.
//
// Note, to access both the pending block and the pending state
// simultaneously, please use Pending(), as the pending state can
// change between multiple method calls
func (miner *Miner) PendingBlock() *types.Block {
	return miner.worker.pendingBlock()
}

func (miner *Miner) SetValidator(addr common.Address) {
	miner.coinbase = addr
	miner.worker.setValidator(addr)
}

// EnablePreseal turns on the preseal mining feature. It's enabled by default.
// Note this function shouldn't be exposed to API, it's unnecessary for users
// (miners) to actually know the underlying detail. It's only for outside project
// which uses this library.
// func (miner *Miner) EnablePreseal() {
//	miner.worker.enablePreseal()
// }

// DisablePreseal turns off the preseal mining feature. It's necessary for some
// fake consensus engine which can seal blocks instantaneously.
// Note this function shouldn't be exposed to API, it's unnecessary for users
// (miners) to actually know the underlying detail. It's only for outside project
// which uses this library.
// func (miner *Miner) DisablePreseal() {
//	miner.worker.disablePreseal()
// }

// SetTxFeeRecipient sets the address where the miner and worker will receive fees
func (miner *Miner) SetTxFeeRecipient(addr common.Address) {
	miner.worker.setTxFeeRecipient(addr)
}

// SubscribePendingLogs starts delivering logs from pending transactions
// to the given channel.
func (miner *Miner) SubscribePendingLogs(ch chan<- []*types.Log) event.Subscription {
	return miner.worker.pendingLogsFeed.Subscribe(ch)
}
