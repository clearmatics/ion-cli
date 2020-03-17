package core

import (
	"context"
	"fmt"
	"github.com/clearmatics/ion-cli/config"
	contract "github.com/clearmatics/ion-cli/contracts"
	"github.com/clearmatics/ion-cli/utils"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"math/big"
)

type Session struct {
	Context   context.Context
	Networks  map[string]*EthClient
	Contracts map[string]*contract.ContractInstance
	Accounts  map[string]*config.Account
	Compilers map[string]string
}

type EthClient struct {
	Client *ethclient.Client
	Url    string
}

func GetClient(url string) (*EthClient, error) {
	eth, err := utils.Client(url)
	if err != nil {
		return nil, err
	}

	client := EthClient{Client: eth, Url: url}

	_, _, err = utils.GetBlockHeaderByNumber(client.Client, big.NewInt(0))

	return &client, err
}

func InitSession() *Session {
	session := Session{}
	session.Context = context.Background()
	session.Networks = make(map[string]*EthClient)
	session.Contracts = make(map[string]*contract.ContractInstance)
	session.Accounts = make(map[string]*config.Account)
	session.Compilers = make(map[string]string)

	return &session
}

func (session *Session) Close() {
	session.RemoveAllCompilers()
}

func (session *Session) RemoveAllCompilers() {
	for _, compiler := range session.Compilers {
		utils.DestroyTempFile(compiler)
	}
}

func (session *Session) AddCompilerIfNotExists(version string) error {
	if _, ok := session.Compilers[version]; !ok {
		fmt.Printf("Compiler for version %s does not exist\n", version)
		solc, err := utils.GetSolidityCompilerVersion(version)
		if err != nil {
			return err
		}

		session.Compilers[version] = solc.Name()
	}
	return nil
}

func AddCompilerAndCompileContract(session *Session, pathToContract string) (*contract.ContractInstance, error) {
	fmt.Printf("Compiling contract %s...\n", pathToContract)
	version, err := utils.GetSolidityContractVersion(pathToContract)
	if err != nil {
		return nil, err
	}

	fmt.Printf("Contract version is %s\n", version)

	err = session.AddCompilerIfNotExists(version)
	if err != nil {
		return nil, err
	}

	solc := session.Compilers[version]
	return utils.CreateContractInstance(pathToContract, solc)
}

func AddCompilerLinkAndCompileContract(session *Session, pathToContract string, libraries map[string]common.Address) (*contract.ContractInstance, error) {
	fmt.Printf("Compiling contract %s...\n", pathToContract)
	version, err := utils.GetSolidityContractVersion(pathToContract)
	if err != nil {
		return nil, err
	}

	fmt.Printf("Contract version is %s\n", version)

	err = session.AddCompilerIfNotExists(version)
	if err != nil {
		return nil, err
	}

	solc := session.Compilers[version]
	return utils.CreateLinkedContractInstance(pathToContract, libraries, solc)
}
