package cmd

import (
	"encoding/hex"
	"fmt"
	"github.com/clearmatics/ion-cli/backend"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// getTransactionCmd represents the getTransaction command
var (

	txHash string
	getProof bool

	getTransactionCmd = &cobra.Command{
		Use:   "getTransaction",
		Short: "Retrieve a transaction object by its hash",
		Long: `Perform a eth_getTransactionByHash rpc call and cache the tx object. 
		Calculate the merkle proof of that tx if the --getProof flag is passed`,
		Run: func(cmd *cobra.Command, args []string) {

			fmt.Println("Connecting to the RPC client..")

			eth, err := backend.GetClient(viper.GetString("rpc"))
			returnIfError(err)

			fmt.Println("Calling getTransactionByHash for tx:", txHash)
			session.Transaction.Tx, _, err = eth.GetTransactionByHash(txHash)
			returnIfError(err)

			if getProof {
				// calculate the merkle proof
				fmt.Println("Calculating the merkle proof..")
				proof, err := eth.GetProof(txHash)
				returnIfError(err)

				session.Transaction.Proof = hex.EncodeToString(proof)
			}

			// update session file
			returnIfError(session.PersistSession(sessionPath))

			fmt.Println("Success! Session file updated")
		},
	}
)

func init() {
	getTransactionCmd.Flags().StringVarP(&txHash, "hash", "", "", "Hexadecimal tx hash")
	getTransactionCmd.Flags().BoolVarP(&getProof, "getProof", "", false, "Calculate and store the merkle proof of that transaction ")

	getTransactionCmd.MarkFlagRequired("hash")

	rootCmd.AddCommand(getTransactionCmd)

}
