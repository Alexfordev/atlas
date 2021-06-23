// Copyright 2018 The go-ethereum Authors
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

// Package rawdb contains a collection of low level database accessors.
package rawdb

import (
	"encoding/binary"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/metrics"
)

// The fields below define the low level database schema prefixing.
var (
	// headHeaderKey tracks the latest know header's hash.
	headHeaderKey = []byte("LastHeader")

	// headBlockKey tracks the latest know full block's hash.
	headBlockKey = []byte("LastBlock")

	headRewardKey = []byte("LastReward")

	lastBlockKey = []byte("LastBlockIndex")

	// headFastBlockKey tracks the latest known incomplete block's hash duirng fast sync.
	headFastBlockKey = []byte("LastFast")

	// fastTrieProgressKey tracks the number of trie entries imported during fast sync.
	fastTrieProgressKey = []byte("TrieSync")

	// stateGcBodyReceiptKey tracks the number of body and receipt entries delete during state sync.
	stateGcBodyReceiptKey = []byte("LastState")

	// Data item prefixes (use single byte to avoid mixing data types, avoid `i`, used for indexes).
	headerPrefix       = []byte("h") // headerPrefix + Mark + num (uint64 big endian) + hash -> header
	headerTDSuffix     = []byte("t") // headerPrefix + Mark + num (uint64 big endian) + hash + headerTDSuffix -> td
	headerHashSuffix   = []byte("n") // headerPrefix + Mark + num (uint64 big endian) + headerHashSuffix -> hash
	headerNumberPrefix = []byte("H") // headerNumberPrefix + hash -> num (uint64 big endian)

	blockBodyPrefix     = []byte("b") // blockBodyPrefix + Mark + num (uint64 big endian) + hash -> block body
	blockReceiptsPrefix = []byte("r") // blockReceiptsPrefix + Mark + num (uint64 big endian) + hash -> block receipts

	txLookupPrefix  = []byte("l") // txLookupPrefix + hash -> transaction/receipt lookup metadata
	bloomBitsPrefix = []byte("B") // bloomBitsPrefix + bit (uint16 big endian) + section (uint64 big endian) + hash -> bloom bits

	preimagePrefix = []byte("secure-key-")        // preimagePrefix + hash -> preimage
	configPrefix   = []byte("atlaschain-config-") // config prefix for the db

	// Chain index prefixes (use `i` + single byte to avoid mixing data types).
	BloomBitsIndexPrefix = []byte("iB") // BloomBitsIndexPrefix is the data table of a chain indexer to track its progress

	preimageCounter    = metrics.NewRegisteredCounter("db/preimage/total", nil)
	preimageHitCounter = metrics.NewRegisteredCounter("db/preimage/hits", nil)
)

//--------------- mark ---------
type Mark uint64

func (m Mark) toByte() []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, uint64(m))
	return b
}
func (m Mark) len() int {
	return 8
}

// change the visit Number by append chain mark   mark + number
func (m Mark) markNumber(n uint64) []byte {
	return append(m.toByte(), encodeBlockNumber(n)...)
}

//change the visit key by append chain mark      key  + mark
func (m Mark) markKey(b []byte) []byte {
	return append(b, m.toByte()...)
}

//  key  + mark + number
func (m Mark) markKeyNum(b []byte, n uint64) []byte {
	return append(append(b, m.toByte()...), encodeBlockNumber(n)...)
}

//--------------- mark ---------
// encodeBlockNumber encodes a block number as big endian uint64
func encodeBlockNumber(number uint64) []byte {
	enc := make([]byte, 8)
	binary.BigEndian.PutUint64(enc, number)
	return enc
}

// headerKey = headerPrefix + num (uint64 big endian) + hash
func headerKey(m Mark, number uint64, hash common.Hash) []byte {
	return append(append(headerPrefix, m.markNumber(number)...), hash.Bytes()...)

}

// headerTDKey = headerPrefix + num (uint64 big endian) + hash + headerTDSuffix
func headerTDKey(m Mark, number uint64, hash common.Hash) []byte {
	return append(headerKey(m, number, hash), headerTDSuffix...)
}

// headerHashKey = headerPrefix + num (uint64 big endian) + headerHashSuffix
func headerHashKey(m Mark, number uint64) []byte {
	return append(append(headerPrefix, m.markNumber(number)...), headerHashSuffix...)
}

// headerNumberKey = headerNumberPrefix + hash
func headerNumberKey(m Mark, hash common.Hash) []byte {
	return append(m.markKey(headerNumberPrefix), hash.Bytes()...)
}

// blockBodyKey = blockBodyPrefix + num (uint64 big endian) + hash
func blockBodyKey(m Mark, number uint64, hash common.Hash) []byte {
	return append(append(blockBodyPrefix, m.markNumber(number)...), hash.Bytes()...)
}

// blockReceiptsKey = blockReceiptsPrefix + num (uint64 big endian) + hash
func blockReceiptsKey(m Mark, number uint64, hash common.Hash) []byte {
	return append(append(blockReceiptsPrefix, m.markNumber(number)...), hash.Bytes()...)
}
