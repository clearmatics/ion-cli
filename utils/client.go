// Copyright (c) 2018 Clearmatics Technologies Ltd
package utils

import (
	"context"
	"encoding/json"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

// Client gets client or fails if no connection
func Client(url string) (*ethclient.Client, error) {
	return ethclient.Dial(url)
}

func GetBlockHeaderByNumber(client *ethclient.Client, blockNum *big.Int) (*types.Header, []byte, error) {
	block, err := client.HeaderByNumber(context.Background(), blockNum)
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

func GetBlockHeaderByHash(client *ethclient.Client, blockHash common.Hash) (*types.Header, []byte, error) {
	block, err := client.HeaderByHash(context.Background(), blockHash)
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

func GetBlockByNumber(client *ethclient.Client, blockNum *big.Int) (*types.Block, []byte, error) {
	block, err := client.BlockByNumber(context.Background(), blockNum)
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

func GetBlockByHash(client *ethclient.Client, blockHash common.Hash) (*types.Block, []byte, error) {
	block, err := client.BlockByHash(context.Background(), blockHash)
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

func GetTransactionByHash(client *ethclient.Client, txHash common.Hash) (*types.Transaction, []byte, error) {
	tx, _, err := client.TransactionByHash(context.Background(), txHash)
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

// BlockNumberByTransactionHash gets a block number by a transaction hash in that block
func BlockNumberByTransactionHash(client *ethclient.Client, txHash common.Hash) (*big.Int, error) {
	receipt, err := client.TransactionReceipt(context.Background(), txHash)
	if err != nil {
		return nil, err
	}

	return receipt.BlockNumber, nil
}

// BlockHashByTransactionHash gets a block hash by a transaction hash in that block
func BlockHashByTransactionHash(client *ethclient.Client, txHash common.Hash) (common.Hash, error) {
	receipt, err := client.TransactionReceipt(context.Background(), txHash)
	if err != nil {
		return common.Hash{}, err
	}

	return receipt.BlockHash, nil
}

// GetAllReceiptsFromBlock get the receipts for all the transactions in a block
func GetAllReceiptsFromBlock(client *ethclient.Client, block *types.Block) types.Receipts {
	var receiptsArr []*types.Receipt
	for _, tx := range block.Transactions() {
		receipt, err := client.TransactionReceipt(context.Background(), tx.Hash())
		if err != nil {
			log.Fatal("TransactionReceipt ERROR:", err)
		}
		receiptsArr = append(receiptsArr, receipt)
	}
	return receiptsArr
}
