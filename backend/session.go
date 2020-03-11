package backend

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Session struct {
	Timestamp int `json:"timestamp"`
	Rpc string `json:"rpc"`
	Active bool `json:"active"`
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