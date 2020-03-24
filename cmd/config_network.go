package cmd

import (
	"fmt"
	"github.com/clearmatics/ion-cli/backend"
	"github.com/spf13/cobra"
	"strings"
)

var (

	addNetworkArgs = []string{"networkID", "URLs"}
	delNetworkArgs = []string{"networkID"}

	networkCmd = &cobra.Command{
		Use: "networks",
		Short: "Manage networks in your configs",
	}

	addNetworkCmd = &cobra.Command{
		Use:   "add [" + strings.Join(addNetworkArgs, ",") + "]",
		Short: "Add a network configuration to the configs",
		Args: func(cmd *cobra.Command, args []string) error {
			return checkArgs(args, addNetworkArgs)
		},
		Run: func(cmd *cobra.Command, args []string) {
			networkId := args[0]
			networkURL := args[1]

			if !configs.IsSet("networks." + networkId) {

				fmt.Println(fmt.Sprintf("Creating network %v with the provided info", networkId))

				configs.Set("networks." + networkId, backend.NetworkInfo{
					Name: networkId,
					Url:  networkURL,
				})

			} else {
				fmt.Println(fmt.Sprintf("The network with id %v already exists!", networkId))
				return
			}

			returnIfError(configs.WriteConfig())

		},
	}

	deleteNetworkCmd = &cobra.Command{
		Use:   "del [" + strings.Join(delNetworkArgs, ",") + "]",
		Short: "Add a network configuration to the configs",
		Args: func(cmd *cobra.Command, args []string) error {
			return checkArgs(args, delNetworkArgs)
		},
		Run: func(cmd *cobra.Command, args []string) {

			networkId := args[0]
			fmt.Println(fmt.Sprintf("Deleting network %v from configs", networkId))
			delete(configs.Get("networks").(map[string]interface{}), networkId)

			returnIfError(configs.WriteConfig())
			fmt.Println("Success")
		},
	}
)

func init() {

	networkCmd.AddCommand(addNetworkCmd)
	networkCmd.AddCommand(deleteNetworkCmd)
	configsCmd.AddCommand(networkCmd)
}