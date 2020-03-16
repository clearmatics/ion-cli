package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/clearmatics/ion-cli/backend"
	"io/ioutil"
	"time"
)

// TODO add log with verbosity level
// TODO we might move also root configs command to backend
var(
	// flags can be persistent or local - more down below
	Verbose bool

	sessionPath string
	configPath string
	session *backend.Session

	timeoutSec =  3600

	rootCmd = &cobra.Command{
		Use:   "ion-cli",
		Short: "Cross-chain framework tool",
		Long: "Ion is a system and function-agnostic framework for building cross-interacting smart contracts between blockchains and/or systems",

		Run: func(cmd *cobra.Command, args []string) {

		},
		Args: func(cmd *cobra.Command, args []string) error {
			// this to validate positional arguments
			return nil
		},
	}
)

func Execute() {
	// start the cli app
	rootCmd.Execute()
}

func init(){

	// flags
	rootCmd.PersistentFlags().BoolVarP(&Verbose, "verbose", "v", false, "verbose output")
	rootCmd.PersistentFlags().StringVarP(&configPath, "config", "c", "./config/test.json", "Config file to populate the session with")
	rootCmd.PersistentFlags().StringVarP(&sessionPath, "session", "s", "./config/session.json", "Session file to populate the session with")

	// choose config
	initConfig(sessionPath, configPath)

}

// choose whether to override configs with session fields
func initConfig(sessionPath string, configPath string) {

	b, _ := ioutil.ReadFile(sessionPath)
	json.Unmarshal(b, &session)

	if int(time.Now().Unix()) - session.Timestamp < timeoutSec  {
		// update the session
		session.Active = true
		session.Timestamp = int(time.Now().Unix())

		err := session.PersistSession(sessionPath)
		if err != nil {
			fmt.Println(err)
		} else {
			viper.SetConfigFile(sessionPath)
			fmt.Println("Using session configs")
		}

	} else {
		fmt.Println("Using default configs")

		viper.SetConfigFile(configPath)
	}

	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}

}


