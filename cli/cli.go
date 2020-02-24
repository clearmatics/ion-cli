// Copyright (c) 2018 Clearmatics Technologies Ltd
package cli

import (
	//"errors"
	"flag"
	"fmt"
	"github.com/clearmatics/ion-cli/cli/cmd"
	"github.com/clearmatics/ion-cli/cli/core"
	//"math/big"
	//"strconv"
	//"strings"
	//
	//"github.com/ethereum/go-ethereum/accounts/abi/bind"
	//"github.com/ethereum/go-ethereum/common"

	"github.com/abiosoft/ishell"

	"github.com/clearmatics/ion-cli/config"
)

func printWelcome() {
	// display welcome info.
	fmt.Println("===============================================================")
	fmt.Print("Ion Command Line Interface\n\n")
	fmt.Println("Use 'help' to list commands")
	fmt.Println("===============================================================")
}

// Launch - definition of commands and creates the interface
func Launch(setup *config.Setup) {
	// by default, new shell includes 'exit', 'help' and 'clear' commands.
	shell := ishell.New()

	// Create new context
	//ctx := context.Background()

	session := core.InitSession()

	if setup != nil {
		// Add all accounts in config to memory
		for _, account := range setup.Accounts {
			user, err := config.InitUser(account.Keyfile, account.Password)
			if err != nil {
				fmt.Printf("Setup Failed: Adding Account %s from configuration failed %s", account.Name, err.Error())
				return
			}
			session.Accounts[account.Name] = &user
		}

		// Compile and add all contract instances to memory
		for _, configContract := range setup.Contracts {
			compiledContract, err := core.CompileContract(session, configContract.File)
			if err != nil {
				fmt.Printf("Setup Failed: Compiling contract %s from configuration failed: %s", configContract.Name, err.Error())
				return
			}
			session.Contracts[configContract.Name] = compiledContract
		}

		// Compile and add all contract instances to memory
		for _, configNetwork := range setup.Networks {
			client, err := core.GetClient(configNetwork.Uri)
			if err != nil {
				fmt.Printf("Could not connect to client %s\n", configNetwork.Name)
				return
			}

			session.Networks[configNetwork.Name] = client
		}
	}

	//---------------------------------------------------------------------------------------------
	// 	RPC Client Specific Commands
	//---------------------------------------------------------------------------------------------

	shell.AddCmd(&ishell.Cmd{
		Name: "test",
		Help: "use: \tconnectToClient [rpc Url] \n\t\t\t\tdescription: Connects to an RPC Client to be used",
		Func: func(c *ishell.Context) {
			f1 := flag.NewFlagSet("f1", flag.ContinueOnError)
			silent := f1.String("me", "no one", "")

			if err := f1.Parse([]string{"-me=you"}); err == nil {
				fmt.Println("apply", *silent)
			}

			c.Println("===============================================================")
		},
	})

	for _, command := range cmd.CoreCommands(session) {
		shell.AddCmd(command)
	}

	/*shell.AddCmd(&ishell.Cmd{
			Name: "callMessage",
			Help: "use: \tcallMessage [contract name] [function name] [from account name] [deployed contract address] \n\t\t\t\tdescription: Connects to an RPC Client to be used",
			Func: func(c *ishell.Context) {
				if len(c.Args) != 4 {
	                c.Println("Usage: \tcallMessage [contract name] [function name] [from account name] [deployed contract address] \n")
				} else {
				    if ethClient == nil {
				        c.Println("Please connect to a Client before invoking this function.\nUse \tconnectToClient [rpc Url] \n")
				        return
				    }

	                instance := contracts[c.Args[0]]
	                methodName := c.Args[1]
	                account := accounts[c.Args[2]]
	                contractDeployedAddress := common.HexToAddress(c.Args[3])

	                if instance == nil {
	                    errStr := fmt.Sprintf("Contract instance %s not found.\nUse \taddContractInstances [name] [path/to/solidity/contract] [deployed address] \n", c.Args[0])
	                    c.Println(errStr)
				        return
	                }
	                if account == nil {
	                    errStr := fmt.Sprintf("Account %s not found.\nUse \taddAccount [name] [path/to/keystore]\n", c.Args[2])
				        c.Println(errStr)
				        return
	                }

	                if instance.Abi.Methods[methodName].Name == "" {
	                    c.Printf("Method name \"%s\" not found for contract \"%s\"\n", methodName, c.Args[0])
	                    return
	                }

	                inputs, err := parseMethodParameters(c, instance.Abi, methodName)
	                if err != nil {
	                    c.Printf("Error parsing parameters: %s\n", err)
	                    return
	                }

	                var out interface{}

	                out, err = contract.CallContract(
	                    ctx,
	                    ethClient.Client,
	                    instance.Contract,
	                    account.Key.Address,
	                    contractDeployedAddress,
	                    c.Args[1],
	                    out,
	                    inputs...
	                )
	                 if err != nil {
	                    c.Println(err)
	                    return
	                 } else {
	                    c.Printf("Result: %s\n", out)
	                 }
				}
				c.Println("===============================================================")
			},
		})*/

	//shell.AddCmd(&ishell.Cmd{
	//	Name: "getTransactionByHash",
	//	Help: "use: \tgetTransactionByHash [optional rpc Url] [hash]\n\t\t\t\tdescription: Returns transaction specified by hash from connected Client or specified endpoint",
	//	Func: func(c *ishell.Context) {
	//		var json []byte
	//		var err error
	//
	//		if len(c.Args) == 1 {
	//			if ethClient != nil {
	//				_, json, err = getTransactionByHash(ethClient, c.Args[0])
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
	//			_, json, err = getTransactionByHash(client, c.Args[1])
	//			if err != nil {
	//				c.Println(err)
	//				return
	//			}
	//		} else {
	//			c.Println("Usage: \tgetTransactionByHash [optional rpc Url] [hash]\n")
	//			return
	//		}
	//		if err != nil {
	//			c.Println(err)
	//			return
	//		}
	//		c.Printf("Transaction: %s\n", json)
	//		c.Println("===============================================================")
	//	},
	//})
	//
	//shell.AddCmd(&ishell.Cmd{
	//	Name: "GetBlockByNumber",
	//	Help: "use: \tGetBlockByNumber [optional rpc Url] [integer]\n\t\t\t\tdescription: Returns block header specified by height from connected Client or from specified endpoint",
	//	Func: func(c *ishell.Context) {
	//		var json []byte
	//		var err error
	//
	//		if len(c.Args) == 1 {
	//			if ethClient != nil {
	//				_, json, err = GetBlockByNumber(ethClient, c.Args[0])
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
	//			_, json, err = GetBlockByNumber(client, c.Args[1])
	//			if err != nil {
	//				c.Println(err)
	//				return
	//			}
	//		} else {
	//			c.Println("Usage: \tGetBlockByNumber [optional rpc Url] [integer]\n")
	//			return
	//		}
	//		if err != nil {
	//			c.Println(err)
	//			return
	//		}
	//		c.Printf("Block: %s\n", json)
	//		c.Println("===============================================================")
	//	},
	//})
	//
	//shell.AddCmd(&ishell.Cmd{
	//	Name: "getBlockByHash",
	//	Help: "use: \tgetBlockByHash [optional rpc Url] [hash] \n\t\t\t\tdescription: Returns block header specified by hash from connected Client or from specific endpoint",
	//	Func: func(c *ishell.Context) {
	//		var json []byte
	//		var err error
	//
	//		if len(c.Args) == 1 {
	//			if ethClient != nil {
	//				_, json, err = getBlockByHash(ethClient, c.Args[0])
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
	//			_, json, err = getBlockByHash(client, c.Args[1])
	//			if err != nil {
	//				c.Println(err)
	//				return
	//			}
	//		} else {
	//			c.Println("Usage: \tgetBlockByHash [optional rpc Url] [hash] \n")
	//			return
	//		}
	//		if err != nil {
	//			c.Println(err)
	//			return
	//		}
	//		c.Printf("Block: %s\n", json)
	//		c.Println("===============================================================")
	//	},
	//})
	//
	//shell.AddCmd(&ishell.Cmd{
	//	Name: "getEncodedBlockByHash",
	//	Help: "use: \tgetEncodedBlockByHash [optional rpc Url] [hash] \n\t\t\t\tdescription: Returns RLP-encoded block header specified by hash from connected Client or from specific endpoint",
	//	Func: func(c *ishell.Context) {
	//		if len(c.Args) == 1 {
	//			if ethClient != nil {
	//				block, _, err := getBlockByHash(ethClient, c.Args[0])
	//				if err != nil {
	//					c.Println(err)
	//					return
	//				}
	//				encodedBlock, err := RlpEncode(block)
	//				if err != nil {
	//					c.Println(err)
	//					return
	//				}
	//				c.Printf("Encoded Block: %+x\n", encodedBlock)
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
	//			block, _, err := getBlockByHash(client, c.Args[0])
	//			if err != nil {
	//				c.Println(err)
	//				return
	//			}
	//			encodedBlock, err := RlpEncode(block)
	//			if err != nil {
	//				c.Println(err)
	//				return
	//			}
	//			c.Printf("Encoded Block:\n %+x\n", encodedBlock)
	//		} else {
	//			c.Println("Usage: \tgetEncodedBlockByHash [optional rpc Url] [integer]\n")
	//			return
	//		}
	//		c.Println("===============================================================")
	//	},
	//})
	//
	//shell.AddCmd(&ishell.Cmd{
	//	Name: "getEncodedBlockByNumber",
	//	Help: "use: \tgetEncodedBlockByNumber [optional rpc Url] [hash] \n\t\t\t\tdescription: Returns RLP-encoded block header specified by number from connected Client or from specific endpoint",
	//	Func: func(c *ishell.Context) {
	//		if len(c.Args) == 1 {
	//			if ethClient != nil {
	//				block, _, err := GetBlockByNumber(ethClient, c.Args[0])
	//				if err != nil {
	//					c.Println(err)
	//					return
	//				}
	//				encodedBlock, err := RlpEncode(block)
	//				if err != nil {
	//					c.Println(err)
	//					return
	//				}
	//				c.Printf("Encoded Block: %+x\n", encodedBlock)
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
	//			block, _, err := GetBlockByNumber(client, c.Args[0])
	//			if err != nil {
	//				c.Println(err)
	//				return
	//			}
	//			encodedBlock, err := RlpEncode(block)
	//			if err != nil {
	//				c.Println(err)
	//				return
	//			}
	//			c.Printf("Encoded Block:\n %+x\n", encodedBlock)
	//		} else {
	//			c.Println("Usage: \tgetEncodedBlockByNumber [optional rpc Url] [integer]\n")
	//			return
	//		}
	//		c.Println("===============================================================")
	//	},
	//})
	//
	//shell.AddCmd(&ishell.Cmd{
	//	Name: "getProof",
	//	Help: "use: \tgetProof [optional rpc Url] [Transaction Hash] \n\t\t\t\tdescription: Returns a merkle patricia proof of a specific transaction and its receipt in a block",
	//	Func: func(c *ishell.Context) {
	//		if len(c.Args) == 1 {
	//			if ethClient != nil {
	//				getProof(ethClient, c.Args[0])
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
	//			getProof(client, c.Args[1])
	//		} else {
	//			c.Println("Usage: \tgetProof [optional rpc Url] [Transaction hash] \n")
	//			return
	//		}
	//		c.Println("===============================================================")
	//	},
	//})

	printWelcome()
	shell.Run()
	session.Close()
	shell.Close()
}
