package vm

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/ethdb/memorydb"
	"github.com/ethereum/go-ethereum/light"
	"github.com/ethereum/go-ethereum/rlp"
	"github.com/ethereum/go-ethereum/trie"
	"gopkg.in/urfave/cli.v1"

	"github.com/mapprotocol/atlas/chains"
	"github.com/mapprotocol/atlas/chains/chainsdb"
	"github.com/mapprotocol/atlas/chains/txverify"
	"github.com/mapprotocol/atlas/core/rawdb"
	"github.com/mapprotocol/atlas/params"
)

var ReceiptsJSON = `[
  {
    "blockHash": "0xe02bf0c849a67807d9465398c829938c560af09617eecaff598ba820ae0c1729",
    "blockNumber": "0x111",
    "contractAddress": null,
    "cumulativeGasUsed": "0xbfdf",
    "from": "0x1aec262a9429eb9167ac4033aaf8b4239c2743fe",
    "gasUsed": "0xbfdf",
    "logs": [
      {
        "address": "0xd6199276959b95a68c1ee30e8569f5fe060903a6",
        "topics": [
          "0x155e433be3576195943c515e1096620bc754e11b3a4b60fda7c4628caf373635",
          "0x000000000000000000000000000068656164657273746f726541646472657373",
          "0x0000000000000000000000001aec262a9429eb9167ac4033aaf8b4239c2743fe",
          "0x000000000000000000000000970e05ffbb2c4a3b80082e82b24f48a29a9c7651"
        ],
        "data": "0x0000000000000000000000000000000000000000000000000000000000000001000000000000000000000000000000000000000000000000000000000000024c000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000d3",
        "blockNumber": "0x111",
        "transactionHash": "0x58e102c383f926992093192bdfb6c6d1197013fd0470475dca6b4c3749484755",
        "transactionIndex": "0x0",
        "blockHash": "0xe02bf0c849a67807d9465398c829938c560af09617eecaff598ba820ae0c1729",
        "logIndex": "0x0",
        "removed": false
      }
    ],
    "logsBloom": "0x00000000000000000000000000000000000000000040000800000000000000000000000000000000000000000000000400000000008000000000000000000000000000000000000000000000000000000000000000000000000000000200200000000000000000021000000000000000000000000080000000000000000004000000000000040000000000000000000000002000000000000000000000000000000000000000000000000000000000000100000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000008000000000001000000000000000",
    "status": "0x1",
    "to": "0xd6199276959b95a68c1ee30e8569f5fe060903a6",
    "transactionHash": "0x58e102c383f926992093192bdfb6c6d1197013fd0470475dca6b4c3749484755",
    "transactionIndex": "0x0",
    "type": "0x0"
  }
]`

var (
	fromAddr  = common.HexToAddress("0x1aec262a9429eb9167ac4033aaf8b4239c2743fe")
	toAddr    = common.HexToAddress("0x970e05ffbb2c4a3b80082e82b24f48a29a9c7651")
	SendValue = big.NewInt(588)
)

type TxParams struct {
	From  []byte
	To    []byte
	Value *big.Int
}

type TxProve struct {
	Tx               *TxParams
	Receipt          *types.Receipt
	Prove            light.NodeList
	BlockNumber      uint64
	TransactionIndex uint
}

func dialConn() *ethclient.Client {
	conn, err := ethclient.Dial("http://192.168.10.215:8545")
	if err != nil {
		log.Fatalf("Failed to connect to the eth: %v", err)
	}
	return conn
}

func getTransactionsHashByBlockNumber(conn *ethclient.Client, number *big.Int) []common.Hash {
	block, err := conn.BlockByNumber(context.Background(), number)
	if err != nil {
		panic(err)
	}
	if block == nil {
		panic("failed to connect to the eth node, please check the network")
	}

	txs := make([]common.Hash, 0, len(block.Transactions()))
	for _, tx := range block.Transactions() {
		txs = append(txs, tx.Hash())
	}
	return txs
}

func getReceiptsByTxsHash(conn *ethclient.Client, txsHash []common.Hash) []*types.Receipt {
	//rs := make([]*types.Receipt, 0, len(txsHash))
	//for _, h := range txsHash {
	//	r, err := conn.TransactionReceipt(context.Background(), h)
	//	if err != nil {
	//		panic(err)
	//	}
	//	if r == nil {
	//		panic("failed to connect to the eth node, please check the network")
	//	}
	//	rs = append(rs, r)
	//}
	//return rs

	return GetReceiptsFromJSON()
}

func GetReceiptsFromJSON() []*types.Receipt {
	var rs []*types.Receipt
	if err := json.Unmarshal([]byte(ReceiptsJSON), &rs); err != nil {
		panic(err)
	}
	return rs
}

func getTxProve() []byte {
	var (
		blockNumber           = big.NewInt(273)
		transactionIndex uint = 0
	)

	// 调用以太坊接口获取 receipts
	//conn := dialConn()
	//txsHash := getTransactionsHashByBlockNumber(conn, blockNumber)
	receipts := getReceiptsByTxsHash(nil, nil)

	// 根据 receipts 生成 trie
	tr, err := trie.New(common.Hash{}, trie.NewDatabase(memorydb.New()))
	if err != nil {
		panic(err)
	}
	for i, r := range receipts {
		key, err := rlp.EncodeToBytes(uint(i))
		if err != nil {
			panic(err)
		}
		value, err := rlp.EncodeToBytes(r)
		if err != nil {
			panic(err)
		}

		tr.Update(key, value)
	}

	proof := light.NewNodeSet()
	key, err := rlp.EncodeToBytes(transactionIndex)
	if err != nil {
		panic(err)
	}
	if err = tr.Prove(key, 0, proof); err != nil {
		panic(err)
	}

	txProve := TxProve{
		Tx: &TxParams{
			From:  fromAddr.Bytes(),
			To:    toAddr.Bytes(),
			Value: SendValue,
		},
		Receipt:          receipts[transactionIndex],
		Prove:            proof.NodeList(),
		BlockNumber:      blockNumber.Uint64(),
		TransactionIndex: transactionIndex,
	}

	input, err := rlp.EncodeToBytes(txProve)
	if err != nil {
		panic(err)
	}
	return input
}

func TestReceiptsRootAndProof(t *testing.T) {
	var (
		srcChain = big.NewInt(1)
		dstChain = big.NewInt(211)
		router   = common.HexToAddress("0xd6199276959b95a68c1ee30e8569f5fe060903a6")
	)

	group, err := chains.ChainType2ChainGroup(rawdb.ChainType(srcChain.Uint64()))
	if err != nil {
		t.Fatal(err)
	}

	set := flag.NewFlagSet("test", 0)
	chainsdb.NewStoreDb(cli.NewContext(nil, set, nil), 10, 2)

	v, err := txverify.Factory(group)
	if err != nil {
		t.Fatal(err)
	}
	if err := v.Verify(router, srcChain, dstChain, getTxProve()); err != nil {
		t.Fatal(err)
	}
}

func TestAddr(t *testing.T) {
	fmt.Println("============================== addr: ", params.TxVerifyAddress)

}
