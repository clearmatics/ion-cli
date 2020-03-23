package backend

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Chain struct {
	Network Network `json:"network"`
	Accounts map[string]AccountInfo `json:"accounts"`

	// cache
	Blocks BlockMap `json:"blocks"`
	Transaction Transaction `json:"transaction"`

	// ion proofs

	// contracts
}

type Network struct {
	Name string `json:"name"`
	Url string `json:"url"`
}

type Profile struct {
	Name string `json:"name"`
	Chains Chains`json:"networks"`
}

type Profiles map[string]Profile
type Chains map[string]Chain

// store all the profiles to disk
func (p Profiles) Save(path string) error {
	b, err := json.MarshalIndent(p, "", "	")
	if err != nil {
		fmt.Errorf("error marshaling the session object")
		return err
	}

	err = ioutil.WriteFile(path, b, 0644)
	if err != nil {
		fmt.Errorf("error updating the session file")
		return err
	}

	fmt.Println("Profiles updated!")
	return nil
}

func (p Profiles) Exist(id string) bool {
	return p[id].Name != ""
}

func (c Chains) Exist(id string) bool {
	return c[id].Network != Network{}
}

// initialize a profile object
func InitProfile(id string) *Profile {
	return &Profile{
		Name:   id,
		Chains: make(map[string]Chain),
	}
}


