package abis

import (
	"fmt"
	"strings"

	"github.com/Alexfordev/atlas/accounts/abi"
	"github.com/Alexfordev/atlas/params"
	"github.com/ethereum/go-ethereum/common"
)

var (
	Registry             *abi.ABI = mustParseAbi("Registry", RegistryStr)
	BlockchainParameters *abi.ABI = mustParseAbi("BlockchainParameters", BlockchainParametersStr)
	ERC20                *abi.ABI = mustParseAbi("ERC20", ERC20Str)
	FeeCurrency          *abi.ABI = mustParseAbi("FeeCurrency", FeeCurrencyStr)
	Elections            *abi.ABI = mustParseAbi("Elections", ElectionsStr)
	EpochRewards         *abi.ABI = mustParseAbi("EpochRewards", EpochRewardsStr)
	GasPriceMinimum      *abi.ABI = mustParseAbi("GasPriceMinimum", GasPriceMinimumStr)
	GoldToken            *abi.ABI = mustParseAbi("GoldToken", GoldTokenStr)
	Random               *abi.ABI = mustParseAbi("Random", RandomStr)
	Validators           *abi.ABI = mustParseAbi("Validators", ValidatorsStr)
	Accounts             *abi.ABI = mustParseAbi("Accounts", AccountsStr)
)

func mustParseAbi(name, abiStr string) *abi.ABI {
	parsedAbi, err := abi.JSON(strings.NewReader(abiStr))
	if err != nil {
		panic(fmt.Sprintf("Error reading ABI %s err=%s", name, err))
	}
	return &parsedAbi
}

var byRegistryId = map[common.Hash]*abi.ABI{
	params.BlockchainParametersRegistryId: BlockchainParameters,
	params.FeeCurrencyWhitelistRegistryId: FeeCurrency,
	params.ElectionRegistryId:             Elections,
	params.EpochRewardsRegistryId:         EpochRewards,
	params.GasPriceMinimumRegistryId:      GasPriceMinimum,
	params.GoldTokenRegistryId:            GoldToken,
	params.RandomRegistryId:               Random,
	params.ValidatorsRegistryId:           Validators,
}

func AbiFor(registryId common.Hash) *abi.ABI {
	found, ok := byRegistryId[registryId]
	if !ok {
		return nil
	}
	return found
}
