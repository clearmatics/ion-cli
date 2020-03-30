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

	chainArgs = []string{"profileID", "chainID", "networkID", "chainType"}
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
			networkId := args[2]
			chainType := args[3]

			if profiles.Exist(profileId) {

				if profiles[profileId].Chains.Exist(chainId) || forceFlag {
					fmt.Println(fmt.Sprintf("Creating chain %v in profile %v", chainId, profileId))
					returnIfError(loadConfig(configPath))

					network := backend.NetworkInfo{}
					returnIfError(configs.UnmarshalKey("networks." + networkId, &network))

					profiles[profileId].Chains.Add(chainId, network, chainType)
				} else {
					fmt.Println(fmt.Sprintf("Chain with id %v already exists! Use flag -f to overwrite it", chainId))
				}

			} else {
				fmt.Println("This profile does not exists yet! Initialize it first")
				// TODO or we can initialize it here
			}
		},
	}

	delChainArgs = []string{"profileID", "chainID"}
	delChainCmd = &cobra.Command{
		Use:   "del [" + strings.Join(delChainArgs, ",") + "]",
		Short: "Delete a chain from a profile",
		Long: ``,
		Args: func(cmd *cobra.Command, args []string) error {
			return checkArgs(args, delChainArgs)
		},
		Run: func(cmd *cobra.Command, args []string) {

			profileId := args[0]
			chainId := args[1]

			fmt.Println(fmt.Sprintf("Deleting chain %v in profile %v", chainId, profileId))
			profiles[profileId].Chains.Remove(chainId)
		},
	}
)
func init() {
	// tree of commands
	chainsCmd.AddCommand(addChainCmd)
	chainsCmd.AddCommand(delChainCmd)

	profileCmd.AddCommand(chainsCmd)
}
