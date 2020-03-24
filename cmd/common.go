package cmd

import (
	"encoding/json"
	"errors"
	"fmt"
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
	b, _ := ioutil.ReadFile(profilesPath)
	return json.Unmarshal(b, &profiles)
}

func checkArgs(args []string, expected []string) error {
	if len(args) != len(expected) {
		return errors.New("invalid args")
	}

	return nil
}