package backend

import (
	"encoding/json"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/core/types"
)

type Block struct {
	Type string `json:"type"`
	Header json.RawMessage `json:"header"` // this is the polymorphic bit
	Interface BlockInterface `json:"-"`
}

type Blocks map[string]Block

type BlockInterface interface {
	RlpEncode() (err error)
	GetByNumber(rpcURL string, number string) (err error)
	GetByHash(rpcURL string, hash string) (err error)
	Marshal() (header []byte, err error)
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

// keeps all config of a particular chain
type NetworkInfo struct {
	Name string `json:"name"`
	Url string `json:"url"`
	Header string `json:"header"`
}
