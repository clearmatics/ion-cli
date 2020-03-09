package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

// sessionCmd represents the session command
var (

	path string

	sessionCmd = &cobra.Command{
		Use:   "session",
		Short: "Initialize a new session of ion",
		Long: "Initialize a session file further calls would read the configs from and populate with needed data for other calls:",
		Run: func(cmd *cobra.Command, args []string) {

			// TODO we could have flags similar to the truffle configs indicating what part of the more general configs to use

			// initialize new viper config object called session that would be checked against by commands
			// in order to understand whether to read from default configs or from session
			// configs should never store in transit data (from a call needed into another)
			// session should store the configs plus the transit data

			Session.SetConfigFile(path)
			err := Session.ReadInConfig()
			if err != nil {
				panic(fmt.Errorf("Fatal error session file: %s \n", err))
			}

			fmt.Println("Using session configs from now on..")
		},
	}
)

func init() {
	fmt.Println("session init called")

	//rootCmd.AddCommand(sessionCmd)
	//sessionCmd.AddCommand(versionCmd)

	// TODO id of session file
	sessionCmd.Flags().StringVarP(&path, "config", "c", "./config/session.json", "Config file to populate the session with")
}
