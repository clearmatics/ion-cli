package backend

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/core/types"
)

type BlockMap map[string]BlockInterface

type BlockInterface interface {
	RlpEncode() (err error)
	GetByNumber(rpcURL string, number string) (err error)
	GetByHash(rpcURL string, hash string) (err error)
}

type Transaction struct {
	Tx    *types.Transaction `json:"tx"`
	Proof string             `json:"proof"`
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
