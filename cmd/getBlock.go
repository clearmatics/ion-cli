package cmd

import (
	"encoding/hex"
	"errors"
	"fmt"
	"github.com/clearmatics/ion-cli/utils"
	"github.com/ethereum/go-ethereum/common"
	"github.com/spf13/viper"
	"math/big"

	"github.com/spf13/cobra"
)

var (
	blockInfo  string // either the hash or the number
	byHash     bool
	rlpEncoded bool

	getBlockCmd = &cobra.Command{
		Use:   "getBlock",
		Short: "Allow to retrieve a block through a rpc call",
		Long:  `Allow to retrieve a block through a rpc call, either by number or by hash, rlp encoded or as object`,
		Run: func(cmd *cobra.Command, args []string) {

			fmt.Println("Connecting to the RPC client..")

			eth, err := utils.Client(viper.GetString("rpc"))
			returnIfError(err)

			// assign the block to session object
			if !byHash {
				fmt.Printf("Retrieving block by number: %v\n", blockInfo)

				number := new(big.Int)
				_, ok := number.SetString(blockInfo, 10)
				if !ok {
					returnIfError(errors.New(fmt.Sprintf("Invalid block number: %s", blockInfo)))
				}

				session.Block.Header, _, err = utils.GetBlockHeaderByNumber(eth, number)
				returnIfError(err)
			} else {
				fmt.Printf("Retrieving block by hash: %v\n", blockInfo)
				hash := common.HexToHash(blockInfo)
				session.Block.Header, _, err = utils.GetBlockHeaderByHash(eth, hash)
				returnIfError(err)
			}

			// add the rlp encoding if flagged
			if rlpEncoded {
				// cache the rlp encoding of that block in the session
				fmt.Println("Rlp encoding it..")
				rlp, err := utils.RlpEncodeBlock(session.Block.Header)
				returnIfError(err)

				session.Block.RlpEncoded = hex.EncodeToString(rlp)
			}

			// update session file
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
