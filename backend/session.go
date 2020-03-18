package backend

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"time"
)

type Session struct {
	Timestamp int `json:"timestamp"`
	// lenght of the session

	// network
	Rpc         string `json:"rpc"`
	Active      bool   `json:"active"`
	AccountName string `json:"account"`

	// enable polymorphism
	Blocks BlockMap `json:"blocks"`

	Transaction Transaction `json:"transaction"`
}

func (s *Session) PersistSession(path string) error {
	b, err := json.MarshalIndent(s, "", "	")
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

func (s *Session) DeleteSession(path string) error {
	// TODO other logic might be needed
	s.Active = false
	s.Timestamp = 0

	// update the file
	return s.PersistSession(path)
}

func (s *Session) IsValid(timeoutSec int) bool {
	// TODO basic logic here
	return int(time.Now().Unix()) - s.Timestamp < timeoutSec
}