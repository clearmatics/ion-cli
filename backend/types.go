package backend

import (
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
)

// TODO might need to specify on which consensus
// TODO the rlp encoded is supposed to be used by ION to submit blocks. how do we identify which block
type BlockHeader struct {
	Header *types.Header `json:"header"`
	RlpEncoded string`json:"rlp_encoded"`
}

type Transaction struct {
	Tx *types.Transaction `json:"tx"`
	Proof string `json:"proof"`
}

type Session struct {
	Timestamp int `json:"timestamp"`
	// lenght of the session

	// network
	Rpc string `json:"rpc"`
	Active bool `json:"active"`
	AccountName string `json:"account"`

	// fields that have to be cached for subsequent calls
	Block BlockHeader `json:"block"`
	Transaction Transaction `json:"transaction"`
}

type EthClient struct {
	client    *ethclient.Client
	rpcClient *rpc.Client
	url       string
}

// an unlocked wallet object
type Wallet struct {
	Auth *bind.TransactOpts `json:"auth"`
	Key  *keystore.Key `json:"key"`
	Name string `json:"name"`
}

// an account info as stored in the configs
type AccountInfo struct {
	Name string `json:"name"`
	Keyfile string `json:"keyfile"`
	Password string `json:"password"`
}