package cmd

import (
	"fmt"
	//"github.com/clearmatics/ion-cli/backend"
	"github.com/spf13/cobra"
	"time"
)

// TODO use command format as profile and config
var (
	deleteSession bool

	sessionCmd = &cobra.Command{
		Use:   "session",
		Short: "Manage a session within ION",
		Long:  "A session keeps track of the latest profile using ION and will use that as default for all calls",
		Run: func(cmd *cobra.Command, args []string) {

			if deleteSession {

				// delete the session
				fmt.Println("Deleting session..")

				err := session.Delete(sessionPath)
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

			if len(args) != 1 {
				fmt.Println("Specify the profile you want to initialize the session for")
				return
			}

			// create a new session
			session.LastAccess = int(time.Now().Unix())
			session.Profile = args[0]

			fmt.Println(fmt.Sprintf("Creating a new session for profile %v", args[0]))

			err := session.Save(sessionPath)
			returnIfError(err)

			fmt.Println("Success!")

		},
	}
)

func init() {
	// root command
	sessionCmd.Flags().BoolVarP(&deleteSession, "delete", "d", false, "Delete the current session")

	// create the tree of commands
	rootCmd.AddCommand(sessionCmd)
	sessionCmd.AddCommand(addSessionCmd)
}
