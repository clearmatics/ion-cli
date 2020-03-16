package backend

import (
	"context"
	"encoding/json"
	"github.com/clearmatics/ion-cli/utils"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"math/big"
)

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

func (eth *EthClient) GetTransactionByHash(hash string) (*types.Transaction, []byte, error) {
	txHash := common.HexToHash(hash)

	tx, _, err := eth.client.TransactionByHash(context.Background(), txHash)
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

/*
func GetProof(eth *EthClient, transactionHash string) ([]byte, error) {
	// Get the transaction hash
	bytesTxHash := common.HexToHash(transactionHash)

	// Generate the proof
	proof, err := utils.GenerateProof(
		context.Background(),
		eth.RpcClient,
		bytesTxHash,
	)

	if err != nil {
		return nil, err
	}

	//fmt.Printf( "Path:           0x%x\n" +
	//            "TxValue:        0x%x\n" +
	//            "TxNodes:        0x%x\n" +
	//            "ReceiptValue:   0x%x\n" +
	//            "ReceiptNodes:   0x%x\n", txPath, txValue, txNodes, receiptValue, receiptNodes)

	return proof, nil
}*/