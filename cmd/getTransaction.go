package cmd

import (
	"encoding/hex"
	"fmt"
	"github.com/clearmatics/ion-cli/utils"
	"github.com/ethereum/go-ethereum/common"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// getTransactionCmd represents the getTransaction command
var (
	txHash   string
	getProof bool

	getTransactionCmd = &cobra.Command{
		Use:   "getTransaction",
		Short: "Retrieve a transaction object by its hash",
		Long: `Perform a eth_getTransactionByHash rpc call and cache the tx object. 
		Calculate the merkle proof of that tx if the --getProof flag is passed`,
		Run: func(cmd *cobra.Command, args []string) {

			fmt.Println("Connecting to the RPC client..")

			eth, err := utils.Client(viper.GetString("rpc"))
			returnIfError(err)

			fmt.Println("Calling getTransactionByHash for tx:", txHash)
			hash := common.HexToHash(txHash)
			session.Transaction.Tx, _, err = utils.GetTransactionByHash(eth, hash)
			returnIfError(err)

			if getProof {
				// calculate the merkle proof
				fmt.Println("Retrieving the merkle data..")
				data, err := utils.FetchProofData(eth, hash)
				returnIfError(err)

				fmt.Println("Calculating the merkle proof..")
				proof, err := utils.GenerateIonProof(*data)
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
