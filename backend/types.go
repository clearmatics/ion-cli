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


type BlockMap map[string]BlockInterface

type BlockInterface interface {
	RlpEncode() (err error)
}

type EthBlockHeader struct {
	Header     *types.Header 	 `json:"header"`
	RlpEncoded string        `json:"rlp_encoded"`
}

type Transaction struct {
	Tx    *types.Transaction `json:"tx"`
	Proof string             `json:"proof"`
}

type EthClient struct {
	client    *ethclient.Client
	rpcClient *rpc.Client
	url       string
}

// an unlocked wallet object
type Wallet struct {
	Auth *bind.TransactOpts `json:"auth"`
	Key  *keystore.Key      `json:"key"`
	Name string             `json:"name"`
}

// an account info as stored in the configs
type AccountInfo struct {
	Name     string `json:"name"`
	Keyfile  string `json:"keyfile"`
	Password string `json:"password"`
}
