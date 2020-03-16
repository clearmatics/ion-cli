package backend

import (
	"encoding/json"
	"fmt"
	"github.com/ethereum/go-ethereum/core/types"
	"io/ioutil"
)

// TODO might need to specify on which consensus
// blockheader fields that are needed to be cached for the app
type BlockHeader struct {
	Header *types.Header `json:"header"`
	RlpEncoded string`json:"rlp_encoded"`
}

type Session struct {
	Timestamp int `json:"timestamp"`
	// lenght of the session

	// network
	Rpc string `json:"rpc"`
	Active bool `json:"active"`
	AccountName string `json:"account"`

	// fields that have to be cached for subsequent calls
	Block BlockHeader `json:"block"`
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