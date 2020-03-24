package cmd

import (
	"fmt"
	"github.com/clearmatics/ion-cli/backend"
	"strings"

	"github.com/spf13/cobra"
)

/* ACCOUNTS OF A CHAIN OF A PROFILE */

var (
	profileId, chainId, accountId string
	profileAccountsArgs = []string{"profileID", "chainID", "accountID"}

	// TODO doc
	profileAccountsCmd = &cobra.Command{
		Use: "accounts",
		Short: "Manage the accounts on the profile chains",
	}

	addProfileAccountCmd = &cobra.Command{
		Use:   "add [" + strings.Join(profileAccountsArgs, ",") + "]",
		Short: "Add account configuration to the profile",
		Long: `Add account configurations to be used for a specific chain and profile`,
		Args: func(cmd *cobra.Command, args []string) error {
			return checkArgs(args, profileAccountsArgs)
		},
		Run: func(cmd *cobra.Command, args []string) {

			profileId = args[0]
			chainId = args[1]
			accountId = args[2]

			if profiles.Exist(profileId) {

				if profiles[profileId].Chains.Exist(chainId) {

					fmt.Println(fmt.Sprintf("Creating accounts %v of chain %v in profile %v", accountId, chainId, profileId))
					returnIfError(loadConfig(configPath))

					account := backend.Account{}
					returnIfError(configs.UnmarshalKey("accounts."+ accountId, &account))

					profiles[profileId].Chains[chainId].Accounts.Add(accountId, account)

				} else {
					fmt.Println("This chain does not exists yet! Add it to the profile first")
				}

			} else {
				fmt.Println("This profile does not exists yet! Create it first")
			}
		},
	}

	delProfileAccountCmd = &cobra.Command{
		Use:   "del [" + strings.Join(profileAccountsArgs, ",") + "]",
		Short: "Delete account configuration to the profile",
		Long: `Delete account configurations for a specific chain and profile`,
		Args: func(cmd *cobra.Command, args []string) error {
			return checkArgs(args, profileAccountsArgs)
		},
		Run: func(cmd *cobra.Command, args []string) {

			profileId = args[0]
			chainId = args[1]
			accountId = args[2]

			fmt.Println(fmt.Sprintf("Deleting account %v of chain %v in profile %v", accountId, chainId, profileId))
			profiles[profileId].Chains[chainId].Accounts.Remove(accountId)
		},
	}
)

func init() {

	profileAccountsCmd.AddCommand(addProfileAccountCmd)
	profileAccountsCmd.AddCommand(delProfileAccountCmd)

	profileCmd.AddCommand(profileAccountsCmd)
}
