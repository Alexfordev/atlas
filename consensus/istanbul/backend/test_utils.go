package backend

import (
	"bytes"
	"crypto/ecdsa"
	"errors"
	"fmt"
	"github.com/mapprotocol/atlas/core/chain"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/crypto/ecies"
	"github.com/ethereum/go-ethereum/rlp"
	"github.com/mapprotocol/atlas/accounts"
	"github.com/mapprotocol/atlas/consensus/consensustest"
	"github.com/mapprotocol/atlas/consensus/istanbul"
	"github.com/mapprotocol/atlas/consensus/istanbul/backend/backendtest"
	"github.com/mapprotocol/atlas/consensus/istanbul/validator"
	"github.com/mapprotocol/atlas/core"
	"github.com/mapprotocol/atlas/core/rawdb"
	"github.com/mapprotocol/atlas/core/state"
	"github.com/mapprotocol/atlas/core/types"
	"github.com/mapprotocol/atlas/core/vm"
	blscrypto "github.com/mapprotocol/atlas/helper/bls"
	"github.com/mapprotocol/atlas/params"
)

// in this test, we can set n to 1, and it means we can process Istanbul and commit a
// block by one node. Otherwise, if n is larger than 1, we have to generate
// other fake events to process Istanbul.
func newBlockChain(n int, isFullChain bool) (*chain.BlockChain, *Backend) {
	genesis, nodeKeys := getGenesisAndKeys(n, isFullChain)

	bc, be, _ := newBlockChainWithKeys(false, common.Address{}, false, genesis, nodeKeys[0])
	return bc, be
}

func newBlockChainWithKeys(isProxy bool, proxiedValAddress common.Address, isProxied bool, genesis *chain.Genesis, privateKey *ecdsa.PrivateKey) (*chain.BlockChain, *Backend, *istanbul.Config) {
	memDB := rawdb.NewMemoryDatabase()
	config := *istanbul.DefaultConfig
	config.ReplicaStateDBPath = ""
	config.ValidatorEnodeDBPath = ""
	config.VersionCertificateDBPath = ""
	config.RoundStateDBPath = ""
	config.Proxy = isProxy
	config.ProxiedValidatorAddress = proxiedValAddress
	config.Proxied = isProxied
	config.Validator = !isProxy
	istanbul.ApplyParamsChainConfigToConfig(genesis.Config, &config)

	b, _ := New(&config, memDB).(*Backend)

	var publicKey ecdsa.PublicKey
	if !isProxy {
		publicKey = privateKey.PublicKey
		address := crypto.PubkeyToAddress(publicKey)
		decryptFn := DecryptFn(privateKey)
		signerFn := SignFn(privateKey)
		signerBLSFn := SignBLSFn(privateKey)
		signerHashFn := SignHashFn(privateKey)
		b.Authorize(address, address, &publicKey, decryptFn, signerFn, signerBLSFn, signerHashFn)
	} else {
		proxyNodeKey, _ := crypto.GenerateKey()
		publicKey = proxyNodeKey.PublicKey
	}

	genesis.MustCommit(memDB)

	blockchain, err := chain.NewBlockChain(memDB, nil, genesis.Config, b, vm.Config{}, nil, nil)
	if err != nil {
		panic(err)
	}

	b.SetChain(
		blockchain,
		blockchain.CurrentBlock,
		func(hash common.Hash) (*state.StateDB, error) {
			stateRoot := blockchain.GetHeaderByHash(hash).Root
			return blockchain.StateAt(stateRoot)
		},
	)
	b.SetBroadcaster(&consensustest.MockBroadcaster{})
	b.SetP2PServer(consensustest.NewMockP2PServer(&publicKey))
	b.StartAnnouncing()

	if !isProxy {
		b.SetCallBacks(blockchain.HasBadBlock,
			func(block *types.Block, state *state.StateDB) (types.Receipts, []*types.Log, uint64, error) {
				return blockchain.Processor().Process(block, state, *blockchain.GetVMConfig())
			},
			blockchain.Validator().ValidateState,
			func(block *types.Block, receipts []*types.Receipt, logs []*types.Log, state *state.StateDB) {
				if err := blockchain.WriteBlockWithState(block, receipts, logs, state, true); err != nil {
					panic(fmt.Sprintf("could not InsertPreprocessedBlock: %v", err))
				}
			})
		if isProxied {
			b.StartProxiedValidatorEngine()
		}
		b.StartValidating()
	}

	return blockchain, b, &config
}

func getGenesisAndKeys(n int, isFullChain bool) (*chain.Genesis, []*ecdsa.PrivateKey) {
	// Setup validators
	var nodeKeys = make([]*ecdsa.PrivateKey, n)
	validators := make([]istanbul.ValidatorData, n)
	for i := 0; i < n; i++ {
		var addr common.Address
		if i == 0 {
			nodeKeys[i], _ = generatePrivateKey()
			addr = getAddress()
		} else {
			nodeKeys[i], _ = crypto.GenerateKey()
			addr = crypto.PubkeyToAddress(nodeKeys[i].PublicKey)
		}
		blsPrivateKey, _ := blscrypto.CryptoType().ECDSAToBLS(nodeKeys[i])
		blsPublicKey, _ := blscrypto.CryptoType().PrivateToPublic(blsPrivateKey)
		blsG1PublicKey, _ := blscrypto.CryptoType().PrivateToG1Public(blsPrivateKey)
		validators[i] = istanbul.ValidatorData{
			Address:        addr,
			BLSPublicKey:   blsPublicKey,
			BLSG1PublicKey: blsG1PublicKey,
		}

	}

	// generate genesis block
	genesis := chain.DefaultGenesisBlock()
	genesis.Config = params.IstanbulTestChainConfig
	if !isFullChain {
		genesis.Config.FullHeaderChainAvailable = false
	}
	// force enable Istanbul engine
	genesis.Config.Istanbul = &params.IstanbulConfig{
		Epoch:          10,
		LookbackWindow: 3,
	}

	AppendValidatorsToGenesisBlock(genesis, validators)
	return genesis, nodeKeys
}

func AppendValidatorsToGenesisBlock(genesis *chain.Genesis, validators []istanbul.ValidatorData) {
	if len(genesis.ExtraData) < types.IstanbulExtraVanity {
		genesis.ExtraData = append(genesis.ExtraData, bytes.Repeat([]byte{0x00}, types.IstanbulExtraVanity)...)
	}
	genesis.ExtraData = genesis.ExtraData[:types.IstanbulExtraVanity]

	var addrs []common.Address
	var publicKeys []blscrypto.SerializedPublicKey
	var g1publicKeys []blscrypto.SerializedG1PublicKey

	for i := range validators {
		if (validators[i].BLSPublicKey == blscrypto.SerializedPublicKey{}) {
			panic("BLSPublicKey is nil")
		}
		addrs = append(addrs, validators[i].Address)
		publicKeys = append(publicKeys, validators[i].BLSPublicKey)
		g1publicKeys = append(g1publicKeys, validators[i].BLSG1PublicKey)
	}

	ist := &types.IstanbulExtra{
		AddedValidators:             addrs,
		AddedValidatorsPublicKeys:   publicKeys,
		AddedValidatorsG1PublicKeys: g1publicKeys,
		Seal:                        []byte{},
		AggregatedSeal:              types.IstanbulAggregatedSeal{},
		ParentAggregatedSeal:        types.IstanbulAggregatedSeal{},
	}

	istPayload, err := rlp.EncodeToBytes(&ist)
	if err != nil {
		panic("failed to encode istanbul extra")
	}
	genesis.ExtraData = append(genesis.ExtraData, istPayload...)
}

func makeHeader(parent *types.Block, config *istanbul.Config) *types.Header {
	header := &types.Header{
		ParentHash: parent.Hash(),
		Number:     parent.Number().Add(parent.Number(), common.Big1),
		GasUsed:    0,
		Extra:      parent.Extra(),
		Time:       parent.Time() + config.BlockPeriod,
	}
	return header
}

func makeBlock(keys []*ecdsa.PrivateKey, chain *chain.BlockChain, engine *Backend, parent *types.Block) (*types.Block, error) {
	block := makeBlockWithoutSeal(chain, engine, parent)

	// Set up block subscription
	chainHeadCh := make(chan core.ChainHeadEvent, 10)
	sub := chain.SubscribeChainHeadEvent(chainHeadCh)
	defer sub.Unsubscribe()

	// start seal request (this is non-blocking)
	err := engine.Seal(chain, block)
	if err != nil {
		return nil, err
	}

	// Wait for and then save the mined block.
	select {
	case ev := <-chainHeadCh:
		block = ev.Block
	case <-time.After(6 * time.Second):
		return nil, errors.New("Timed out when making a block")
	}

	// Notify the core engine to stop working on current Seal.
	go engine.istanbulEventMux.Post(istanbul.FinalCommittedEvent{})

	return block, nil
}

func makeBlockWithoutSeal(chain *chain.BlockChain, engine *Backend, parent *types.Block) *types.Block {
	header := makeHeader(parent, engine.config)
	// The worker that calls Prepare is the one filling the Coinbase
	header.Coinbase = engine.wallets().Ecdsa.Address
	engine.Prepare(chain, header)
	time.Sleep(time.Until(time.Unix(int64(header.Time), 0)))

	state, err := chain.StateAt(parent.Root())
	if err != nil {
		fmt.Printf("Error!! %v\n", err)
	}
	engine.Finalize(chain, header, state, nil)

	block, err := engine.FinalizeAndAssemble(chain, header, state, nil, nil, nil)
	if err != nil {
		fmt.Printf("Error!! %v\n", err)
	}

	return block
}

/**
 * SimpleBackend
 * Private key: bb047e5940b6d83354d9432db7c449ac8fca2248008aaa7271369880f9f11cc1
 * Public key: 04a2bfb0f7da9e1b9c0c64e14f87e8fb82eb0144e97c25fe3a977a921041a50976984d18257d2495e7bfd3d4b280220217f429287d25ecdf2b0d7c0f7aae9aa624
 * Address: 0x70524d664ffe731100208a0154e556f9bb679ae6
 */
func getAddress() common.Address {
	return common.HexToAddress("0x70524d664ffe731100208a0154e556f9bb679ae6")
}

func getInvalidAddress() common.Address {
	return common.HexToAddress("0xc63597005f0da07a9ea85b5052a77c3b0261bdca")
}

func generatePrivateKey() (*ecdsa.PrivateKey, error) {
	key := "bb047e5940b6d83354d9432db7c449ac8fca2248008aaa7271369880f9f11cc1"
	return crypto.HexToECDSA(key)
}

func newTestValidatorSet(n int) (istanbul.ValidatorSet, []*ecdsa.PrivateKey) {
	// generate validators
	keys := make(Keys, n)
	validators := make([]istanbul.ValidatorData, n)
	for i := 0; i < n; i++ {
		privateKey, _ := crypto.GenerateKey()
		blsPrivateKey, _ := blscrypto.CryptoType().ECDSAToBLS(privateKey)
		blsPublicKey, _ := blscrypto.CryptoType().PrivateToPublic(blsPrivateKey)
		keys[i] = privateKey
		validators[i] = istanbul.ValidatorData{
			Address:      crypto.PubkeyToAddress(privateKey.PublicKey),
			BLSPublicKey: blsPublicKey,
		}
	}
	vset := validator.NewSet(validators)
	return vset, keys
}

type Keys []*ecdsa.PrivateKey

func (slice Keys) Len() int {
	return len(slice)
}

func (slice Keys) Less(i, j int) bool {
	return strings.Compare(crypto.PubkeyToAddress(slice[i].PublicKey).String(), crypto.PubkeyToAddress(slice[j].PublicKey).String()) < 0
}

func (slice Keys) Swap(i, j int) {
	slice[i], slice[j] = slice[j], slice[i]
}

func DecryptFn(key *ecdsa.PrivateKey) istanbul.DecryptFn {
	if key == nil {
		key, _ = generatePrivateKey()
	}

	return func(_ accounts.Account, c, s1, s2 []byte) ([]byte, error) {
		eciesKey := ecies.ImportECDSA(key)
		return eciesKey.Decrypt(c, s1, s2)
	}
}

func SignFn(key *ecdsa.PrivateKey) istanbul.SignerFn {
	if key == nil {
		key, _ = generatePrivateKey()
	}

	return func(_ accounts.Account, mimeType string, data []byte) ([]byte, error) {
		return crypto.Sign(crypto.Keccak256(data), key)
	}
}

func SignBLSFn(key *ecdsa.PrivateKey) istanbul.BLSSignerFn {
	if key == nil {
		key, _ = generatePrivateKey()
	}

	return func(_ accounts.Account, data []byte, extraData []byte, useComposite, cip22 bool) (blscrypto.SerializedSignature, error) {
		from := crypto.PubkeyToAddress(key.PublicKey)

		keybytes, err := blscrypto.CryptoType().ECDSAToBLS(key)
		if err != nil {
			return blscrypto.SerializedSignature{}, err
		}
		prikey, err := blscrypto.DeserializePrivateKey(keybytes)
		if err != nil {
			return blscrypto.SerializedSignature{}, err
		}

		pkbytes, err := blscrypto.CryptoType().PrivateToPublic(keybytes)
		if err != nil {
			return blscrypto.SerializedSignature{}, err
		}
		pubkey, err := blscrypto.UnmarshalPk(pkbytes[:])
		if err != nil {
			return blscrypto.SerializedSignature{}, err
		}
		signature, err := blscrypto.Sign(prikey, pubkey, from.Bytes())
		if err != nil {
			return blscrypto.SerializedSignature{}, err
		}
		signature2 := blscrypto.SerializedSignature{}
		copy(signature2[:], signature.Marshal())
		return signature2, nil
	}
}

func SignHashFn(key *ecdsa.PrivateKey) istanbul.HashSignerFn {
	if key == nil {
		key, _ = generatePrivateKey()
	}

	return func(_ accounts.Account, data []byte) ([]byte, error) {
		return crypto.Sign(data, key)
	}
}

func newBackend() (b *Backend) {
	_, b = newBlockChain(4, true)

	key, _ := generatePrivateKey()
	address := crypto.PubkeyToAddress(key.PublicKey)
	b.Authorize(address, address, &key.PublicKey, DecryptFn(key), SignFn(key), SignBLSFn(key), SignHashFn(key))
	return
}

type testBackendFactoryImpl struct{}

// TestBackendFactory can be passed to backendtest.InitTestBackendFactory
var TestBackendFactory backendtest.TestBackendFactory = testBackendFactoryImpl{}

// New is part of TestBackendInterface.
func (testBackendFactoryImpl) New(isProxy bool, proxiedValAddress common.Address, isProxied bool, genesisCfg *chain.Genesis, privateKey *ecdsa.PrivateKey) (backendtest.TestBackendInterface, *istanbul.Config) {
	_, be, config := newBlockChainWithKeys(isProxy, proxiedValAddress, isProxied, genesisCfg, privateKey)
	return be, config
}

// GetGenesisAndKeys is part of TestBackendInterface
func (testBackendFactoryImpl) GetGenesisAndKeys(numValidators int, isFullChain bool) (*chain.Genesis, []*ecdsa.PrivateKey) {
	return getGenesisAndKeys(numValidators, isFullChain)
}
