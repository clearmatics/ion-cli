package cmd

import (
	_ "context"
	"encoding/hex"
	"errors"
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
	"github.com/ethereum/go-ethereum/ethclient"
	"math/big"
	"reflect"
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
					client, err := getClient(*uri)
					if err != nil {
						c.Println("Could not connect to client.\n")
						return
					}

					session.Networks[*name] = client
					c.Println("Connected!")
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

					compiledContract, err := core.CompileContract(session, *path)
					if err != nil {
						c.Println(err)
						return
					}

					fmt.Printf("Compiled contract bytecode: \n%s\n", compiledContract.BinStr)
					fmt.Printf("Compiled contract abistr: \n%s\n", compiledContract.AbiStr)
					fmt.Printf("Compiled contract abi: \n%s\n", compiledContract.Abi)

					session.Contracts[*name] = compiledContract
					c.Println("Added!")
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

					constructorInputs, err := parseMethodParameters(c, contract.Abi, "")
					if err != nil {
						c.Printf("Error parsing constructor parameters: %s\n", err)
						return
					}

					payload, err := contracts.CompilePayload(contract.BinStr, contract.AbiStr, constructorInputs...)
					if err != nil {
						c.Printf("Error compiling tx payload: %s\n", err)
						return
					}

					tx, err := contracts.DeployContract(
						session.Context,
						client.Client,
						account.Key.PrivateKey,
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
				}

				c.Println("===============================================================")
			},
		},
		{
			Name: "linkAndDeployContract",
			Help: "use: \tlinkAndDeployContract -contract -account -client -gasLimit -libraries\n\t\t\t\tdescription: Deploys specified contract instance while linking to existing deployed libraries",
			Func: func(c *ishell.Context) {
				flagSet := flag.NewFlagSet("deployContract", flag.ContinueOnError)
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

					contract, err := core.CompileLinkedContract(session, session.Contracts[*contractName].Path, libraries)
					if err != nil {
						c.Println(err)
						return
					}

					client := session.Networks[*clientName]
					account := session.Accounts[*accountName]

					constructorInputs, err := parseMethodParameters(c, contract.Abi, "")
					if err != nil {
						c.Printf("Error parsing constructor parameters: %s\n", err)
						return
					}

					payload, err := contracts.CompilePayload(contract.BinStr, contract.AbiStr, constructorInputs...)
					if err != nil {
						c.Printf("Error compiling tx payload: %s\n", err)
						return
					}

					tx, err := contracts.DeployContract(
						session.Context,
						client.Client,
						account.Key.PrivateKey,
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
				}

				c.Println("===============================================================")
			},
		},
		{
			Name: "transactionMessage",
			Help: "use: \ttransactionMessage -contract -address -function -account -client -ether -gasLimit \n\t\t\t\tdescription: Calls a contract function as a transaction.",
			Func: func(c *ishell.Context) {
				//if len(c.Args) != 6 {
				//	c.Println("Usage: \ttransactionMessage [contract name] [function name] [from account name] [deployed contract address] [amount] [gasLimit] \n")
				//} else {
				//	if ethClient == nil {
				//		c.Println("Please connect to a Client before invoking this function.\nUse \tconnectToClient [rpc Url] \n")
				//		return
				//	}
				//
				//	instance := contracts[c.Args[0]]
				//	methodName := c.Args[1]
				//	account := accounts[c.Args[2]]
				//	contractDeployedAddress := common.HexToAddress(c.Args[3])
				//
				//	if instance == nil {
				//		errStr := fmt.Sprintf("Contract instance %s not found.\nUse \taddContractInstance [name] [path/to/solidity/contract] \n", c.Args[0])
				//		c.Println(errStr)
				//		return
				//	}
				//	if account == nil {
				//		errStr := fmt.Sprintf("Account %s not found.\nUse \taddAccount [name] [path/to/keystore]\n", c.Args[2])
				//		c.Println(errStr)
				//		return
				//	}
				//
				//	amount := new(big.Int)
				//	amount, ok := amount.SetString(c.Args[4], 10)
				//	if !ok {
				//		c.Err(errors.New("Please enter an integer for <amount>"))
				//	}
				//	gasLimit, err := strconv.ParseUint(c.Args[5], 10, 64)
				//	if err != nil {
				//		c.Err(errors.New("Please enter an integer for <gasLimit>"))
				//	}
				//
				//	if instance.Abi.Methods[methodName].Name == "" {
				//		c.Printf("Method name \"%s\" not found for contract \"%s\"\n", methodName, c.Args[0])
				//		return
				//	}
				//
				//	inputs, err := parseMethodParameters(c, instance.Abi, methodName)
				//	if err != nil {
				//		c.Printf("Error parsing parameters: %s\n", err)
				//		return
				//	}
				//
				//	tx, err := contract.TransactionContract(
				//		ctx,
				//		ethClient.client,
				//		account.Key.PrivateKey,
				//		instance.Contract,
				//		contractDeployedAddress,
				//		amount,
				//		gasLimit,
				//		c.Args[1],
				//		inputs...,
				//	)
				//	if err != nil {
				//		c.Println(err)
				//		return
				//	} else {
				//		c.Println("Waiting for transaction to be mined...")
				//		receipt, err := bind.WaitMined(ctx, ethClient.client, tx)
				//		if err != nil {
				//			c.Println(err)
				//			return
				//		}
				//		c.Printf("Transaction hash: %s\n", receipt.TxHash.String())
				//	}
				//}

				flagSet := flag.NewFlagSet("deployContract", flag.ContinueOnError)
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

					inputs, err := parseMethodParameters(c, contract.Abi, *functionName)
					if err != nil {
						c.Printf("Error parsing parameters: %s\n", err)
						return
					}

					tx, err := contracts.FunctionCallTransaction(
						session.Context,
						client.Client,
						account.Key.PrivateKey,
						contract.Contract,
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
				}

				c.Println("===============================================================")
			},
		},
	}
}

func getClient(url string) (*core.EthClient, error) {
	rpc := utils.ClientRPC(url)
	eth := ethclient.NewClient(rpc)

	client := core.EthClient{eth, rpc, url}

	_, _, err := core.GetBlockByNumber(&client, "0")

	return &client, err
}

func parseMethodParameters(c *ishell.Context, abiStruct *abi.ABI, methodName string) (args []interface{}, err error) {
	var inputParameters abi.Arguments
	if methodName != "" {
		inputParameters = abiStruct.Methods[methodName].Inputs
	} else {
		inputParameters = abiStruct.Constructor.Inputs
	}

	c.ShowPrompt(false)
	defer c.ShowPrompt(true)

	for i := 0; i < len(inputParameters); i++ {
		argument := inputParameters[i]
		c.Printf("Enter input data for parameter %s:\n", argument.Name)

		input := c.ReadLine()

		// bytes = []byte{} argument type = slice, no element, type equates to []uint8
		// byte[] = [][1]byte{} argument type = slice, element type = array, type equates to [][1]uint8
		// byte = bytes1
		// bytesn = [n]byte{} 0 < n < 33, argument type = array, no element, type equates to [n]uint8
		// bytesn[] = [][n]byte{} argument type = slice, element type = array, type equares to [][n]uint8
		// bytesn[m] = [m][n]byte{} argument type = array, element type = array, type equates to [m][n]uint8
		// Many annoying cases of byte arrays

		if argument.Type.Kind == reflect.Array || argument.Type.Kind == reflect.Slice {
			c.Println("Argument is array\n")

			// One dimensional byte array
			// Accepts all byte arrays as hex string with pre-pended '0x' only
			if argument.Type.Elem == nil {
				if argument.Type.Type == reflect.TypeOf(common.Address{}) {
					// address solidity type
					item, err := utils.ConvertToType(input, &argument.Type)
					if err != nil {
						c.Err(err)
					}
					args = append(args, item)
					continue
				} else if argument.Type.Type == reflect.TypeOf([]byte{}) {
					// bytes solidity type
					bytes, err := hex.DecodeString(input[2:])
					if err != nil {
						c.Err(err)
					}
					args = append(args, bytes)
					continue
				} else {
					// Fixed byte array of size n; bytesn solidity type
					// Any submitted bytes longer than the expected size will be truncated

					bytes, err := hex.DecodeString(input[2:])
					if err != nil {
						c.Err(err)
					}

					// Fixed sized arrays can't be created with variables as size
					switch argument.Type.Size {
					case 1:
						var byteArray [1]byte
						copy(byteArray[:], bytes[:1])
						args = append(args, byteArray)
					case 2:
						var byteArray [2]byte
						copy(byteArray[:], bytes[:2])
						args = append(args, byteArray)
					case 3:
						var byteArray [3]byte
						copy(byteArray[:], bytes[:3])
						args = append(args, byteArray)
					case 4:
						var byteArray [4]byte
						copy(byteArray[:], bytes[:4])
						args = append(args, byteArray)
					case 5:
						var byteArray [5]byte
						copy(byteArray[:], bytes[:5])
						args = append(args, byteArray)
					case 6:
						var byteArray [6]byte
						copy(byteArray[:], bytes[:6])
						args = append(args, byteArray)
					case 7:
						var byteArray [7]byte
						copy(byteArray[:], bytes[:7])
						args = append(args, byteArray)
					case 8:
						var byteArray [8]byte
						copy(byteArray[:], bytes[:8])
						args = append(args, byteArray)
					case 9:
						var byteArray [9]byte
						copy(byteArray[:], bytes[:9])
						args = append(args, byteArray)
					case 10:
						var byteArray [10]byte
						copy(byteArray[:], bytes[:10])
						args = append(args, byteArray)
					case 11:
						var byteArray [11]byte
						copy(byteArray[:], bytes[:11])
						args = append(args, byteArray)
					case 12:
						var byteArray [12]byte
						copy(byteArray[:], bytes[:12])
						args = append(args, byteArray)
					case 13:
						var byteArray [13]byte
						copy(byteArray[:], bytes[:13])
						args = append(args, byteArray)
					case 14:
						var byteArray [14]byte
						copy(byteArray[:], bytes[:14])
						args = append(args, byteArray)
					case 15:
						var byteArray [15]byte
						copy(byteArray[:], bytes[:15])
						args = append(args, byteArray)
					case 16:
						var byteArray [16]byte
						copy(byteArray[:], bytes[:16])
						args = append(args, byteArray)
					case 17:
						var byteArray [17]byte
						copy(byteArray[:], bytes[:17])
						args = append(args, byteArray)
					case 18:
						var byteArray [18]byte
						copy(byteArray[:], bytes[:18])
						args = append(args, byteArray)
					case 19:
						var byteArray [19]byte
						copy(byteArray[:], bytes[:19])
						args = append(args, byteArray)
					case 20:
						var byteArray [20]byte
						copy(byteArray[:], bytes[:20])
						args = append(args, byteArray)
					case 21:
						var byteArray [21]byte
						copy(byteArray[:], bytes[:21])
						args = append(args, byteArray)
					case 22:
						var byteArray [22]byte
						copy(byteArray[:], bytes[:22])
						args = append(args, byteArray)
					case 23:
						var byteArray [23]byte
						copy(byteArray[:], bytes[:23])
						args = append(args, byteArray)
					case 24:
						var byteArray [24]byte
						copy(byteArray[:], bytes[:24])
						args = append(args, byteArray)
					case 25:
						var byteArray [25]byte
						copy(byteArray[:], bytes[:25])
						args = append(args, byteArray)
					case 26:
						var byteArray [26]byte
						copy(byteArray[:], bytes[:26])
						args = append(args, byteArray)
					case 27:
						var byteArray [27]byte
						copy(byteArray[:], bytes[:27])
						args = append(args, byteArray)
					case 28:
						var byteArray [28]byte
						copy(byteArray[:], bytes[:28])
						args = append(args, byteArray)
					case 29:
						var byteArray [29]byte
						copy(byteArray[:], bytes[:29])
						args = append(args, byteArray)
					case 30:
						var byteArray [30]byte
						copy(byteArray[:], bytes[:30])
						args = append(args, byteArray)
					case 31:
						var byteArray [31]byte
						copy(byteArray[:], bytes[:31])
						args = append(args, byteArray)
					case 32:
						var byteArray [32]byte
						copy(byteArray[:], bytes[:32])
						args = append(args, byteArray)
					default:
						errStr := fmt.Sprintf("Error parsing fixed size byte array. Array of size %d incompatible", argument.Type.Size)
						return nil, errors.New(errStr)
					}
					continue
				}

			}

			array := strings.Split(input, ",")
			argSize := argument.Type.Size
			size := len(array)
			if argSize != 0 {
				for size != argSize {
					c.Printf("Please enter %i comma-separated list of elements:\n", argSize)
					input = c.ReadLine()
					array = strings.Split(input, ",")
					size = len(array)
				}
			}

			size = len(array)

			elementType := argument.Type.Elem

			// Elements cannot be kind slice                                        only mean slice
			if elementType.Kind == reflect.Array && elementType.Type != reflect.TypeOf(common.Address{}) {
				// Is 2D byte array
				/* Nightmare to implement, have to account for:
				   * Slice of fixed byte arrays; bytes32[] in solidity for example, generally bytesn[]
				   * Fixed array of fixed byte arrays; bytes32[10] in solidity for example bytesn[m]
				   * Slice or fixed array of string; identical to above two cases as string in solidity is array of bytes

				   Since the upper bound of elements in an array in solidity is 2^256-1, and each fixed byte array
				   has a limit of bytes32 (bytes1, bytes2, ..., bytes31, bytes32), and Golang array creation takes
				   constant length values, we would have to paste the switch-case containing 1-32 fixed byte arrays
				   2^256-1 times to handle every possibility. Since arrays of arrays in seldom used, we have not
				   implemented it.
				*/

				return nil, errors.New("2D Arrays unsupported. Use \"bytes\" instead.")

				/*
				   slice := make([]interface{}, 0, size)
				   err = addFixedByteArrays(array, elementType.Size, slice)
				   if err != nil {
				       return nil, err
				   }
				   args = append(args, slice)
				   continue
				*/
			} else {
				switch elementType.Type {
				case reflect.TypeOf(bool(false)):
					convertedArray := make([]bool, 0, size)
					for _, item := range array {
						b, err := utils.ConvertToBool(item)
						if err != nil {
							return nil, err
						}
						convertedArray = append(convertedArray, b)
					}
					args = append(args, convertedArray)
				case reflect.TypeOf(int8(0)):
					convertedArray := make([]int8, 0, size)
					for _, item := range array {
						i, err := strconv.ParseInt(item, 10, 8)
						if err != nil {
							return nil, err
						}
						convertedArray = append(convertedArray, int8(i))
					}
					args = append(args, convertedArray)
				case reflect.TypeOf(int16(0)):
					convertedArray := make([]int16, 0, size)
					for _, item := range array {
						i, err := strconv.ParseInt(item, 10, 16)
						if err != nil {
							return nil, err
						}
						convertedArray = append(convertedArray, int16(i))
					}
					args = append(args, convertedArray)
				case reflect.TypeOf(int32(0)):
					convertedArray := make([]int32, 0, size)
					for _, item := range array {
						i, err := strconv.ParseInt(item, 10, 32)
						if err != nil {
							return nil, err
						}
						convertedArray = append(convertedArray, int32(i))
					}
					args = append(args, convertedArray)
				case reflect.TypeOf(int64(0)):
					convertedArray := make([]int64, 0, size)
					for _, item := range array {
						i, err := strconv.ParseInt(item, 10, 64)
						if err != nil {
							return nil, err
						}
						convertedArray = append(convertedArray, int64(i))
					}
					args = append(args, convertedArray)
				case reflect.TypeOf(uint8(0)):
					convertedArray := make([]uint8, 0, size)
					for _, item := range array {
						u, err := strconv.ParseUint(item, 10, 8)
						if err != nil {
							return nil, err
						}
						convertedArray = append(convertedArray, uint8(u))
					}
					args = append(args, convertedArray)
				case reflect.TypeOf(uint16(0)):
					convertedArray := make([]uint16, 0, size)
					for _, item := range array {
						u, err := strconv.ParseUint(item, 10, 16)
						if err != nil {
							return nil, err
						}
						convertedArray = append(convertedArray, uint16(u))
					}
					args = append(args, convertedArray)
				case reflect.TypeOf(uint32(0)):
					convertedArray := make([]uint32, 0, size)
					for _, item := range array {
						u, err := strconv.ParseUint(item, 10, 32)
						if err != nil {
							return nil, err
						}
						convertedArray = append(convertedArray, uint32(u))
					}
					args = append(args, convertedArray)
				case reflect.TypeOf(uint64(0)):
					convertedArray := make([]uint64, 0, size)
					for _, item := range array {
						u, err := strconv.ParseUint(item, 10, 64)
						if err != nil {
							return nil, err
						}
						convertedArray = append(convertedArray, uint64(u))
					}
					args = append(args, convertedArray)
				case reflect.TypeOf(&big.Int{}):
					convertedArray := make([]*big.Int, 0, size)
					for _, item := range array {
						newInt := new(big.Int)
						newInt, ok := newInt.SetString(item, 10)
						if !ok {
							return nil, errors.New("Could not convert string to big.int")
						}
						convertedArray = append(convertedArray, newInt)
					}
					args = append(args, convertedArray)
				case reflect.TypeOf(common.Address{}):
					convertedArray := make([]common.Address, 0, size)
					for _, item := range array {
						a := common.HexToAddress(item)
						convertedArray = append(convertedArray, a)
					}
					args = append(args, convertedArray)
				default:
					errStr := fmt.Sprintf("Type %s not found", elementType.Type)
					return nil, errors.New(errStr)
				}
			}
		} else {
			switch argument.Type.Kind {
			case reflect.String:
				args = append(args, input)
			case reflect.Bool:
				b, err := utils.ConvertToBool(input)
				if err != nil {
					return nil, err
				}
				args = append(args, b)
			case reflect.Int8:
				i, err := strconv.ParseInt(input, 10, 8)
				if err != nil {
					return nil, err
				}
				args = append(args, int8(i))
			case reflect.Int16:
				i, err := strconv.ParseInt(input, 10, 16)
				if err != nil {
					return nil, err
				}
				args = append(args, int16(i))
			case reflect.Int32:
				i, err := strconv.ParseInt(input, 10, 32)
				if err != nil {
					return nil, err
				}
				args = append(args, int32(i))
			case reflect.Int64:
				i, err := strconv.ParseInt(input, 10, 64)
				if err != nil {
					return nil, err
				}
				args = append(args, int64(i))
			case reflect.Uint8:
				u, err := strconv.ParseUint(input, 10, 8)
				if err != nil {
					return nil, err
				}
				args = append(args, uint8(u))
			case reflect.Uint16:
				u, err := strconv.ParseUint(input, 10, 16)
				if err != nil {
					return nil, err
				}
				args = append(args, uint16(u))
			case reflect.Uint32:
				u, err := strconv.ParseUint(input, 10, 32)
				if err != nil {
					return nil, err
				}
				args = append(args, uint32(u))
			case reflect.Uint64:
				u, err := strconv.ParseUint(input, 10, 64)
				if err != nil {
					return nil, err
				}
				args = append(args, uint64(u))
			case reflect.Ptr:
				newInt := new(big.Int)
				newInt, ok := newInt.SetString(input, 10)
				if !ok {
					return nil, errors.New("Could not convert string to big.int")
				}
				if err != nil {
					return nil, err
				}
				args = append(args, newInt)
			case reflect.Array:
				if argument.Type.Type == reflect.TypeOf(common.Address{}) {
					address := common.HexToAddress(input)
					args = append(args, address)
				} else {
					return nil, errors.New("Conversion failed. Item is array type, cannot parse")
				}
			default:
				errStr := fmt.Sprintf("Error, type not found: %s", argument.Type.Kind)
				return nil, errors.New(errStr)
			}
		}
	}

	return
}
