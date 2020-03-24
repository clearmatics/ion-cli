package cmd

// TODO refactor with profile logic
//var (
//	blockInfo  string // either the hash or the number
//	blockType  string
//	byHash     bool
//	rlpEncoded bool
//
//	block backend.BlockInterface
//	network string
//
//	getBlockCmd = &cobra.Command{
//		Use:   "getBlock",
//		Short: "Allow to retrieve a block through a rpc call",
//		Long:  `Allow to retrieve a block through a rpc call, either by number or by hash, rlp encoded or as object`,
//		Run: func(cmd *cobra.Command, args []string) {
//
//			// TODO checks on args if allowed before going on
//
//			// define which block type
//			switch blockType {
//			case "eth":
//				block = &ethereum.EthBlockHeader{}
//			case "clique":
//				block = &ethereum.CliqueBlockHeader{}
//			default:
//				// TODO
//				fmt.Println("This block type is not recognised. Availables are..")
//				return
//			}
//
//			// rpc call
//			if !byHash {
//				fmt.Printf("Request of retrieving %v block by number: %v\n", blockType, blockInfo)
//
//				err := block.GetByNumber(viper.GetString("networks." + network), blockInfo)
//				returnIfError(err)
//
//			} else {
//				fmt.Printf("Request of retrieving %v block by hash: %v\n", blockType, blockInfo)
//
//				err := block.GetByHash(viper.GetString("networks." + network), blockInfo)
//				returnIfError(err)
//			}
//
//			// add the rlp encoding to the object if flagged
//			if rlpEncoded {
//				fmt.Println("Rlp encoding it..")
//				err := block.RlpEncode()
//				returnIfError(err)
//
//			}
//			// add block to session struct
//			session.Blocks[blockType] = block
//
//			// update session file
//			returnIfError(session.PersistSession(sessionPath))
//
//			fmt.Println("Success! Session file updated")
//		},
//	}
//)
//
//func init() {
//
//	getBlockCmd.Flags().BoolVarP(&rlpEncoded, "rlp", "", false, "Specify if the returned block header should be rlp encoded or not")
//	getBlockCmd.Flags().BoolVarP(&byHash, "byHash", "", false, "Specify if reading the block by number or by hash")
//	getBlockCmd.Flags().StringVarP(&blockInfo, "block", "b", "latest", "Block number or hash")
//	getBlockCmd.Flags().StringVarP(&blockType, "type", "t", "eth", "Block header type format")
//	getBlockCmd.Flags().StringVarP(&network, "network", "n", "local", "Network name as in config")
//
//	rootCmd.AddCommand(getBlockCmd)
//
//}
