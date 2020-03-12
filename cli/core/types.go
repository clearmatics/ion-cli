package core

import (
	"context"
	"github.com/clearmatics/ion-cli/config"
	contract "github.com/clearmatics/ion-cli/contracts"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
)

type Session struct {
	Context   context.Context
	Networks  map[string]*EthClient
	Contracts map[string]*contract.ContractInstance
	Accounts  map[string]*config.Account
	Compilers map[string]string
}

type EthClient struct {
	Client    *ethclient.Client
	RpcClient *rpc.Client
	Url       string
}
