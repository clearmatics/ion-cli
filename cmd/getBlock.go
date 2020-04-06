package cmd

import (
	"fmt"
	"github.com/clearmatics/ion-cli/backend"
	"github.com/spf13/cobra"
	"strings"
)

var (
	blockType  string
	byHash     bool
	rlpEncoded bool
	err error
	//block *backend.BlockInterface
	chain string

	getBlkArgs = []string{"Block hash or number"}
	getBlockCmd = &cobra.Command{
		Use:   "getBlock [" + strings.Join(getBlkArgs, ",") + "]",
		Short: "Allow to retrieve a block through a rpc call",
		Long:  `Allow to retrieve a block through a rpc call, either by number or by hash, rlp encoded or as object`,
		Args: func(cmd *cobra.Command, args []string) error {
			return checkArgs(args, getBlkArgs)
		},
		Run: func(cmd *cobra.Command, args []string) {
			blockInfo := args[0]

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

			// assign the type implementing interfaces in the active chain
			returnIfError(activeChain.AssignImplementers())

			// rpc call
			if !byHash {
				fmt.Println(fmt.Sprintf("Request of retrieving on chain %v block by number: %v\n", chain, blockInfo))
				fmt.Println(rpcURL)

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

			if activeProfile.IsActive() {
				// marshal the typed header into json raw format that will be saved to file
				activeChain.Block.Header, err = activeChain.Block.Interface.Marshal()
				returnIfError(err)

				// update profile chain
				activeProfile.Chains[chain] = activeChain
				returnIfError(profiles.Save(profilesPath))

			} else {
				// just print the block retrieved
				activeChain.Block.Interface.Print()
			}

		},
	}
)

func init() {

	getBlockCmd.Flags().BoolVarP(&rlpEncoded, "rlp", "", false, "Rlp encode the block header as well (default false)")
	getBlockCmd.Flags().BoolVarP(&byHash, "byHash", "", false, "Specify if reading the block by hash (default by number)")
	getBlockCmd.Flags().StringVarP(&chain, "chain", "c", "local", "Chain identifier in the profile")
	getBlockCmd.Flags().StringVarP(&chainType, "type", "", "eth", "Chain structure type")

	// to run with no profiles
	getBlockCmd.Flags().StringVarP(&rpcURL, "rpc", "", "http://127.0.0.1:8545", "URL of the rpc endpoint")

	rootCmd.AddCommand(getBlockCmd)

}

