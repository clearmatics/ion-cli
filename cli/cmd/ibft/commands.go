package ibft

import (
	"flag"
	"fmt"
	"github.com/abiosoft/ishell"
	"github.com/clearmatics/ion-cli/cli/core"
	"github.com/clearmatics/ion-cli/utils"
	"github.com/ethereum/go-ethereum/rlp"
	"math/big"
	//"github.com/clearmatics/ion-cli/cli"
)

func IBFTCommands(session *core.Session) []*ishell.Cmd {
	return []*ishell.Cmd{
		{
			Name: "getBlockByNumber_IBFT",
			Help: "use: \tgetBlockByNumber_IBFT [optional rpc url] [integer]\n\t\t\t\tdescription: Returns proposal, commital RLP-encoded block headers and commit seals by block number required for submission to IBFT validation from connected client or specified endpoint",
			Func: func(c *ishell.Context) {
				flagSet := flag.NewFlagSet("getBlockByNumber_IBFT", flag.ContinueOnError)
				clientName := flagSet.String("client", "", "")
				number := flagSet.Int64("height", 0, "")

				if err := flagSet.Parse(c.Args); err == nil {
					if flagSet.NFlag() != 2 {
						flagSet.Usage()
						return
					}

					if _, ok := session.Networks[*clientName]; !ok {
						c.Println(fmt.Sprintf("Client %s not recognised. Please use addClient to add a new connected client first or specify a correct Client name.", *clientName))
						return
					}

					client := session.Networks[*clientName]
					block, _, err := utils.GetBlockHeaderByNumber(client.Client, big.NewInt(*number))
					if err != nil {
						c.Println(err)
						return
					}

					extra, err := ExtractIstanbulExtra(block)
					if err != nil {
						c.Println(fmt.Sprintf("extracting ibft extra data error: %s", err.Error()))
						return
					}

					committedSeals, err := rlp.EncodeToBytes(extra.CommittedSeal)
					if err != nil {
						c.Println(fmt.Sprintf("rlp encoding committed seals error: %s", err.Error()))
						return
					}

					proposalBlock, commitBlock, err := rlpEncodeIBFT(block)
					if err != nil {
						c.Println(fmt.Sprintf("encoding block error: %s", err.Error()))
					}

					c.Printf("Proposal Block:\n %+x\n", proposalBlock)
					c.Printf("Committed Block:\n %+x\n", commitBlock)
					c.Printf("Commit Seals:\n %+x\n", committedSeals)
				} else {
					c.Println(err.Error())
				}

				c.Println("===============================================================")
			},
		},
		//{
		//	Name: "getBlockByHash_IBFT",
		//	Help: "use: \tgetBlockByHash_IBFT [optional rpc url] [integer]\n\t\t\t\tdescription: Returns proposal, commital RLP-encoded block headers and commit seals by block hash required for submission to IBFT validation from connected client or specified endpoint",
		//	Func: func(c *ishell.Context) {
		//		if len(c.Args) == 1 {
		//			if ethClient != nil {
		//				block, _, err := cli.GetBlockByHash(ethClient, c.Args[0])
		//				if err != nil {
		//					c.Println(err)
		//					return
		//				}
		//				proposalBlock, commitBlock, seals := RlpEncodeIBFT(block)
		//				c.Printf("Proposal Block:\n %+x\n", proposalBlock)
		//				c.Printf("Commital Block:\n %+x\n", commitBlock)
		//				c.Printf("Commit Seals:\n %+x\n", seals)
		//			} else {
		//				c.Println("Please connect to a Client before invoking this function.\nUse \tconnectToClient [rpc url] \n")
		//				return
		//			}
		//		} else if len(c.Args) == 2 {
		//			client, err := getClient(c.Args[0])
		//			if err != nil {
		//				c.Println(err)
		//				return
		//			}
		//			block, _, err := cli.GetBlockByHash(client, c.Args[1])
		//			if err != nil {
		//				c.Println(err)
		//				return
		//			}
		//			proposalBlock, commitBlock, seals := RlpEncodeIBFT(block)
		//			c.Printf("Proposal Block:\n %+x\n", proposalBlock)
		//			c.Printf("Commital Block:\n %+x\n", commitBlock)
		//			c.Printf("Commit Seals:\n %+x\n", seals)
		//		} else {
		//			c.Println("Usage: \tgetBlockByHash_IBFT [optional rpc url] [integer]\n")
		//			return
		//		}
		//		c.Println("===============================================================")
		//	},
		//},
	}
}
