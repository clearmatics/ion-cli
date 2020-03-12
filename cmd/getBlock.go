package cmd

import (
	"fmt"
	"github.com/clearmatics/ion-cli/backend"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/spf13/viper"

	"github.com/spf13/cobra"
)

// TODO getBlock -number/hash -encoded

var (

	blockInfo string // either the hash or the number
	byHash = false
	block *types.Header
	byteBlock []byte

	getBlockCmd = &cobra.Command{
		Use:   "getBlock",
		Short: "Allow to retrieve a block through a rpc call",
		Long: `Allow to retrieve a block through a rpc call, either by number or by hash, rlp encoded or as object`,
		Run: func(cmd *cobra.Command, args []string) {

			fmt.Println("Connecting to the RPC client..")

			eth, err := backend.GetClient(viper.GetString("rpc"))
			if err != nil {
				fmt.Println(err)
				return
			}

			fmt.Println("Retrieving the block you asked..")

			if !byHash {
				block, byteBlock, err = eth.GetBlockByNumber(blockInfo)
			} else {
				block, byteBlock, err = eth.GetBlockByHash(blockInfo)
			}

			fmt.Println("Success!", block)

			// update the session with what's needed to cache
			session.Block = string(byteBlock)
			session.PersistSession(sessionPath)
		},
	}
)

func init() {

	//getBlockCmd.Flags().BoolVarP(&byHash, "byHash", "c", false, "Specify if reading the block by number or by hash")
	getBlockCmd.Flags().StringVarP(&blockInfo, "block", "b", "latest", "Block number or hash")

	getBlockCmd.MarkFlagRequired("block")

	rootCmd.AddCommand(getBlockCmd)


}
