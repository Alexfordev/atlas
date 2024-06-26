package gold_token

import (
	"math/big"

	"github.com/Alexfordev/atlas/contracts"
	"github.com/Alexfordev/atlas/contracts/abis"
	"github.com/Alexfordev/atlas/core/vm"
	"github.com/Alexfordev/atlas/params"
	"github.com/ethereum/go-ethereum/common"
)

var (
	totalSupplyMethod    = contracts.NewRegisteredContractMethod(params.GoldTokenRegistryId, abis.GoldToken, "totalSupply", params.MaxGasForTotalSupply)
	increaseSupplyMethod = contracts.NewRegisteredContractMethod(params.GoldTokenRegistryId, abis.GoldToken, "increaseSupply", params.MaxGasForIncreaseSupply)
	mintMethod           = contracts.NewRegisteredContractMethod(params.GoldTokenRegistryId, abis.GoldToken, "mint", params.MaxGasForMintGas)
)

func GetTotalSupply(vmRunner vm.EVMRunner) (*big.Int, error) {
	var totalSupply *big.Int
	err := totalSupplyMethod.Query(vmRunner, &totalSupply)
	return totalSupply, err
}

func IncreaseSupply(vmRunner vm.EVMRunner, value *big.Int) error {
	err := increaseSupplyMethod.Execute(vmRunner, nil, common.Big0, value)
	return err
}

func Mint(vmRunner vm.EVMRunner, beneficiary common.Address, value *big.Int) error {
	if value.Cmp(new(big.Int)) <= 0 {
		return nil
	}

	err := mintMethod.Execute(vmRunner, nil, common.Big0, beneficiary, value)
	return err
}
