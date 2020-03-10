package cmd

import (
	_ "context"
	"flag"
	"fmt"
	"github.com/abiosoft/ishell"
	"github.com/clearmatics/ion-cli/cli/core"
	"github.com/clearmatics/ion-cli/config"
	"github.com/clearmatics/ion-cli/contracts"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"

	"github.com/clearmatics/ion-cli/utils"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"math/big"
	"strconv"
	"strings"
)

func CoreCommands(session *core.Session) []*ishell.Cmd {
	return []*ishell.Cmd{
		{
			Name: "addClient",
			Help: "use: \taddClient -name -uri \n\t\t\t\tdescription: Connects to an RPC client to be used",
			Func: func(c *ishell.Context) {

				flagSet := flag.NewFlagSet("addClient", flag.ContinueOnError)
				uri := flagSet.String("uri", "", "")
				name := flagSet.String("name", "", "")

				if err := flagSet.Parse(c.Args); err == nil {
					if flagSet.NFlag() != 2 {
						flagSet.Usage()
						return
					}

					c.Println("Connecting to client...\n")
					client, err := core.GetClient(*uri)
					if err != nil {
						c.Println("Could not connect to client.\n")
						return
					}

					session.Networks[*name] = client
					c.Println("Connected!")
				} else {
					c.Println(err.Error())
				}
				c.Println("===============================================================")
			},
		},
		{
			Name: "listClients",
			Help: "use: \tlistClients \n\t\t\t\tdescription: Lists all connected clients",
			Func: func(c *ishell.Context) {
				for key, value := range session.Networks {
					c.Println(fmt.Sprintf("%s: %s", key, value.Url))
				}
				c.Println("===============================================================")
			},
		},
		{
			Name: "addContractInstance",
			Help: "use: \taddContractInstance -name -file\n\t\t\t\tdescription: Compiles a contract for use",
			Func: func(c *ishell.Context) {

				flagSet := flag.NewFlagSet("addContractInstance", flag.ContinueOnError)
				name := flagSet.String("name", "", "Name to use as reference")
				path := flagSet.String("file", "", "Path to solidity contract file")

				if err := flagSet.Parse(c.Args); err == nil {
					if flagSet.NFlag() != 2 {
						flagSet.Usage()
						return
					}

					compiledContract, err := core.AddCompilerAndCompileContract(session, *path)
					if err != nil {
						c.Println(err)
						return
					}

					session.Contracts[*name] = compiledContract
					c.Println("Added!")
				} else {
					c.Println(err.Error())
				}
				c.Println("===============================================================")
			},
		},
		{
			Name: "listContracts",
			Help: "use: \tlistContracts \n\t\t\t\tdescription: List compiled contract instances",
			Func: func(c *ishell.Context) {
				for key := range session.Contracts {
					c.Println(key)
				}
				c.Println("===============================================================")
			},
		},
		{
			Name: "addAccount",
			Help: "use: \taddAccount -name -keyfile\n\t\t\t\tdescription: Add account to be used for transactions",
			Func: func(c *ishell.Context) {
				flagSet := flag.NewFlagSet("addAccount", flag.ContinueOnError)
				name := flagSet.String("name", "", "Name to use as reference")
				path := flagSet.String("keyfile", "", "Path to encrypted keyfile as ethereum wallet format")

				if err := flagSet.Parse(c.Args); err == nil {
					if flagSet.NFlag() != 2 {
						flagSet.Usage()
						return
					}

					c.ShowPrompt(false)
					defer c.ShowPrompt(true)

					c.Println("Please provide your key decryption password.")

					input := c.ReadPassword()
					account, err := config.InitUser(*path, input)
					if err != nil {
						c.Println(err)
						return
					}
					session.Accounts[*name] = &account

					c.Println("Account added succesfully.")
				} else {
					c.Println(err.Error())
				}

				c.Println("===============================================================")
			},
		},
		{
			Name: "listAccounts",
			Help: "use: \tlistAccounts \n\t\t\t\tdescription: List all added accounts",
			Func: func(c *ishell.Context) {
				for key := range session.Accounts {
					c.Println(key)
				}
				c.Println("===============================================================")
			},
		},
		{
			Name: "deployContract",
			Help: "use: \tdeployContract -contract -account -client -gasLimit\n\t\t\t\tdescription: Deploys specified contract instance to connected client",
			Func: func(c *ishell.Context) {
				flagSet := flag.NewFlagSet("deployContract", flag.ContinueOnError)
				contractName := flagSet.String("contract", "", "Name of compiled contract to deploy")
				accountName := flagSet.String("account", "", "Name of account to deploy from")
				clientName := flagSet.String("client", "", "Name of client to deploy to")
				limit := flagSet.String("gasLimit", "", "Gas limit provided for the deployment transaction")

				if err := flagSet.Parse(c.Args); err == nil {
					if flagSet.NFlag() != 4 {
						flagSet.Usage()
						return
					}

					if _, ok := session.Contracts[*contractName]; !ok {
						c.Println(fmt.Sprintf("Contract %s not recognised. Please use addContractInstance to add a new contract or specify a correct contract name.", *contractName))
						return
					}
					if _, ok := session.Accounts[*accountName]; !ok {
						c.Println(fmt.Sprintf("Account %s not recognised. Please use addAccount to add a new account or specify a correct account name.", *accountName))
						return
					}

					if _, ok := session.Networks[*clientName]; !ok {
						c.Println(fmt.Sprintf("Client %s not recognised. Please use addClient to add a new connected client first or specify a correct Client name.", *clientName))
						return
					}

					gasLimit, err := strconv.ParseUint(*limit, 10, 64)
					if err != nil {
						c.Println(err)
						return
					}

					contract := session.Contracts[*contractName]
					client := session.Networks[*clientName]
					account := session.Accounts[*accountName]

					constructorInputs, err := parseMethodArguments(c, contract.Abi, "")
					if err != nil {
						c.Printf("Error parsing constructor parameters: %s\n", err)
						return
					}

					payload, err := contracts.CreateTransactionPayload(contract, "", constructorInputs...)
					if err != nil {
						c.Printf("Error compiling tx payload: %s\n", err)
						return
					}

					tx, err := contracts.SendTransaction(
						session.Context,
						client.Client,
						account.Key.PrivateKey,
						nil,
						payload,
						nil,
						gasLimit,
					)
					if err != nil {
						c.Println(err)
						return
					}

					c.Println("Waiting for contract to be deployed")
					addr, err := bind.WaitDeployed(session.Context, client.Client, tx)
					if err != nil {
						c.Println(err)
						return
					}
					c.Printf("Deployed contract at: %s\n", addr.String())
				} else {
					c.Println(err.Error())
				}

				c.Println("===============================================================")
			},
		},
		{
			Name: "linkAndDeployContract",
			Help: "use: \tlinkAndDeployContract -contract -account -client -gasLimit -libraries\n\t\t\t\tdescription: Deploys specified contract instance while linking to existing deployed libraries",
			Func: func(c *ishell.Context) {
				flagSet := flag.NewFlagSet("linkAndDeployContract", flag.ContinueOnError)
				contractName := flagSet.String("contract", "", "Name of compiled contract to deploy")
				accountName := flagSet.String("account", "", "Name of account to deploy from")
				clientName := flagSet.String("client", "", "Name of client to deploy to")
				limit := flagSet.String("gasLimit", "", "Gas limit provided for the deployment transaction")
				librariesToLink := flagSet.String("libraries", "", "Comma-separated list of libraries to link for compilation. Format: -libraries <LibraryName>:<LibraryAddress>,<LibraryName>:<LibraryAddress>,... e.g. RLP:0x12345678...")

				if err := flagSet.Parse(c.Args); err == nil {
					if flagSet.NFlag() != 4 {
						flagSet.Usage()
						return
					}

					if _, ok := session.Contracts[*contractName]; !ok {
						c.Println(fmt.Sprintf("Contract %s not recognised. Please use addContractInstance to add a new contract or specify a correct contract name.", *contractName))
						return
					}
					if _, ok := session.Accounts[*accountName]; !ok {
						c.Println(fmt.Sprintf("Account %s not recognised. Please use addAccount to add a new account or specify a correct account name.", *accountName))
						return
					}

					if _, ok := session.Networks[*clientName]; !ok {
						c.Println(fmt.Sprintf("Client %s not recognised. Please use addClient to add a new connected client first or specify a correct Client name.", *clientName))
						return
					}

					gasLimit, err := strconv.ParseUint(*limit, 10, 64)
					if err != nil {
						c.Println(err)
						return
					}

					libraryList := strings.Split(*librariesToLink, ",")
					libraries := make(map[string]common.Address)

					for _, lib := range libraryList {
						library := strings.Split(lib, ":")
						name := library[0]
						address := common.HexToAddress(library[1])
						libraries[name] = address
					}

					contract, err := core.AddCompilerLinkAndCompileContract(session, session.Contracts[*contractName].Path, libraries)
					if err != nil {
						c.Println(err)
						return
					}

					client := session.Networks[*clientName]
					account := session.Accounts[*accountName]

					constructorInputs, err := parseMethodArguments(c, contract.Abi, "")
					if err != nil {
						c.Printf("Error parsing constructor parameters: %s\n", err)
						return
					}

					payload, err := contracts.CreateTransactionPayload(contract, "", constructorInputs...)
					if err != nil {
						c.Printf("Error compiling tx payload: %s\n", err)
						return
					}

					tx, err := contracts.SendTransaction(
						session.Context,
						client.Client,
						account.Key.PrivateKey,
						nil,
						payload,
						nil,
						gasLimit,
					)
					if err != nil {
						c.Println(err)
						return
					}

					c.Println("Waiting for contract to be deployed")
					addr, err := bind.WaitDeployed(session.Context, client.Client, tx)
					if err != nil {
						c.Println(err)
						return
					}
					c.Printf("Deployed contract at: %s\n", addr.String())
				} else {
					c.Println(err.Error())
				}

				c.Println("===============================================================")
			},
		},
		{
			Name: "functionTransaction",
			Help: "use: \tfunctionTransaction -contract -address -function -account -client -ether -gasLimit \n\t\t\t\tdescription: Calls a contract function as a transaction.",
			Func: func(c *ishell.Context) {
				flagSet := flag.NewFlagSet("functionTransaction", flag.ContinueOnError)
				contractName := flagSet.String("contract", "", "Name of contract of which a function will be called")
				contractAddress := flagSet.String("address", "", "Address of deployed contract")
				functionName := flagSet.String("function", "", "Name of function to be called")
				accountName := flagSet.String("account", "", "Name of account to sign transaction with")
				clientName := flagSet.String("client", "", "Name of client to transaction will be sent to")
				eth := flagSet.String("ether", "", "Amount of Ether to be sent with the function call")
				limit := flagSet.String("gasLimit", "", "Gas limit provided for the deployment transaction")

				if err := flagSet.Parse(c.Args); err == nil {
					if flagSet.NFlag() != 7 {
						flagSet.Usage()
						return
					}

					if _, ok := session.Contracts[*contractName]; !ok {
						c.Println(fmt.Sprintf("Contract %s not recognised. Please use addContractInstance to add a new contract or specify a correct contract name.", *contractName))
						return
					}
					if _, ok := session.Accounts[*accountName]; !ok {
						c.Println(fmt.Sprintf("Account %s not recognised. Please use addAccount to add a new account or specify a correct account name.", *accountName))
						return
					}

					if _, ok := session.Networks[*clientName]; !ok {
						c.Println(fmt.Sprintf("Client %s not recognised. Please use addClient to add a new connected client first or specify a correct Client name.", *clientName))
						return
					}

					contract := session.Contracts[*contractName]
					client := session.Networks[*clientName]
					account := session.Accounts[*accountName]

					ether := new(big.Int)
					ether, ok := ether.SetString(*eth, 10)
					if !ok {
						c.Println(fmt.Sprintf("Please enter an integer for -ether"))
						return
					}

					gasLimit, err := strconv.ParseUint(*limit, 10, 64)
					if err != nil {
						c.Println(err)
						return
					}
					contractDeployedAddress := common.HexToAddress(*contractAddress)

					if contract.Abi.Methods[*functionName].Name == "" {
						c.Printf("Method name \"%s\" not found for contract \"%s\"\n", *functionName, *contractName)
						return
					}

					inputs, err := parseMethodArguments(c, contract.Abi, *functionName)
					if err != nil {
						c.Printf("Error parsing parameters: %s\n", err)
						return
					}

					tx, err := contracts.FunctionCallTransaction(
						session.Context,
						client.Client,
						account.Key.PrivateKey,
						contract,
						contractDeployedAddress,
						ether,
						gasLimit,
						*functionName,
						inputs...,
					)
					if err != nil {
						c.Println(err)
						return
					} else {
						c.Println("Waiting for transaction to be mined...")
						receipt, err := bind.WaitMined(session.Context, client.Client, tx)
						if err != nil {
							c.Println(err)
							return
						}
						c.Printf("Transaction hash: %s\n", receipt.TxHash.String())
					}
				} else {
					c.Println(err.Error())
				}

				c.Println("===============================================================")
			},
		},
		{
			Name: "functionCall",
			Help: "use: \tfunctionCall [contract name] [function name] [from account name] [deployed contract address] \n\t\t\t\tdescription: Calls a contract function returning result without mutating state.",
			Func: func(c *ishell.Context) {
				flagSet := flag.NewFlagSet("functionCall", flag.ContinueOnError)
				contractName := flagSet.String("contract", "", "Name of contract of which a function will be called")
				contractAddress := flagSet.String("address", "", "Address of deployed contract")
				functionName := flagSet.String("function", "", "Name of function to be called")
				accountName := flagSet.String("account", "", "Name of account to call function from")
				clientName := flagSet.String("client", "", "Name of client to send request to")

				if err := flagSet.Parse(c.Args); err == nil {
					if flagSet.NFlag() != 5 {
						flagSet.Usage()
						return
					}

					if _, ok := session.Contracts[*contractName]; !ok {
						c.Println(fmt.Sprintf("Contract %s not recognised. Please use addContractInstance to add a new contract or specify a correct contract name.", *contractName))
						return
					}
					if _, ok := session.Accounts[*accountName]; !ok {
						c.Println(fmt.Sprintf("Account %s not recognised. Please use addAccount to add a new account or specify a correct account name.", *accountName))
						return
					}

					if _, ok := session.Networks[*clientName]; !ok {
						c.Println(fmt.Sprintf("Client %s not recognised. Please use addClient to add a new connected client first or specify a correct Client name.", *clientName))
						return
					}

					contract := session.Contracts[*contractName]
					client := session.Networks[*clientName]
					account := session.Accounts[*accountName]

					contractDeployedAddress := common.HexToAddress(*contractAddress)

					if contract.Abi.Methods[*functionName].Name == "" {
						c.Printf("Method name \"%s\" not found for contract \"%s\"\n", *functionName, *contractName)
						return
					}

					inputs, err := parseMethodArguments(c, contract.Abi, *functionName)
					if err != nil {
						c.Printf("Error parsing parameters: %s\n", err)
						return
					}

					var out interface{}

					out, err = contracts.CallContract(
						session.Context,
						client.Client,
						contract,
						account.Key.Address,
						contractDeployedAddress,
						*functionName,
						out,
						inputs...,
					)
					if err != nil {
						c.Println(err)
						return
					} else {
						c.Printf("Result: %s\n", out)
					}
				} else {
					c.Println(err.Error())
				}

				c.Println("===============================================================")
			},
		},
		{
			Name: "getTransactionByHash",
			Help: "use: \tgetTransactionByHash -client -hash\n\t\t\t\tdescription: Returns transaction specified by hash from connected Client or specified endpoint",
			Func: func(c *ishell.Context) {
				flagSet := flag.NewFlagSet("getTransactionByHash", flag.ContinueOnError)
				clientName := flagSet.String("client", "", "Name of client to send request to")
				hash := flagSet.String("hash", "", "Transaction hash")

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

					_, tx, err := utils.GetTransactionByHash(client.Client, common.HexToHash(*hash))
					if err != nil {
						c.Println(err)
						return
					}
					c.Printf("Transaction: \n%s\n", tx)
				} else {
					c.Println(err.Error())
				}
				c.Println("===============================================================")
			},
		},
		{
			Name: "getBlockByNumber",
			Help: "use: \tgetBlockByNumber -height -client\n\t\t\t\tdescription: Returns block header specified by height",
			Func: func(c *ishell.Context) {
				flagSet := flag.NewFlagSet("getBlockByNumber", flag.ContinueOnError)
				clientName := flagSet.String("client", "", "Name of client to send request to")
				number := flagSet.Int64("height", 0, "Block height")

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

					_, block, err := utils.GetBlockHeaderByNumber(client.Client, big.NewInt(*number))
					if err != nil {
						c.Println(err)
						return
					}
					c.Printf("Block: %s\n", block)
				} else {
					c.Println(err.Error())
				}
				c.Println("===============================================================")
			},
		},
		{
			Name: "getBlockByHash",
			Help: "use: \tgetBlockByHash -client -hash \n\t\t\t\tdescription: Returns block header specified by hash",
			Func: func(c *ishell.Context) {
				flagSet := flag.NewFlagSet("getBlockByHash", flag.ContinueOnError)
				clientName := flagSet.String("client", "", "Name of client to send request to")
				hash := flagSet.String("hash", "", "Block hash")

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

					_, block, err := utils.GetBlockHeaderByHash(client.Client, common.HexToHash(*hash))
					if err != nil {
						c.Println(err)
						return
					}
					c.Printf("Block: %s\n", block)
				} else {
					c.Println(err.Error())
				}
				c.Println("===============================================================")
			},
		},
		{
			Name: "getRLPEncodedBlockByHash",
			Help: "use: \tgetRLPEncodedBlockByHash -client -hash \n\t\t\t\tdescription: Returns RLP-encoded block header specified by hash",
			Func: func(c *ishell.Context) {
				flagSet := flag.NewFlagSet("getRLPEncodedBlockByHash", flag.ContinueOnError)
				clientName := flagSet.String("client", "", "Name of client to send request to")
				hash := flagSet.String("hash", "", "Block hash")

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
					encodedBlock, err := utils.RlpEncodeBlock(block)
					if err != nil {
						c.Println(err)
						return
					}
					c.Printf("Encoded Block: \n%+x\n", encodedBlock)
				} else {
					c.Println(err.Error())
				}
				c.Println("===============================================================")
			},
		},
		{
			Name: "getRLPEncodedBlockByNumber",
			Help: "use: \tgetRLPEncodedBlockByNumber -client -height\n\t\t\t\tdescription: Returns RLP-encoded block header specified by number",
			Func: func(c *ishell.Context) {
				flagSet := flag.NewFlagSet("getRLPEncodedBlockByNumber", flag.ContinueOnError)
				clientName := flagSet.String("client", "", "Name of client to send request to")
				number := flagSet.Int64("height", 0, "Block height")

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
					encodedBlock, err := utils.RlpEncodeBlock(block)
					if err != nil {
						c.Println(err)
						return
					}
					c.Printf("Encoded Block: \n%+x\n", encodedBlock)
				} else {
					c.Println(err.Error())
				}

				c.Println("===============================================================")
			},
		},
		{
			Name: "getProof",
			Help: "use: \tgetProof -client -txhash \n\t\t\t\tdescription: Returns an RLP-encoded set of merkle proofs of a specific transaction and its receipt in a block",
			Func: func(c *ishell.Context) {
				flagSet := flag.NewFlagSet("getProof", flag.ContinueOnError)
				clientName := flagSet.String("client", "", "Name of client to send request to")
				hash := flagSet.String("txhash", "", "Transaction hash of transaction to be proven")

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

					data, err := utils.FetchProofData(client.Client, common.HexToHash(*hash))
					if err != nil {
						c.Println(err.Error())
						return
					}

					proof, err := utils.GenerateIonProof(*data)
					if err != nil {
						c.Println(err.Error())
						return
					}

					c.Printf("Proof: \n0x%x\n", proof)
				} else {
					c.Println(err.Error())
				}
				c.Println("===============================================================")
			},
		},
	}
}

func parseMethodArguments(c *ishell.Context, abiStruct *abi.ABI, methodName string) (args []interface{}, err error) {
	var inputArguments abi.Arguments
	if methodName != "" {
		inputArguments = abiStruct.Methods[methodName].Inputs
	} else {
		inputArguments = abiStruct.Constructor.Inputs
	}

	c.ShowPrompt(false)
	defer c.ShowPrompt(true)

	for i := 0; i < len(inputArguments); {
		argument := inputArguments[i]
		c.Printf("Enter input data for parameter %s:\n", argument.Name)

		input, err := c.ReadLineErr()
		if err != nil {
			return nil, err
		}

		arg, err := utils.ApplySolidityType(input, argument.Type)
		if err != nil {
			if err == utils.NotArrayFormatError {
				c.Println("Input error:", err.Error())
				continue
			}
			return nil, err
		}
		c.Println(arg)
		args = append(args, arg)
		i++
	}

	return
}
