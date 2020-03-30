package backend

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/keystore"
)

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
