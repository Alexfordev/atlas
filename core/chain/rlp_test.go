// Copyright 2019 The go-ethereum Authors
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

package chain

import (
	"fmt"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/rlp"
	"github.com/Alexfordev/atlas/consensus/consensustest"
	"github.com/Alexfordev/atlas/core/rawdb"
	"github.com/Alexfordev/atlas/core/types"
	"github.com/Alexfordev/atlas/params"
)

func getBlock(transactions int, uncles int, dataSize int) *types.Block {
	var (
		aa = common.HexToAddress("0x000000000000000000000000000000000000aaaa")
		// Generate a canonical chain to act as the main dataset
		engine = consensustest.NewFaker()
		db     = rawdb.NewMemoryDatabase()
		// A sender who makes transactions, has some funds
		key, _  = crypto.HexToECDSA("b71c71a67e1177ad4e901695e1b4b9ee17ae16c6668d313eac2f96dbcda3f291")
		address = crypto.PubkeyToAddress(key.PublicKey)
		funds   = big.NewInt(1000000000)
		gspec   = &Genesis{
			Config: params.TestChainConfig,
			Alloc:  GenesisAlloc{address: {Balance: funds}},
		}
		genesis = gspec.MustCommit(db)
	)

	// We need to generate as many blocks +1 as uncles
	blocks, _ := GenerateChain(params.TestChainConfig, genesis, engine, db, uncles+1,
		func(n int, b *BlockGen) {
			if n == uncles {
				// Add transactions and stuff on the last block
				for i := 0; i < transactions; i++ {
					tx, _ := types.SignTx(types.NewTransaction(uint64(i), aa,
						big.NewInt(0), 50000, big.NewInt(1), make([]byte, dataSize)), types.HomesteadSigner{}, key)
					b.AddTx(tx)
				}
				for i := 0; i < uncles; i++ {
					b.AddUncle(&types.Header{ParentHash: b.PrevBlock(n - 1 - i).Hash(), Number: big.NewInt(int64(n - i))})
				}
			}
		})
	block := blocks[len(blocks)-1]
	return block
}

// TestRlpIterator tests that individual transactions can be picked out
// from blocks without full unmarshalling/marshalling
func TestRlpIterator(t *testing.T) {
	for _, tt := range []struct {
		txs      int
		uncles   int
		datasize int
	}{
		{0, 0, 0},
		{0, 2, 0},
		{10, 0, 0},
		{10, 2, 0},
		{10, 2, 50},
	} {
		testRlpIterator(t, tt.txs, tt.uncles, tt.datasize)
	}
}

func testRlpIterator(t *testing.T, txs, uncles, datasize int) {
	desc := fmt.Sprintf("%d txs [%d datasize] and %d uncles", txs, datasize, uncles)
	bodyRlp, _ := rlp.EncodeToBytes(getBlock(txs, uncles, datasize).Body())
	it, err := rlp.NewListIterator(bodyRlp)
	if err != nil {
		t.Fatal(err)
	}
	// Check that txs exist
	if !it.Next() {
		t.Fatal("expected two elems, got zero")
	}
	txdata := it.Value()
	// Check that uncles exist
	if !it.Next() {
		t.Fatal("expected two elems, got one")
	}
	// No more after that
	if it.Next() {
		t.Fatal("expected only two elems, got more")
	}
	txIt, err := rlp.NewListIterator(txdata)
	if err != nil {
		t.Fatal(err)
	}
	var gotHashes []common.Hash
	var expHashes []common.Hash
	for txIt.Next() {
		gotHashes = append(gotHashes, crypto.Keccak256Hash(txIt.Value()))
	}

	var expBody types.Body
	err = rlp.DecodeBytes(bodyRlp, &expBody)
	if err != nil {
		t.Fatal(err)
	}
	for _, tx := range expBody.Transactions {
		expHashes = append(expHashes, tx.Hash())
	}
	if gotLen, expLen := len(gotHashes), len(expHashes); gotLen != expLen {
		t.Fatalf("testcase %v: length wrong, got %d exp %d", desc, gotLen, expLen)
	}
	// also sanity check against input
	if gotLen := len(gotHashes); gotLen != txs {
		t.Fatalf("testcase %v: length wrong, got %d exp %d", desc, gotLen, txs)
	}
	for i, got := range gotHashes {
		if exp := expHashes[i]; got != exp {
			t.Errorf("testcase %v: hash wrong, got %x, exp %x", desc, got, exp)
		}
	}
}
