package clique

import (
	"flag"
	"fmt"
	"github.com/abiosoft/ishell"
	"github.com/clearmatics/ion-cli/cli/core"
	"github.com/clearmatics/ion-cli/utils"
	"github.com/ethereum/go-ethereum/common"
	"math/big"
)

func CliqueCommands(session *core.Session) []*ishell.Cmd {
	return []*ishell.Cmd{
		{
			Name: "getBlockByNumber_Clique",
			Help: "use: \tgetBlockByNumber_Clique -client -height\n\t\t\t\tdescription: Returns signed and unsigned RLP-encoded block headers by block number required for submission to Clique validation from connected Client",
			Func: func(c *ishell.Context) {
				flagSet := flag.NewFlagSet("getBlockByNumber_Clique", flag.ContinueOnError)
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

					signed, unsigned, err := rlpEncodeClique(block)
					if err != nil {
						c.Println(err)
						return
					}

					c.Printf("Signed Block:\n %+x\n", signed)
					c.Printf("Unsigned Block:\n %+x\n", unsigned)
				} else {
					c.Println(err.Error())
				}

				c.Println("===============================================================")
			},
		},
		{
			Name: "getBlockByHash_Clique",
			Help: "use: \tgetBlockByHash_Clique -client -hash \n\t\t\t\tdescription: Returns signed and unsigned RLP-encoded block headers by block hash required for submission to Clique validation from connected Client",
			Func: func(c *ishell.Context) {
				flagSet := flag.NewFlagSet("getBlockByHash_Clique", flag.ContinueOnError)
				clientName := flagSet.String("client", "", "")
				hash := flagSet.String("hash", "", "")

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
					block, _, err := utils.GetBlockHeaderByHash(client.Client, common.HexToHash(*hash))
					if err != nil {
						c.Println(err)
						return
					}

					signed, unsigned, err := rlpEncodeClique(block)
					if err != nil {
						c.Println(err)
						return
					}

					c.Printf("Signed Block:\n %+x\n", signed)
					c.Printf("Unsigned Block:\n %+x\n", unsigned)
				} else {
					c.Println(err.Error())
				}

				c.Println("===============================================================")
			},
		},
	}
}
