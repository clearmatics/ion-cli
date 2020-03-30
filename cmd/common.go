package cmd

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/clearmatics/ion-cli/backend"
	"github.com/clearmatics/ion-cli/backend/ethereum"
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

func assignChainImplementers(chainType string) error {
	switch chainType {
	case "eth":
		activeProfile.Chains[chain].Transaction.Interface = &ethereum.EthTransaction{}
		activeProfile.Chains[chain].Block.Interface = &ethereum.EthBlockHeader{}
	case "clique":
		activeProfile.Chains[chain].Block.Interface = &ethereum.CliqueBlockHeader{}
		activeProfile.Chains[chain].Transaction.Interface = &ethereum.EthTransaction{}
	default:
		return errors.New(fmt.Sprintf("The chain type %v is not recognised", blockType))
	}

	return nil
}
