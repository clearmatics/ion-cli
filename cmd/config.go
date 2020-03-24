package cmd

import (
	"fmt"
	"github.com/clearmatics/ion-cli/backend"
	"github.com/spf13/cobra"
)

// configsCmd represents the configs command
var (

	accountName, password, keyFile string
	networkName, networkURL string

	/* CONFIGS */
	configsCmd = &cobra.Command{
		Use:   "configs",
		Short: "",
		Long: ``,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("TODO list configs")
		},
	}

	/* WALLETS INFO */
	configWalletCmd = &cobra.Command{
		Use:   "wallets",
		Short: "Add or delete wallet configuration to the configs",
		Long: `Add or delete wallet configurations to the configs`,
		Run: func(cmd *cobra.Command, args []string) {

			if len(args) != 1 {
				fmt.Println("Please specify the walletId your are referring to")
				return
			}

			returnIfError(loadConfig(configPath))

			walletId := args[0]

			if deleteFlag {

				fmt.Println(fmt.Sprintf("Deleting wallet %v from configs", walletId))
				delete(configs.Get("accounts").(map[string]interface{}), walletId)

			} else if !configs.IsSet("accounts." + walletId) {

				fmt.Println(fmt.Sprintf("Creating wallet %v with the provided info", walletId))

				configs.Set("accounts." + walletId, backend.WalletInfo{
					Name: walletId,
					Keyfile: keyFile,
					Password: password,
				})

			} else {
				fmt.Println(fmt.Sprintf("The wallet with id %v already exists!", walletId))
				return
			}

			returnIfError(configs.WriteConfig())

		},
	}

	/* NETWORK INFO */
	configNetworkCmd = &cobra.Command{
		Use:   "network",
		Short: "Add or delete chains configuration to the configs",
		Long: `Add or delete chains configurations to the configs`,
		Run: func(cmd *cobra.Command, args []string) {

			if len(args) != 1 {
				fmt.Println("Please specify the chainId your are referring to")
				return
			}

			returnIfError(loadConfig(configPath))

			networkId := args[0]

			if deleteFlag {

				fmt.Println(fmt.Sprintf("Deleting network %v from configs", networkId))
				delete(configs.Get("networks").(map[string]interface{}), networkId)

			} else if !configs.IsSet("networks." + networkId) {

				fmt.Println(fmt.Sprintf("Creating network %v with the provided info", networkId))

				configs.Set("networks." + networkId, backend.NetworkInfo{
					Name: networkName,
					Url:  networkURL,
				})

			} else {
				fmt.Println(fmt.Sprintf("The network with id %v already exists!", networkId))
				return
			}

			returnIfError(configs.WriteConfig())

		},
	}
)

func init() {
	// init subcommands flags
	initAddWallet()
	initAddNetwork()

	// TREE OF COMMANDS
	rootCmd.AddCommand(configsCmd)
	configsCmd.AddCommand(configWalletCmd)
	configsCmd.AddCommand(configNetworkCmd)
}


func initAddWallet() {
	configWalletCmd.Flags().BoolVarP(&deleteFlag, "delete", "d", false, "Delete the specified wallet info")
	configWalletCmd.Flags().StringVarP(&accountName, "name", "n", "", "The name of the account")
	configWalletCmd.Flags().StringVarP(&keyFile, "keyfile", "k", "", "The path to the keyfile")

	configWalletCmd.Flags().StringVarP(&password, "pwd", "", "", "The password to unlock the account")
}

func initAddNetwork() {
	configNetworkCmd.Flags().BoolVarP(&deleteFlag, "delete", "d", false, "Delete the specified wallet info")
	configNetworkCmd.Flags().StringVarP(&networkName, "name", "n", "", "The name of the network")
	configNetworkCmd.Flags().StringVarP(&networkURL, "keyfile", "u", "url", "The url to connect to the network")
}
