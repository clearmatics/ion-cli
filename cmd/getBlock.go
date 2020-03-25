package cmd

import (
	"errors"
	"fmt"
	"github.com/clearmatics/ion-cli/backend"
	"github.com/clearmatics/ion-cli/backend/ethereum"
	"github.com/spf13/cobra"
)

var (
	blockInfo  string
	blockType  string
	byHash     bool
	rlpEncoded bool

	block backend.BlockInterface
	chain string

	getBlockCmd = &cobra.Command{
		Use:   "getBlock",
		Short: "Allow to retrieve a block through a rpc call",
		Long:  `Allow to retrieve a block through a rpc call, either by number or by hash, rlp encoded or as object`,
		Run: func(cmd *cobra.Command, args []string) {

			if !activeProfile.Chains.Exist(chain) {
				fmt.Println(fmt.Sprintf("The chain %v doesn't exists for profile %v", chain, activeProfile.Name))
				return
			}

			// assign the type implementing the block interface in the chain
			returnIfError(assignBlockImplementer())

			// pointer
			block  = activeProfile.Chains[chain].Blocks["latest"]

			// rpc call
			if !byHash {
				fmt.Println(fmt.Sprintf("Request of retrieving on chain %v block by number: %v\n", chain, blockInfo))

				err := block.GetByNumber(activeProfile.Chains[chain].Network.Url, blockInfo)
				returnIfError(err)

			} else {
				fmt.Printf("Request of retrieving on chain %v block by hash: %v\n", chain, blockInfo)

				err := block.GetByHash(activeProfile.Chains[chain].Network.Url, blockInfo)
				returnIfError(err)
			}

			// add the rlp encoding to the object if flagged
			if rlpEncoded {
				fmt.Println("Rlp encoding it..")
				err := block.RlpEncode()
				returnIfError(err)
			}

			// persist the updates on the active profile
			returnIfError(profiles.Save(profilesPath))
		},
	}
)

func init() {

	getBlockCmd.Flags().BoolVarP(&rlpEncoded, "rlp", "", false, "Specify if the returned block header should be rlp encoded or not")
	getBlockCmd.Flags().BoolVarP(&byHash, "byHash", "", false, "Specify if reading the block by number or by hash")
	getBlockCmd.Flags().StringVarP(&blockInfo, "block", "b", "latest", "Block number or hash")
	getBlockCmd.Flags().StringVarP(&blockType, "type", "t", "eth", "Block header type format")

	// to override profile configs if active
	getBlockCmd.Flags().StringVarP(&chain, "chain", "c", "local", "Chain identifier in the profile")

	rootCmd.AddCommand(getBlockCmd)

}

func assignBlockImplementer() error {

	switch activeProfile.Chains[chain].Network.Header {
	case "eth":
		activeProfile.Chains[chain].Blocks["latest"] = &ethereum.EthBlockHeader{}
	case "clique":
		activeProfile.Chains[chain].Blocks["latest"] = &ethereum.CliqueBlockHeader{}
	default:
		return errors.New(fmt.Sprintf("The block type %v in chain %v is not recognised.", activeProfile.Chains[chain].Network.Header, chain))
	}

	return nil

}
