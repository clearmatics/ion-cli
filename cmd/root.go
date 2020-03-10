package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"io/ioutil"
	"time"
)

type Session struct {
	Timestamp int `json:"timestamp"`
	Rpc string `json:"rpc"`
	Active bool `json:"active"`
}


var(
	// flags can be persistent or local - more down below
	Verbose bool
	Configs *viper.Viper

	sessionPath string
	configPath string
	session Session

	timeoutSec =  3600

	rootCmd = &cobra.Command{
		Use:   "ion-cli",
		Short: "Cross-chain framework tool",
		Long: "Ion is a system and function-agnostic framework for building cross-interacting smart contracts between blockchains and/or systems",

		Run: func(cmd *cobra.Command, args []string) {
			//fmt.Println("Hi from the ION cli. Type ion-cli help to display the help")

			fmt.Println("root called")

		},
		Args: func(cmd *cobra.Command, args []string) error {
			// this to validate positional arguments
			return nil
		},
	}
)

func Execute() {
	// start the app
	rootCmd.Execute()
}

func init(){
	fmt.Println("root init called")

	// flags
	rootCmd.PersistentFlags().BoolVarP(&Verbose, "verbose", "v", false, "verbose output")
	rootCmd.Flags().StringVarP(&configPath, "config", "c", "./config/test.json", "Config file to populate the session with")
	rootCmd.Flags().StringVarP(&sessionPath, "sessionF", "s", "./config/session.json", "Session file to populate the session with")

	// choose config
	initConfig(sessionPath, configPath)

	// add all commands
	rootCmd.AddCommand(versionCmd)
	rootCmd.AddCommand(sessionCmd)
}

// choose whether to override configs with session fields
func initConfig(sessionPath string, configPath string) {

	b, _ := ioutil.ReadFile(sessionPath)
	json.Unmarshal(b, &session)

	// TODO use verbose to determine log
	// TODO unit test

	fmt.Println(session)

	if int(time.Now().Unix()) - session.Timestamp < timeoutSec  {
		viper.SetConfigFile(sessionPath)

		// update the session
		session.Active = true

		session.Timestamp = int(time.Now().Unix())
		b, err := json.Marshal(session)
		if err != nil {
			fmt.Errorf("error marshaling the session object")
		}

		err = ioutil.WriteFile(sessionPath, b, 0644)
		if err != nil {
			fmt.Errorf("error updating the session file")
		}

		fmt.Println("Overriding with session configs")
	} else {
		fmt.Println("Using default configs")
		viper.SetConfigFile(configPath)
	}

	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}

}

