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



			fmt.Println("Using session configs from now on..")
		},
	}
)

func init() {
	fmt.Println("session init called")

	// TODO overrite single flags instead of using file
	sessionCmd.Flags().StringVarP(&path, "config", "c", "./config/session.json", "Config file to populate the session with")
}
