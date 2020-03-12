package backend

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Session struct {
	Timestamp int `json:"timestamp"`
	// lenght of the session
	// network
	Rpc string `json:"rpc"`
	Active bool `json:"active"`
	Account Wallet `json:"account"`
}

// TODO reads from the accounts object and store in session the one requested
func (s *Session) ParseAccount(accountName string) error {

	// retrieve private key
	//configs, err := ioutil.ReadFile(configFile)
	//if err != nil {
	//	return err
	//}
	//
	//userkey, err = keystore.DecryptKey(keyjson, password)
	//if err != nil {
	//	return nil, nil, err
	//}
	//
	//// Create an authorized transactor
	//key := ReadString(privkeystore)
	//auth, err = bind.NewTransactor(strings.NewReader(key), password)
	//if err != nil {
	//	return nil, nil, err
	//}

	return nil
}

func (s *Session) PersistSession(path string) error {
	b, err := json.Marshal(s)
	if err != nil {
		fmt.Errorf("error marshaling the session object")
		return err
	}

	err = ioutil.WriteFile(path, b, 0644)
	if err != nil {
		fmt.Errorf("error updating the session file")
		return err
	}

	return nil
}