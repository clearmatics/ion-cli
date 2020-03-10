package utils

import (
	contract "github.com/clearmatics/ion-cli/contracts"
	"github.com/ethereum/go-ethereum/common"
)

func CreateContractInstance(pathToContract string, solc string) (*contract.ContractInstance, error) {
	compiledContract, err := contract.CompileContractAt(pathToContract, solc)
	if err != nil {
		return nil, err
	}

	bytecode, abi, err := contract.GetContractBytecodeAndABI(compiledContract)
	if err != nil {
		return nil, err
	}

	return &contract.ContractInstance{Contract: compiledContract, Abi: abi, Path: pathToContract, Bytecode: bytecode}, nil
}

func CreateLinkedContractInstance(pathToContract string, libraries map[string]common.Address, solc string) (*contract.ContractInstance, error) {
	compiledContract, err := contract.CompileContractWithLibraries(pathToContract, libraries, solc)
	if err != nil {
		return nil, err
	}

	bytecode, abi, err := contract.GetContractBytecodeAndABI(compiledContract)
	if err != nil {
		return nil, err
	}

	return &contract.ContractInstance{Contract: compiledContract, Abi: abi, Path: pathToContract, Bytecode: bytecode}, nil
}
