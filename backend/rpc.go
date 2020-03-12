package backend

import (
	"context"
	"encoding/json"
	"github.com/clearmatics/ion-cli/utils"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
	"math/big"
)

type EthClient struct {
	client    *ethclient.Client
	rpcClient *rpc.Client
	url       string
}

func GetClient(url string) (*EthClient, error) {
	rpc := utils.ClientRPC(url)
	eth := ethclient.NewClient(rpc)

	client := EthClient{client: eth, rpcClient: rpc, url: url}

	_, _, err := client.GetBlockByNumber("0")

	return &client, err
}

func (eth *EthClient) GetBlockByNumber(number string) (*types.Header, []byte, error) {
	// var blockHeader header
	blockNum := new(big.Int)
	blockNum.SetString(number, 10)

	block, err := eth.client.HeaderByNumber(context.Background(), blockNum)
	if err != nil {
		return nil, nil, err
	}
	// Marshal into a JSON
	b, err := json.MarshalIndent(block, "", " ")
	if err != nil {
		return nil, nil, err
	}
	return block, b, nil
}

func (eth *EthClient) GetBlockByHash(hash string) (*types.Header, []byte, error) {
	blockHash := common.HexToHash(hash)

	block, err := eth.client.HeaderByHash(context.Background(), blockHash)
	if err != nil {
		return nil, nil, err
	}
	// Marshal into a JSON
	b, err := json.MarshalIndent(block, "", " ")
	if err != nil {
		return nil, nil, err
	}
	return block, b, nil
}