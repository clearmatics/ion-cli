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

func checkArgs(args []string, expected []string) error {
	if len(args) != len(expected) {
		return errors.New("invalid args")
	}

	return nil
}

func assignChainImplementers(chainObj *backend.Chain) error {
	switch chainObj.Type {
	case "eth":
		chainObj.Transaction.Interface = &ethereum.EthTransaction{}
		chainObj.Block.Interface = &ethereum.EthBlockHeader{}
	case "clique":
		chainObj.Block.Interface = &ethereum.CliqueBlockHeader{}
		chainObj.Transaction.Interface = &ethereum.EthTransaction{}
	case "ibft":
		chainObj.Block.Interface = &ethereum.IBFTBlockHeader{}
		chainObj.Transaction.Interface = &ethereum.EthTransaction{}
	default:
		return errors.New(fmt.Sprintf("The chain type %v is not recognised", blockType))
	}

	return nil
}


