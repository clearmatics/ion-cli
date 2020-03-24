package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"strings"
)

// profileCmd represents the profile command
var (

	/* PROFILES */
	profileCmd = &cobra.Command{
		Use: "profile",
		Long: "Manage the profiles to interact with ION",
	}

	profileArgs = []string{"profileID"}
	addProfileCmd = &cobra.Command{
		Use:   "add [" + strings.Join(profileArgs, ",") + "]",
		Short: "Add a profile",
		Long: `Initialize a profile with the specified profileID`,
		Args: func(cmd *cobra.Command, args []string) error {
			return checkArgs(args, profileArgs)
		},
		Run: func(cmd *cobra.Command, args []string) {

			profileId := args[0]

			returnIfError(loadProfiles(profilesPath))

			if profiles.Exist(profileId) {

				if deleteFlag {
					fmt.Println("Deleting profile", profileId)
					profiles.Remove(profileId)
				} else {
					fmt.Println("This profile already exist:\n", profiles[profileId])
				}

			} else {
				fmt.Println("Creating a profile named", profileId)
				profiles.Add(profileId)
			}

			profiles.Save(profilesPath)
		},
	}

)

func init() {
	addProfileCmd.Flags().BoolVarP(&deleteFlag, "delete", "d", false, "Delete the specified profile")

	// tree of commands
	profileCmd.AddCommand(addProfileCmd)
	rootCmd.AddCommand(profileCmd)
}


