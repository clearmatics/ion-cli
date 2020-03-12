package backend

import(
	"encoding/json"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"io/ioutil"
	"strings"
)

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

var accounts []AccountInfo

// initialize a wallet ready to be used
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
	key := ReadString(keyStore)
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

	return PersistObject(accounts, accountsPath)
}
