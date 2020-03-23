package backend

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Profile struct {
	Name string `json:"name"`
	Chains Chains`json:"networks"`
}

type Profiles map[string]Profile

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

func (p Profiles) Add (profileId string) {
	p[profileId] = Profile{
		Name:   profileId,
		Chains: make(map[string]Chain),
	}
}

func (p Profiles) Remove (profileId string) {
	delete(p, profileId)
}


