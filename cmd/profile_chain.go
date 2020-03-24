package cmd

import (
	"fmt"
	"github.com/clearmatics/ion-cli/backend"
	"strings"

	"github.com/spf13/cobra"
)

/* CHAINS OF A PROFILE */

var (
	// TODO doc
	chainsCmd = &cobra.Command{
		Use: "chains",
		Short: "Manage the chains configuration of a profile",
	}

	chainArgs = []string{"profileID", "networkID"}
	addChainCmd = &cobra.Command{
		Use:   "add [" + strings.Join(chainArgs, ",") + "]",
		Short: "Add or delete a chain from a profile",
		Long: ``,
		Args: func(cmd *cobra.Command, args []string) error {
			return checkArgs(args, chainArgs)
		},
		Run: func(cmd *cobra.Command, args []string) {

			profileId := args[0]
			chainId := args[1]

			returnIfError(loadProfiles(profilesPath))

			if profiles.Exist(profileId) {

				if deleteFlag {

					fmt.Println(fmt.Sprintf("Deleting chain %v in profile %v", chainId, profileId))
					profiles[profileId].Chains.Remove(chainId)

				} else {

					fmt.Println(fmt.Sprintf("Creating chain %v in profile %v", chainId, profileId))
					returnIfError(loadConfig(configPath))

					network := backend.NetworkInfo{}
					returnIfError(configs.UnmarshalKey("networks." + chainId, &network))

					profiles[profileId].Chains.Add(chainId, network)

				}

				profiles.Save(profilesPath)
			} else {
				fmt.Println("This profile does not exists yet! Create it first")
				// or we can initialize it here
			}
		},
	}
)
func init() {
	chainsCmd.Flags().BoolVarP(&deleteFlag, "delete", "d", false, "Delete the specified chain")

	// tree of commands
	chainsCmd.AddCommand(addChainCmd)

	profileCmd.AddCommand(chainsCmd)
}
