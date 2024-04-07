package mempool

import (
	"io"
	"log"

	"github.com/Alexfordev/atlas/tools/utils/btcapi"
	"github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcd/wire"
)

type MempoolClient struct {
	baseURL string
}

func NewClient(netParams *chaincfg.Params) *MempoolClient {
	baseURL := ""
	if netParams.Net == wire.MainNet {
		baseURL = "https://mempool.space/api"
	} else if netParams.Net == wire.TestNet3 {
		baseURL = "https://mempool.space/testnet/api"
	} else {
		log.Fatal("mempool don't support other netParams")
	}
	return &MempoolClient{
		baseURL: baseURL,
	}
}

func (c *MempoolClient) request(method, subPath string, requestBody io.Reader) ([]byte, error) {
	return btcapi.Request(method, c.baseURL, subPath, requestBody)
}

var _ btcapi.BTCAPIClient = (*MempoolClient)(nil)
