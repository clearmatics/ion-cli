package backend

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/core/types"
)

// TODO might need to specify on which consensus
type BlockHeader struct {
	Header     *types.Header `json:"header"`
	RlpEncoded string        `json:"rlp_encoded"`
}

type Transaction struct {
	Tx    *types.Transaction `json:"tx"`
	Proof string             `json:"proof"`
}

type Session struct {
	Timestamp int `json:"timestamp"`
	// lenght of the session

	// network
	Rpc         string `json:"rpc"`
	Active      bool   `json:"active"`
	AccountName string `json:"account"`

	// fields that have to be cached for subsequent calls
	Block       BlockHeader `json:"block"`
	Transaction Transaction `json:"transaction"`
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
