// Copyright (c) 2018 Clearmatics Technologies Ltd

package config
// TODO to be deleted once all the old cli under cli/ is refactored or won-t build
import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/keystore"
)

// Settings
type Setup struct {
	Accounts  []ConfigAccount   `json:"accounts"`
	Contracts []ConfigContracts `json:"contracts"`
	Networks  []ConfigNetworks  `json:"networks"`
}

type ConfigAccount struct {
	Name     string `json:"name"`
	Keyfile  string `json:"keyfile"`
	Password string `json:"password"`
}

type ConfigContracts struct {
	Name string `json:"name"`
	File string `json:"file"`
}

type ConfigNetworks struct {
	Name string `json:"name"`
	Uri  string `json:"uri"`
}

type Account struct {
	Auth *bind.TransactOpts
	Key  *keystore.Key
}

// Takes path to a JSON and returns a struct of the contents
func ReadSetup(config string) (Setup, error) {
	setup := Setup{}
	raw, err := ioutil.ReadFile(config)
	if err != nil {
		return setup, err
	}

	err = json.Unmarshal(raw, &setup)
	if err != nil {
		return setup, err
	}

	return setup, nil
}

// Takes path to a JSON and returns a string of the contents
func ReadString(path string) (contents string) {
	raw, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Print(err, "\n")
	}

	contents = string(raw)

	return
}

func InitUser(privkeystore string, password string) (user Account, err error) {
	// retrieve private key
	keyjson, err := ioutil.ReadFile(privkeystore)
	if err != nil {
		return Account{}, err
	}

	userkey, err := keystore.DecryptKey(keyjson, password)
	if err != nil {
		return Account{}, err
	}

	// Create an authorized transactor
	key := ReadString(privkeystore)
	auth, err := bind.NewTransactor(strings.NewReader(key), password)
	if err != nil {
		return Account{}, err
	}

	return Account{auth, userkey}, nil
}
