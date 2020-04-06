package cmd

import (
	"fmt"
	"github.com/clearmatics/ion-cli/backend"
	"github.com/spf13/cobra"
	"strings"
)


var (

	getTxArgs = []string{"transaction Hash"}
	getTransactionCmd = &cobra.Command{
		Use:   "getTransactionProof [" + strings.Join(getTxArgs, ",") + "]",
		Short: "Retrieve a transaction object by its hash and assign to it the ION proof",
		Long: `Perform a eth_getTransactionByHash rpc call and cache the tx object.
		Calculate the merkle proof of that tx if the --getProof flag is passed`,
		Args: func(cmd *cobra.Command, args []string) error {
			return checkArgs(args, getTxArgs)
		},
		Run: func(cmd *cobra.Command, args []string) {
			txHash := args[0]

			if !activeProfile.IsActive() {

				fmt.Println("No profile in use..")

				// run with no profile
				activeChain = backend.Chain{
					Network: backend.NetworkInfo{
						Name:   "",
						Url:    rpcURL,
						Header: "",
					},
					Type:chainType,
				}
			} else {

				if !activeProfile.Chains.Exist(chain){
					fmt.Println(fmt.Sprintf("The chain %v doesn't exists for profile %v", chain, activeProfile.Name))
					return
				}

				// use profile chain
				activeChain = activeProfile.Chains[chain]
			}

			// assign the type implementing the tx interface in the chain
			returnIfError(assignChainImplementers(&activeChain))


			fmt.Println("Retrieving and generating ION proof for tx:", txHash)
			err := activeChain.Transaction.Interface.GenerateIonProof(activeChain.Network.Url, txHash)
			returnIfError(err)


			if activeProfile.IsActive() {
				// marshal the typed header into json raw format that will be saved to file
				activeChain.Transaction.Tx, err = activeChain.Transaction.Interface.Marshal()
				returnIfError(err)

				// update profile chain
				activeProfile.Chains[chain] = activeChain
				returnIfError(profiles.Save(profilesPath))

			} else {
				// just print the block retrieved
				activeChain.Transaction.Interface.Print()
			}

		},
	}
)

func init() {
	getTransactionCmd.Flags().StringVarP(&chain, "chain", "c", "local", "Chain identifier in the profile")
	getTransactionCmd.Flags().StringVarP(&rpcURL, "rpc", "", "http://127.0.0.1:8545", "URL of the rpc endpoint")

	rootCmd.AddCommand(getTransactionCmd)
}


