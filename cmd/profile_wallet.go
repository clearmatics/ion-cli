package cmd

import (
	"fmt"
	"github.com/clearmatics/ion-cli/backend"
	"strings"

	"github.com/spf13/cobra"
)

/* WALLETS OF A CHAIN OF A PROFILE */

var (

	walletsCmd = &cobra.Command{
		Use: "wallets",
		Long: "Manage the wallets of a profile",
	}

	walletArgs = []string{"profileID", "chainID", "walletID"}
	addWalletCmd = &cobra.Command{
		Use:   "add [" + strings.Join(walletArgs, ",") + "]",
		Short: "Add or delete wallet configuration to the profile",
		Long: `Add or delete wallet configurations to be used for a specific chain and profile`,
		Args: func(cmd *cobra.Command, args []string) error {
			return checkArgs(args, walletArgs)
		},
		Run: func(cmd *cobra.Command, args []string) {

			profileId := args[0]
			chainId := args[1]
			walletId := args[2]

			returnIfError(loadProfiles(profilesPath))

			if profiles.Exist(profileId){

				if profiles[profileId].Chains.Exist(chainId) {

					if deleteFlag {

						fmt.Println(fmt.Sprintf("Deleting wallet %v of chain %v in profile %v", walletId, chainId, profileId))
						profiles[profileId].Chains[chainId].Accounts.Remove(walletId)

					} else {
						fmt.Println(fmt.Sprintf("Creating wallet %v of chain %v in profile %v", walletId, chainId, profileId))

						returnIfError(loadConfig(configPath))

						account := backend.WalletInfo{}
						returnIfError(configs.UnmarshalKey("accounts."+ walletId, &account))

						profiles[profileId].Chains[chainId].Accounts.Add(walletId, account)
					}

					profiles.Save(profilesPath)

				} else {
					fmt.Println("This chain does not exists yet! Add it to the profile first")
				}

			} else {
				fmt.Println("This profile does not exists yet! Create it first")
			}

		},
	}
)

func init() {
	addWalletCmd.Flags().BoolVarP(&deleteFlag, "delete", "d", false, "Delete the specified wallet")

	walletsCmd.AddCommand(addWalletCmd)

	profileCmd.AddCommand(walletsCmd)
}
