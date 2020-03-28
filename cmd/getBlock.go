package cmd

import (
	"errors"
	"fmt"
	"github.com/clearmatics/ion-cli/backend/ethereum"
	"github.com/spf13/cobra"
)

var (
	blockInfo  string
	blockType  string
	byHash     bool
	rlpEncoded bool
	err error
	//block *backend.BlockInterface
	chain string

	getBlockCmd = &cobra.Command{
		Use:   "getBlock",
		Short: "Allow to retrieve a block through a rpc call",
		Long:  `Allow to retrieve a block through a rpc call, either by number or by hash, rlp encoded or as object`,
		Run: func(cmd *cobra.Command, args []string) {

			// TODO active profile might not exist
			if !activeProfile.Chains.Exist(chain) {
				fmt.Println(fmt.Sprintf("The chain %v doesn't exists for profile %v", chain, activeProfile.Name))
				return
			}

			// assign the type implementing the block interface in the chain
			returnIfError(assignBlockImplementer(blockType))

			activeChain := activeProfile.Chains[chain]

			// rpc call
			if !byHash {
				fmt.Println(fmt.Sprintf("Request of retrieving on chain %v block by number: %v\n", chain, blockInfo))

				err = activeChain.Block.Interface.GetByNumber(activeChain.Network.Url, blockInfo)
				returnIfError(err)

			} else {
				fmt.Printf("Request of retrieving on chain %v block by hash: %v\n", chain, blockInfo)

				err = activeChain.Block.Interface.GetByHash(activeProfile.Chains[chain].Network.Url, blockInfo)
				returnIfError(err)
			}


			if rlpEncoded {
				fmt.Println("Rlp encoding it..")
				err := activeChain.Block.Interface.RlpEncode()
				returnIfError(err)
			}

			activeChain.Block.Header, err = activeChain.Block.Interface.Marshal()
			activeChain.Block.Type = blockType

			returnIfError(err)

			activeProfile.Chains[chain] = activeChain

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

func assignBlockImplementer(blockType string) error {
	switch blockType {
	case "eth":
		activeProfile.Chains[chain].Block.Interface = &ethereum.EthBlockHeader{}
	case "clique":
		activeProfile.Chains[chain].Block.Interface = &ethereum.CliqueBlockHeader{}
	default:
		return errors.New(fmt.Sprintf("The block type %v is not recognised", blockType))
	}

	return nil

}
