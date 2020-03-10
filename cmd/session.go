package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"io/ioutil"
	"time"
)

// sessionCmd represents the session command
var (

	deleteSession bool

	sessionCmd = &cobra.Command{
		Use:   "session",
		Short: "Initialize a new session of ion",
		Long: "Initialize a session file further calls would read the configs from and populate with needed data for other calls:",
		Run: func(cmd *cobra.Command, args []string) {

			// TODO we could have flags similar to the truffle configs indicating what part of the more general configs to use
			if deleteSession {

				// delete the session
				session.Active = false
				session.Timestamp = 0

				err := persistSession(session, sessionPath)
				if err == nil {
					fmt.Println("Session deleted..")
				}

			} else {

				// create a new session
				session.Active = true
				session.Timestamp = int(time.Now().Unix())

				err := persistSession(session, sessionPath)
				if err == nil {
					fmt.Println("Session created..")
				}
			}
		},
	}
)

func init() {
	fmt.Println("session init called")

	// TODO override single flags instead of using file
	sessionCmd.Flags().BoolVarP(&deleteSession, "delete", "d", false, "Delete the current session")
}

func persistSession(session Session, path string) error {
	b, err := json.Marshal(session)
	if err != nil {
		fmt.Errorf("error marshaling the session object")
		return err
	}

	err = ioutil.WriteFile(path, b, 0644)
	if err != nil {
		fmt.Errorf("error updating the session file")
		return err
	}

	return nil
}
