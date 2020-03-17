// Copyright (c) 2018 Clearmatics Technologies Ltd
package contracts

import (
	"context"
	"crypto/ecdsa"
	"encoding/json"
	"errors"
	"fmt"
	"math/big"
	"path/filepath"
	"strings"

	"github.com/Shirikatsu/go-ethereum/common/compiler"
	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
)

// ContractInstance is just an util type to output contract and address
type ContractInstance struct {
	Contract *compiler.Contract
	Abi      *abi.ABI
	Path     string
	Bytecode []byte
}

func GetContractBytecodeAndABI(c *compiler.Contract) ([]byte, *abi.ABI, error) {
	bytecode := common.FromHex(c.Code)

	abiStr, err := json.Marshal(c.Info.AbiDefinition)
	if err != nil {
		return []byte{}, nil, err
	}

	abiObj, err := abi.JSON(strings.NewReader(string(abiStr)))
	if err != nil {
		return []byte{}, nil, err
	}

	return bytecode, &abiObj, nil
}

func CreateTransactionPayload(contract *ContractInstance, methodName string, inputs ...interface{}) ([]byte, error) {
	payload, err := contract.Abi.Pack(methodName, inputs...)
	if err != nil {
		return nil, err
	}

	// Is constructor, therefore we are deploying a contract
	if methodName == "" {
		return append(contract.Bytecode, payload...), nil
	}

	return payload, nil
}

func SendTransaction(
	ctx context.Context,
	backend bind.ContractBackend,
	userKey *ecdsa.PrivateKey,
	to *common.Address,
	payload []byte,
	amount *big.Int,
	gasLimit uint64,
) (*types.Transaction, error) {
	userAddr := crypto.PubkeyToAddress(userKey.PublicKey)

	tx, err := createNewTransaction(ctx, backend, &userAddr, to, amount, gasLimit, payload)
	if err != nil {
		return nil, err
	}

	signedTx, err := signTransaction(tx, userKey)
	if err != nil {
		return nil, err
	}

	err = backend.SendTransaction(ctx, signedTx)
	if err != nil {
		return nil, err
	}

	return signedTx, nil
}

// FunctionCallTransaction execute function in contract
func FunctionCallTransaction(
	ctx context.Context,
	backend bind.ContractBackend,
	userKey *ecdsa.PrivateKey,
	contract *ContractInstance,
	to common.Address,
	amount *big.Int,
	gasLimit uint64,
	methodName string,
	args ...interface{},
) (*types.Transaction, error) {
	fmt.Println("Creating transaction payload...")
	payload, err := CreateTransactionPayload(contract, methodName, args...)
	if err != nil {
		errStr := fmt.Sprintf("error creating transaction payload: %s\n", err)
		return nil, errors.New(errStr)
	}

	return SendTransaction(ctx, backend, userKey, &to, payload, amount, gasLimit)
}

// CallContract without changing the state
func CallContract(
	ctx context.Context,
	client bind.ContractCaller,
	contract *ContractInstance,
	from, to common.Address,
	methodName string,
	out interface{},
	args ...interface{},
) (res interface{}, err error) {
	payload, err := CreateTransactionPayload(contract, methodName, args...)
	if err != nil {
		errStr := fmt.Sprintf("error creating transaction payload: %s\n", err)
		return nil, errors.New(errStr)
	}

	callMsg := ethereum.CallMsg{From: from, To: &to, Data: payload}
	output, err := client.CallContract(ctx, callMsg, nil)
	if err != nil {
		errStr := fmt.Sprintf("error calling contract function: %s\n", err)
		return nil, errors.New(errStr)
	}

	err = contract.Abi.Unpack(out, methodName, output)
	if err != nil {
		errStr := fmt.Sprintf("error unpacking call message result: %s\n", err)
		return nil, errors.New(errStr)
	}

	return out, nil
}

func createNewTransaction(
	ctx context.Context,
	backend bind.ContractBackend,
	from, to *common.Address,
	amount *big.Int,
	gasLimit uint64,
	payload []byte,
) (*types.Transaction, error) {
	nonce, err := backend.PendingNonceAt(ctx, *from) // uint64(0)
	if err != nil {
		return nil, err
	}

	gasPrice, err := backend.SuggestGasPrice(ctx) //new(big.Int)
	if err != nil {
		return nil, err
	}

	// create contract transaction NewContractCreation is the same has NewTransaction with `to` == nil
	// tx := types.NewTransaction(nonce, nil, amount, gasLimit, gasPrice, payload)
	var tx *types.Transaction
	if to == nil {
		tx = types.NewContractCreation(nonce, amount, gasLimit, gasPrice, payload)
	} else {
		tx = types.NewTransaction(nonce, *to, amount, gasLimit, gasPrice, payload)
	}

	return tx, nil
}

// method created just to easily sign a tranasaction
func signTransaction(tx *types.Transaction, userKey *ecdsa.PrivateKey) (*types.Transaction, error) {
	signer := types.HomesteadSigner{} // this functions makes it easier to change signer if needed

	signedTx, err := types.SignTx(tx, signer, userKey)
	if err != nil {
		return nil, err
	}

	return signedTx, nil
}

func CompileContractAt(contractPath string, solc string) (compiledContract *compiler.Contract, err error) {
	absPath, err := filepath.Abs(contractPath)
	if err != nil {
		return nil, err
	}
	path := strings.Split(absPath, "/")
	contractFolder := path[len(path)-2]
	contractFile := path[len(path)-1]

	i := strings.Index(absPath, contractFolder)
	remapping := fmt.Sprintf("../=%s", absPath[:i])

	contract, err := compiler.CompileSolidity(solc, []string{remapping}, absPath)
	if err != nil {
		return nil, err
	}

	for key := range contract {
		if strings.Contains(key, contractFile) {
			return contract[key], nil
		}
	}

	return nil, errors.New("compiled contract contains no data")
}

func CompileContractWithLibraries(contractPath string, libraries map[string]common.Address, solc string) (compiledContract *compiler.Contract, err error) {
	absPath, err := filepath.Abs(contractPath)
	if err != nil {
		return nil, err
	}
	path := strings.Split(absPath, "/")
	contractFolder := path[len(path)-2]
	contractFile := path[len(path)-1]

	var args []string

	// Add libraries to args
	for name, address := range libraries {
		libraryArg := name + ":" + address.String()
		args = append(args, fmt.Sprintf("--libraries=%s", libraryArg))
	}

	// Add remapping values to args
	i := strings.Index(absPath, contractFolder)
	args = append(args, fmt.Sprintf("../=%s ", absPath[:i]))

	contract, err := compiler.CompileSolidity(solc, args, absPath)
	if err != nil {
		return nil, err
	}

	for key := range contract {
		if strings.Contains(key, contractFile) {
			return contract[key], nil
		}
	}

	return nil, errors.New("compiled contract contains no data")

}
