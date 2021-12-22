package tools

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/rlp"
	"golang.org/x/crypto/sha3"
)

func RlpHash(x interface{}) (h common.Hash) {
	hw := sha3.NewLegacyKeccak256()
	if e := rlp.Encode(hw, x); e != nil {
		log.Warn("RlpHash", "error", e.Error())
	}
	hw.Sum(h[:0])
	return h
}
