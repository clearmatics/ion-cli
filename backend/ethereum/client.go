package ethereum

import (
	"context"
	"encoding/json"
	"github.com/clearmatics/ion-cli/utils"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/ethereum/go-ethereum/ethclient"
	"math/big"
)

type EthClient struct {
	Client    *ethclient.Client
	RpcClient *rpc.Client
	Url       string
}

func GetClient(url string) (*EthClient, error) {
	rpc := utils.ClientRPC(url)
	eth := ethclient.NewClient(rpc)

	client := EthClient{Client: eth, RpcClient: rpc, Url: url}

	_, _, err := client.GetBlockByNumber("0")

	return &client, err
}

func (eth *EthClient) GetBlockByNumber(number string) (block *types.Header, b []byte, err error) {

	if number == "latest" {
		err = eth.RpcClient.Call(&block, "eth_getBlockByNumber", "latest", false)
		if err != nil {
			return nil, nil, err
		}
	} else {
		blockNum := new(big.Int)
		blockNum.SetString(number, 10)

		block, err = eth.Client.HeaderByNumber(context.Background(), blockNum)
		if err != nil {
			return nil, nil, err
		}
	}

	// Marshal into a JSON
	b, err = json.MarshalIndent(block, "", " ")
	if err != nil {
		return nil, nil, err
	}
	return block, b, nil
}

func (eth *EthClient) GetBlockByHash(hash string) (*types.Header, []byte, error) {
	blockHash := common.HexToHash(hash)

	block, err := eth.Client.HeaderByHash(context.Background(), blockHash)
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

func (eth *EthClient) GetTransactionByHash(hash string) (*types.Transaction, []byte, error) {
	txHash := common.HexToHash(hash)

	tx, _, err := eth.Client.TransactionByHash(context.Background(), txHash)
	if err != nil {
		return nil, nil, err
	}
	// Marshal into a JSON
	t, err := json.MarshalIndent(tx, "", " ")
	if err != nil {
		return nil, nil, err
	}
	return tx, t, nil
}


//func (eth *EthClient) GetProof(transactionHash string) ([]byte, error) {
//	// Get the transaction hash
//	bytesTxHash := common.HexToHash(transactionHash)
//
//	// Generate the proof
//	proof, err := utils.GenerateProof(
//		context.Background(),
//		eth.rpcClient,
//		bytesTxHash,
//	)
//
//	if err != nil {
//		return nil, err
//	}
//
//	return proof, nil
//}