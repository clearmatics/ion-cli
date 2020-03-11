package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"time"
)
// TODO we could have flags similar to the truffle configs indicating what part of the more general configs to use
// todo we might restore a session entirely from file
// TODO override single flags instead of using file

var (
	deleteSession bool

	sessionCmd = &cobra.Command{
		Use:   "session",
		Short: "Manage a session within ION",
		Long: "Allow to create, restore or delete a session file further calls would read the configs from and populate with needed data for other calls:",
		Run: func(cmd *cobra.Command, args []string) {

			if deleteSession {
				// delete the session
				session.Active = false
				session.Timestamp = 0

				fmt.Println("Deleting session..")

			} else {

				// create a new session
				session.Active = true
				session.Timestamp = int(time.Now().Unix())

				fmt.Println("Creating a new session..")
			}

			err := session.PersistSession(sessionPath)
			if err == nil {
				fmt.Println("Success!")
			}
		},
	}
)

func init() {
	sessionCmd.Flags().BoolVarP(&deleteSession, "delete", "d", false, "Delete the current session")

	rootCmd.AddCommand(sessionCmd)
}


