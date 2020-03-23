package backend

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"time"
)

type Session struct {
	LastAccess int `json:"timestamp"`
	Profile string `json:"profile"`
}

func (s *Session) Save(path string) error {
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

func (s *Session) Delete(path string) error {
	// TODO other logic might be needed
	s.LastAccess = 0

	// update the file
	return s.Save(path)
}

func (s *Session) IsValid(timeoutSec int) bool {
	// TODO basic logic here
	return int(time.Now().Unix()) - s.LastAccess < timeoutSec
}