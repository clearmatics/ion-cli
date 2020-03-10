package cmd

import (
	"fmt"
	"github.com/abiosoft/ishell"
	"github.com/clearmatics/ion-cli/cli/core"
	"github.com/clearmatics/ion-cli/utils"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/rlp"
)

func CliqueCommands(session *core.Session) []*ishell.Cmd {
	return []*ishell.Cmd{
		//{
		//	Name: "getBlockByNumber_Clique",
		//	Help: "use: \tgetBlockByNumber_Clique [optional rpc Url] [integer]\n\t\t\t\tdescription: Returns signed and unsigned RLP-encoded block headers by block number required for submission to Clique validation from connected Client or specified endpoint",
		//	Func: func(c *ishell.Context) {
		//		if len(c.Args) == 1 {
		//			if ethClient != nil {
		//				block, _, err := GetBlockByNumber(ethClient, c.Args[0])
		//				if err != nil {
		//					c.Println(err)
		//					return
		//				}
		//				signedBlock, unsignedBlock := RlpEncodeClique(block)
		//				c.Printf("Signed Block: %+x\n", signedBlock)
		//				c.Printf("Unsigned Block: %+x\n", unsignedBlock)
		//			} else {
		//				c.Println("Please connect to a Client before invoking this function.\nUse \tconnectToClient [rpc Url] \n")
		//				return
		//			}
		//		} else if len(c.Args) == 2 {
		//			client, err := getClient(c.Args[0])
		//			if err != nil {
		//				c.Println(err)
		//				return
		//			}
		//			block, _, err := GetBlockByNumber(client, c.Args[1])
		//			if err != nil {
		//				c.Println(err)
		//				return
		//			}
		//			signedBlock, unsignedBlock := RlpEncodeClique(block)
		//			c.Printf("Signed Block:\n %+x\n", signedBlock)
		//			c.Printf("Unsigned Block:\n %+x\n", unsignedBlock)
		//		} else {
		//			c.Println("Usage: \tgetBlockByNumber_Clique [optional rpc Url] [integer]\n")
		//			return
		//		}
		//		c.Println("===============================================================")
		//	},
		//},
		//{
		//	Name: "getBlockByHash_Clique",
		//	Help: "use: \tgetBlockByHash_Clique [optional rpc Url] [hash] \n\t\t\t\tdescription: Returns signed and unsigned RLP-encoded block headers by block hash required for submission to Clique validation from connected Client or specified endpoint",
		//	Func: func(c *ishell.Context) {
		//		if len(c.Args) == 1 {
		//			if ethClient != nil {
		//				block, _, err := getBlockByHash(ethClient, c.Args[0])
		//				if err != nil {
		//					c.Println(err)
		//					return
		//				}
		//				signedBlock, unsignedBlock := RlpEncodeClique(block)
		//				c.Printf("Signed Block: 0x%+x\n", signedBlock)
		//				c.Printf("Unsigned Block: 0x%+x\n", unsignedBlock)
		//			} else {
		//				c.Println("Please connect to a Client before invoking this function.\nUse \tconnectToClient [rpc Url] \n")
		//				return
		//			}
		//		} else if len(c.Args) == 2 {
		//			client, err := getClient(c.Args[0])
		//			if err != nil {
		//				c.Println(err)
		//				return
		//			}
		//			block, _, err := getBlockByHash(client, c.Args[1])
		//			if err != nil {
		//				c.Println(err)
		//				return
		//			}
		//			signedBlock, unsignedBlock := RlpEncodeClique(block)
		//			c.Printf("Signed Block:\n %+x\n", signedBlock)
		//			c.Printf("Unsigned Block:\n %+x\n", unsignedBlock)
		//		} else {
		//			c.Println("Usage: \tgetBlockByHash_Clique [optional rpc Url] [hash]\n")
		//			return
		//		}
		//		c.Println("===============================================================")
		//	},
		//},
	}
}

func rlpEncodeClique(blockHeader *types.Header) (rlpSignedBlock []byte, rlpUnsignedBlock []byte, err error) {
	// Encode the orginal block header
	_, err = rlp.EncodeToBytes(&blockHeader)
	if err != nil {
		fmt.Println("can't RLP encode requested block:", err)
		return
	}

	// Generate an interface to encode the blockheader without the signature in the extraData
	rlpSignedBlock, err = utils.RlpEncodeBlock(blockHeader)
	if err != nil {
		return
	}

	rlpUnsignedBlock, err = utils.RlpEncodeUnsignedBlock(blockHeader)
	if err != nil {
		return
	}

	return rlpSignedBlock, rlpUnsignedBlock, nil
}
