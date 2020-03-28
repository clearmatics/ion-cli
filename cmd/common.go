package cmd

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/clearmatics/ion-cli/backend"
	"github.com/spf13/viper"
	"io/ioutil"
	"os"
	//"strings"
)

// utils for cli commands these functions will be available to all commands

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
	profiles = backend.Profiles{}
	b, _ := ioutil.ReadFile(profilesPath)
	err := json.Unmarshal(b, &profiles)

	return err
}

func checkArgs(args []string, expected []string) error {
	if len(args) != len(expected) {
		return errors.New("invalid args")
	}

	return nil
}