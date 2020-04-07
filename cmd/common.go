package cmd

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/clearmatics/ion-cli/backend"
	"github.com/spf13/viper"
	"io/ioutil"
	"os"
	"time"

	//"strings"
)

// utils for commands
func returnIfError(err error) {
	if err != nil {
		fmt.Println(err)
		fmt.Println("Quitting..")
		os.Exit(-1)
	}
}

func loadConfig(configPath string) error {
	configs = viper.New()
	configs.SetConfigFile(configPath)
	return configs.ReadInConfig()
}

func loadProfiles(profilesPath string) error {
	jsonFile, err := os.Open(profilesPath)
	if err != nil {
		return err
	}
	defer jsonFile.Close()

	profiles = backend.Profiles{}

	b, _ := ioutil.ReadAll(jsonFile)
	err = json.Unmarshal(b, &profiles)
	return err
}

func initProfile() error {
	returnIfError(loadProfiles(profilesPath))

	// if profile flag is set use that if it's a valid profile
	if profiles.Exist(profileName){
		fmt.Println("Using profile", profileName, "from the flag")

		activeProfile = profiles[profileName]
	} else {
		// check if a session is active
		b, _ := ioutil.ReadFile(sessionPath)
		returnIfError(json.Unmarshal(b, &session))

		if session.IsValid(timeoutSec) && profiles.Exist(session.Profile) {
			fmt.Println("Loading profile", session.Profile, "from the session")

			session.LastAccess = int(time.Now().Unix())
			session.Save(sessionPath)

			activeProfile = profiles[session.Profile]
		}
	}

	return nil
}

func checkArgs(args []string, expected []string) error {
	if len(args) != len(expected) {
		return errors.New("invalid args")
	}

	return nil
}


