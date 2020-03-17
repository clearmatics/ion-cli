package backend

import(
	"encoding/json"
	"fmt"
	"github.com/clearmatics/ion-cli/utils"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"io/ioutil"
	"reflect"
	"strings"
)

var accounts []AccountInfo

// initialize a wallet ready to be used
// TODO to be called with viper account name in use to have a transactor ready
func InitAccount(name string, keyStore string, password string) (Wallet, error) {

	w := Wallet{}

	// retrieve private key
	b, err := ioutil.ReadFile(keyStore)
	if err != nil {
		return Wallet{}, err
	}

	w.Key, err = keystore.DecryptKey(b, password)
	if err != nil {
		return Wallet{}, err
	}

	// Create an authorized transactor
	key := utils.ReadString(keyStore)
	w.Auth, err = bind.NewTransactor(strings.NewReader(key), password)
	if err != nil {
		return Wallet{}, err
	}

	// Add its identifier
	w.Name = name

	return w, nil
}

// store info to accounts file
func AddAccount(name string, keystore string, password string, accountsPath string) error {

	// get all accounts
	b, err := ioutil.ReadFile(accountsPath)
	if err != nil {
		return err
	}
	err = json.Unmarshal(b, &accounts)
	if err != nil {
		return err
	}

	// append this new one
	accounts = append(accounts, AccountInfo{name, keystore, password})

	return utils.PersistObject(accounts, accountsPath)
}

// unmarshal all the available accounts
func FetchAccounts(accountsPath string) ([]AccountInfo, error) {
	// get all accounts
	b, err := ioutil.ReadFile(accountsPath)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(b, &accounts)
	if err != nil {
		return nil, err
	}

	return accounts, nil
}

// print the list of accounts
func ListAccounts(accountPath string) (error) {
	accounts, err := FetchAccounts(accountPath)

	if err != nil {
		return err
	}

	if len(accounts) == 0 {
		fmt.Println("You have no accounts yet! Type ion-cli help to see how to add one")
		return nil
	}

	for i := 0; i < len(accounts); i++ {
		fmt.Println("---------------------")

		fields := reflect.TypeOf(accounts[i])
		values := reflect.ValueOf(accounts[i])

		for j:=0; j < fields.NumField(); j++ {
			fmt.Println(fields.Field(j).Name, "=", values.Field(j))
		}
	}

	return nil
}
