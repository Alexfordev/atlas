// Copyright 2015 The go-ethereum Authors
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

package runtime

import (
	"github.com/Alexfordev/atlas/core/chain"
	"github.com/Alexfordev/atlas/core/vm"
	"github.com/Alexfordev/atlas/core/vm/vmcontext"
)

func NewEnv(cfg *Config) *vm.EVM {
	txContext := vm.TxContext{
		Origin:   cfg.Origin,
		GasPrice: cfg.GasPrice,
	}
	blockContext := vm.BlockContext{
		CanTransfer:          chain.CanTransfer,
		Transfer:             vmcontext.TobinTransfer,
		GetHash:              cfg.GetHashFn,
		Coinbase:             cfg.Coinbase,
		BlockNumber:          cfg.BlockNumber,
		Time:                 cfg.Time,
		Difficulty:           cfg.Difficulty,
		GasLimit:             cfg.GasLimit,
		BaseFee:              cfg.BaseFee,
		GetRegisteredAddress: vmcontext.GetRegisteredAddress,
	}
	if cfg.ChainConfig.Istanbul != nil {
		blockContext.EpochSize = cfg.ChainConfig.Istanbul.Epoch
	}
	return vm.NewEVM(blockContext, txContext, cfg.State, cfg.ChainConfig, cfg.EVMConfig)
}
