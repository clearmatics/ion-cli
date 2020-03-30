package cmd

import (
	"fmt"
	"github.com/clearmatics/ion-cli/backend/ethereum"
	"github.com/spf13/cobra"
	"strings"
)


var (
	txHash   string
	rpcURL string

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

			// if no active profiles just do a on the fly call and print the output
			if !activeProfile.IsActive() {
				runWithNoProfiles(txHash)
				return
			}

			if !activeProfile.Chains.Exist(chain) {
				fmt.Println(fmt.Sprintf("The chain %v doesn't exists for profile %v", chain, activeProfile.Name))
				return
			}

			// assign the type implementing the tx interface in the chain
			returnIfError(assignChainImplementers(activeProfile.Chains[chain].Type))

			activeChain := activeProfile.Chains[chain]

			fmt.Println("Retrieving and generating ION proof for tx:", txHash)
			err := activeChain.Transaction.Interface.GenerateIonProof(activeChain.Network.Url, txHash)
			returnIfError(err)

			// marshal the typed header into json raw format that will be saved to file
			activeChain.Transaction.Tx, err = activeChain.Transaction.Interface.Marshal()
			returnIfError(err)

			// persist the updates on the active profile
			activeProfile.Chains[chain] = activeChain
			returnIfError(profiles.Save(profilesPath))
		},
	}
)

func init() {
	getTransactionCmd.Flags().StringVarP(&chain, "chain", "c", "local", "Chain identifier in the profile")
	getTransactionCmd.Flags().StringVarP(&rpcURL, "rpc", "", "", "URL of the rpc endpoint")

	getTransactionCmd.MarkFlagRequired("hash")

	
	rootCmd.AddCommand(getTransactionCmd)

}

func runWithNoProfiles(txHash string) {
	transactionObj := ethereum.EthTransaction{
		Tx:    nil,
		Proof: "",
	}
	returnIfError(transactionObj.GenerateIonProof(rpcURL, txHash))

	fmt.Println("Success! Here's the proof", transactionObj.Proof)
}

