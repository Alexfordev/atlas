package chains

import (
	"errors"
	"github.com/mapprotocol/atlas/core/rawdb"
)

const (
	ChainTypeMAP rawdb.ChainType = 1000
	ChainTypeETH rawdb.ChainType = 1001
)

const (
	ChainNameMAP = "MAP"
	ChainNameETH = "ETH"
)

var name2type = map[string]rawdb.ChainType{
	ChainNameMAP: ChainTypeMAP,
	ChainNameETH: ChainTypeETH,
}

func ChainNameToChainType(chain string) (rawdb.ChainType, error) {
	chainType, ok := name2type[chain]
	if !ok {
		return 0, errors.New("unsupported chain ")
	}
	return chainType, nil
}
