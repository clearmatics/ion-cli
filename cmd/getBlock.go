package cmd

import (
	"encoding/hex"
	"fmt"
	"github.com/clearmatics/ion-cli/backend"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/spf13/viper"

	"github.com/spf13/cobra"
)

// TODO getBlock -encoded

var (

	blockInfo string // either the hash or the number
	byHash bool
	rlpEncoded bool
	block *types.Header
	byteBlock []byte
	sessionBlock backend.BlockHeader

	getBlockCmd = &cobra.Command{
		Use:   "getBlock",
		Short: "Allow to retrieve a block through a rpc call",
		Long: `Allow to retrieve a block through a rpc call, either by number or by hash, rlp encoded or as object`,
		Run: func(cmd *cobra.Command, args []string) {

			fmt.Println("Connecting to the RPC client..")

			eth, err := backend.GetClient(viper.GetString("rpc"))
			returnIfError(err)

			if !byHash {
				fmt.Printf("Retrieving block by number: %v\n", blockInfo)
				block, byteBlock, err = eth.GetBlockByNumber(blockInfo)
				returnIfError(err)
			} else {
				fmt.Printf("Retrieving block by hash: %v\n", blockInfo)
				block, byteBlock, err = eth.GetBlockByHash(blockInfo)
				returnIfError(err)
			}

			// cache the block in the session
			session.Block.Header = block

			if rlpEncoded {
				// cache the rlp encoding of that block in the session
				fmt.Println("Rlp encoding it..")
				rlp, err := backend.RlpEncode(block)
				returnIfError(err)

				session.Block.RlpEncoded = hex.EncodeToString(rlp)
			}


			// update the session
			returnIfError(session.PersistSession(sessionPath))

			fmt.Println("Success! Session file updated")
		},
	}
)

func init() {

	getBlockCmd.Flags().BoolVarP(&rlpEncoded, "rlp", "", false, "Specify if the returned block header should be rlp encoded or not")
	getBlockCmd.Flags().BoolVarP(&byHash, "byHash", "", false, "Specify if reading the block by number or by hash")
	getBlockCmd.Flags().StringVarP(&blockInfo, "block", "b", "latest", "Block number or hash")

	rootCmd.AddCommand(getBlockCmd)

}

