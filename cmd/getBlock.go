package cmd

import (
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

			// define which block type // TODO should go into the network configs
			switch blockType {
			case "eth":
				block = &ethereum.EthBlockHeader{}
			case "clique":
				block = &ethereum.CliqueBlockHeader{}
			default:
				// TODO
				fmt.Println("This block type is not recognised. Availables are..")
				return
			}

			if !activeProfile.Chains.Exist(chain) {
				fmt.Println(fmt.Sprintf("The chain %v doesn't exists for profile %v", chain, activeProfile.Name))
				return
			}

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

			// add block to profile
			// TODO we might store multiple blocks depending on some counter or smt
			activeProfile.Chains[chain].Blocks["latest"] = block

			// update the profiles
			returnIfError(profiles.Save(profilesPath))

			fmt.Println("Success! Profile updated with the block info")
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
