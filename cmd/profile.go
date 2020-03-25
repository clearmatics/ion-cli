package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"strings"
)

/* PROFILES */

var (

	profileCmd = &cobra.Command{
		Use: "profile",
		Short: "Manage the profiles to interact with ION",
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			return loadProfiles(profilesPath)
		},
		PersistentPostRunE: func(cmd *cobra.Command, args []string) error {
			return profiles.Save(profilesPath)
		},
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

			if profiles.Exist(profileId) {

				fmt.Println("This profile already exist:\n", profiles[profileId])

			} else {
				fmt.Println("Creating a profile named", profileId)
				profiles.Add(profileId)
			}

		},
	}

	delProfileCmd = &cobra.Command{
		Use:   "del [" + strings.Join(profileArgs, ",") + "]",
		Short: "Delete a profile",
		Run: func(cmd *cobra.Command, args []string) {
			profileId := args[0]

			fmt.Println("Deleting profile", profileId)
			profiles.Remove(profileId)

		},
	}
)

func init() {

	// tree of commands
	profileCmd.AddCommand(addProfileCmd)
	profileCmd.AddCommand(delProfileCmd)

	rootCmd.AddCommand(profileCmd)
}


