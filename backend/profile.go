package backend

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Profile struct {
	Name string `json:"name"`
	Chains Chains`json:"chains"`
}

type Profiles map[string]Profile

// store all the profiles to disk
func (p Profiles) Save(path string) error {
	b, err := json.MarshalIndent(p, "", "	")
	if err != nil {
		fmt.Errorf("error marshaling the profile object")
		return err
	}

	err = ioutil.WriteFile(path, b, 0644)
	if err != nil {
		fmt.Errorf("error updating the profile file")
		return err
	}

	fmt.Println("Profiles updated!")
	return nil
}

// tells if a profile name id exists
func (p Profiles) Exist(id string) bool {
	return p[id].Name != ""
}

// initialize a profile with profileId
func (p Profiles) Add (profileId string) {
	p[profileId] = Profile{
		Name:   profileId,
		Chains: Chains{},
	}
}

// remove profile with profiledId from profiles
func (p Profiles) Remove (profileId string) {
	if p.Exist(profileId) {
		delete(p, profileId)
	}
}
