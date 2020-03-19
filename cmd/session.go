package cmd

import (
	"fmt"
	//"github.com/clearmatics/ion-cli/backend"
	"github.com/spf13/cobra"
	"time"
)

// todo we might restore a session entirely from file
// TODO override single flags of using file
// TODO all the flags containing network, account, etc to use
// TODo clean session cached objects
var (
	deleteSession bool

	sessionCmd = &cobra.Command{
		Use:   "session",
		Short: "Manage a session within ION",
		Long:  "Allow to create, restore or delete a session file further calls would read the configs from and populate with needed data for other calls:",
		Run: func(cmd *cobra.Command, args []string) {

			if deleteSession {

				// delete the session
				fmt.Println("Deleting session..")

				err := session.DeleteSession(sessionPath)
				returnIfError(err)

				fmt.Println("Success!")

			} else {
				fmt.Printf("These are the session parameters you are using: \n%+v", session)
			}
		},
	}

	addSessionCmd = &cobra.Command{
		Use:   "init",
		Short: "Add a session within ION",
		Long:  "Allow to create, a session file further calls would read the configs from and populate with needed data for other calls:",
		Run: func(cmd *cobra.Command, args []string) {

			// create a new session
			session.Active = true
			session.Timestamp = int(time.Now().Unix())
			session.AccountName = accountName

			fmt.Println("Creating a new session..")

			err := session.PersistSession(sessionPath)
			returnIfError(err)

			fmt.Println("Success!")

		},
	}
)

func init() {
	// root command
	sessionCmd.Flags().BoolVarP(&deleteSession, "delete", "d", false, "Delete the current session")

	// sub commands
	initAddCmd()

	// create the tree of commands
	rootCmd.AddCommand(sessionCmd)
	sessionCmd.AddCommand(addSessionCmd)
}

func initAddCmd() {
	// add sub command
	addSessionCmd.Flags().StringVarP(&accountName, "account", "a", "", "The account name to use in the session")
	addSessionCmd.MarkFlagRequired("account")
}
