package backend

import (
	//"encoding/json"
	//"fmt"
	"github.com/clearmatics/ion-cli/config"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"io/ioutil"
	//"reflect"
	"strings"
)

// an account info as stored in the configs
type Account struct {
	Name     string `json:"name"`
	Keyfile  string `json:"keyfile"`
	Password string `json:"password"`
}
type Accounts map[string]Account

func (a Accounts) Exist(id string) bool {
	return a[id].Name != ""
}

func (a Accounts) Add (id string, account Account) {
	a[id] = account
}

func (a Accounts) Remove (id string) {
	delete(a, id)
}

// initialize a wallet ready to be used to transact
func (a Account) Unlock() (Wallet, error) {

	w := Wallet{}

	// retrieve private key
	b, err := ioutil.ReadFile(a.Keyfile)
	if err != nil {
		return Wallet{}, err
	}

	w.Key, err = keystore.DecryptKey(b, a.Password)
	if err != nil {
		return Wallet{}, err
	}

	// Create an authorized transactor
	key := config.ReadString(a.Keyfile)
	w.Auth, err = bind.NewTransactor(strings.NewReader(key), a.Password)
	if err != nil {
		return Wallet{}, err
	}

	// Add its identifier
	w.Name = a.Name

	return w, nil
}


// print the list of accounts
//func (a Accounts) ListAccounts(accountPath string) error {
//	accounts, err := FetchAccounts(accountPath)
//
//	if err != nil {
//		return err
//	}
//
//	if len(accounts) == 0 {
//		fmt.Println("You have no accounts yet! Type ion-cli help to see how to add one")
//		return nil
//	}
//
//	for i := 0; i < len(accounts); i++ {
//		fmt.Println("---------------------")
//
//		fields := reflect.TypeOf(accounts[i])
//		values := reflect.ValueOf(accounts[i])
//
//		for j := 0; j < fields.NumField(); j++ {
//			fmt.Println(fields.Field(j).Name, "=", values.Field(j))
//		}
//	}
//
//	return nil
//}
