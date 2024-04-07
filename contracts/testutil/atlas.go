package testutil

import (
	"github.com/Alexfordev/atlas/params"
	"github.com/ethereum/go-ethereum/common"
)

type AtlasMock struct {
	Runner               *MockEVMRunner
	Registry             *RegistryMock
	BlockchainParameters *BlockchainParametersMock
}

func NewAtlasMock() AtlasMock {
	atlas := AtlasMock{
		Runner:               NewMockEVMRunner(),
		Registry:             NewRegistryMock(),
		BlockchainParameters: NewBlockchainParametersMock(),
	}

	atlas.Runner.RegisterContract(params.RegistrySmartContractAddress, atlas.Registry)

	atlas.Registry.AddContract(params.BlockchainParametersRegistryId, common.HexToAddress("0x01"))
	atlas.Runner.RegisterContract(common.HexToAddress("0x01"), atlas.BlockchainParameters)

	return atlas
}
