package interfaces

import (
	"github.com/Alexfordev/atlas/chains"
	"github.com/Alexfordev/atlas/chains/ethereum"
	"github.com/Alexfordev/atlas/core/types"
)

type IValidate interface {
	ValidateHeaderChain(db types.StateDB, headers []byte, chainType chains.ChainType) (int, error)
}

func ValidateFactory(group chains.ChainGroup) (IValidate, error) {
	switch group {
	case chains.ChainGroupETH:
		return new(ethereum.Validate), nil
	}
	return nil, chains.ErrNotSupportChain
}

type Validate interface {
}
