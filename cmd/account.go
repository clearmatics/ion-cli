
package cmd

import (
	"fmt"
	"github.com/clearmatics/ion-cli/backend"
	"github.com/spf13/cobra"
)

// TODO remove subcommands

// accountCmd represents the account command
var (

	keyFile string
	password string
	accountsFile string
	accountName string

	accountCmd = &cobra.Command{
		Use:   "accounts",
		Short: "Manage the accounts",
		Long: `Manage the accounts you will use to interact with the ION smart contracts`,
		Run: func(cmd *cobra.Command, args []string) {

			fmt.Println("Fetching the accounts available..")

			err := backend.ListAccounts(accountsFile)
			returnIfError(err)
		},
	}

	addAccountCmd = &cobra.Command{
		Use:   "add",
		Short: "Add an account to the config file",
		Long: `Add an account object to the config file. You will be able to use that in a session just by his name`,
		Run: func(cmd *cobra.Command, args []string) {

			fmt.Println("Storing the new account object.. ")

			// add it to the accounts
			err := backend.AddAccount(accountName, keyFile, password, accountsFile)
			if err != nil {
				fmt.Println(err)
				return
			}

			fmt.Println("Success!")
		},
	}
)

func init() {

	// init sub commands
	initAddAccount()

	rootCmd.AddCommand(accountCmd)
	accountCmd.AddCommand(addAccountCmd)
}

func initAddAccount() {
	addAccountCmd.Flags().StringVarP(&accountName, "name", "n", "", "The name of the account")
	addAccountCmd.Flags().StringVarP(&keyFile, "keyfile", "k", "", "The path to the keyfile")
	addAccountCmd.Flags().StringVarP(&password, "pwd", "p", "", "The password to unlock the account")
	addAccountCmd.Flags().StringVarP(&accountsFile, "accountsFile", "a", "./config/accounts.json", "The file containing the accounts")

	returnIfError(addAccountCmd.MarkFlagRequired("name"))
	returnIfError(addAccountCmd.MarkFlagRequired("keyfile"))
	returnIfError(addAccountCmd.MarkFlagRequired("pwd"))
}
