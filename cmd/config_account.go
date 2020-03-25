package cmd

import (
	"fmt"
	"github.com/clearmatics/ion-cli/backend"
	"github.com/spf13/cobra"
	"strings"
)

var (
	addAccountsArgs = []string{"accountID", "keyfile Path", "password"}
	delAccountArgs = []string{"accountID"}

	accountCmd = &cobra.Command{
		Use: "accounts",
		Short: "Manage the accounts in the configs",
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			// load configs
			returnIfError(loadConfig(configPath))
			return nil
		},
	}

	addAccountCmd = &cobra.Command{
		Use:   "add [" + strings.Join(addAccountsArgs, ",") + "]",
		Short: "Add an account configuration to the configs",
		Args: func(cmd *cobra.Command, args []string) error {
			return checkArgs(args, addAccountsArgs)
		},
		Run: func(cmd *cobra.Command, args []string) {
			accountID := args[0]
			keyFile := args[1]
			password := args[2]

			if !configs.IsSet("accounts." + accountID) || forceFlag {

				fmt.Println(fmt.Sprintf("Creating account %v with the provided info", accountID))

				configs.Set("accounts." + accountID, backend.Account{
					Name: accountID,
					Keyfile: keyFile,
					Password: password,
				})

				returnIfError(configs.WriteConfig())
				fmt.Println("Success!")

			} else {
				fmt.Println(fmt.Sprintf("The account with id %v already exists! Use flag -f to overwrite it", accountID))
				return
			}

		},
	}
	
	deleteAccountCmd = &cobra.Command{
		Use:   "del [" + strings.Join(delAccountArgs, ",") + "]",
		Short: "Delete an account configuration from the configs",
		Args: func(cmd *cobra.Command, args []string) error {
			return checkArgs(args, delAccountArgs)
		},
		Run: func(cmd *cobra.Command, args []string) {

			accountID := args[0]

			fmt.Println(fmt.Sprintf("Deleting account %v from configs", accountID))
			delete(configs.Get("accounts").(map[string]interface{}), accountID)

			returnIfError(configs.WriteConfig())
			fmt.Println("Success!")
		},
	}
)

func init() {
	accountCmd.AddCommand(addAccountCmd)
	accountCmd.AddCommand(deleteAccountCmd)

	configsCmd.AddCommand(accountCmd)
}