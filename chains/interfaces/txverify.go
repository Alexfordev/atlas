package interfaces

import (
	"github.com/ethereum/go-ethereum/common"

	"github.com/Alexfordev/atlas/chains"
	"github.com/Alexfordev/atlas/chains/ethereum"
	"github.com/Alexfordev/atlas/core/types"
)

type IVerify interface {
	Verify(db types.StateDB, router common.Address, txProveBytes []byte) (logs []byte, err error)
}

func VerifyFactory(group chains.ChainGroup) (IVerify, error) {
	switch group {
	case chains.ChainGroupETH:
		return new(ethereum.Verify), nil
	}
	return nil, chains.ErrNotSupportChain
}
