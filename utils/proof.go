// Copyright (c) 2018-2020 Clearmatics Technologies Ltd
package utils

import (
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/ethdb/memorydb"
	"github.com/ethereum/go-ethereum/rlp"
	"github.com/ethereum/go-ethereum/trie"
)

type IonProofData struct {
	Transaction       *types.Transaction
	TransactionIndex  byte
	BlockTransactions []*types.Transaction
	BlockReceipts     []*types.Receipt
}

func FetchProofData(client *ethclient.Client, txHash common.Hash) (*IonProofData, error) {
	tx, _, err := GetTransactionByHash(client, txHash)
	if err != nil {
		return nil, err
	}

	blockNum, err := BlockNumberByTransactionHash(client, txHash)
	if err != nil {
		return nil, err
	}

	block, _, err := GetBlockByNumber(client, blockNum)
	if err != nil {
		return nil, err
	}

	var txIndex byte
	txs := block.Transactions()

	// Calculate transaction index)
	for i := 0; i < len(txs); i++ {
		if txHash == txs[i].Hash() {
			txIndex = byte(i)
		}
	}

	receipts := GetAllReceiptsFromBlock(client, block)

	return &IonProofData{
		Transaction:       tx,
		TransactionIndex:  txIndex,
		BlockTransactions: txs,
		BlockReceipts:     receipts,
	}, nil
}

func GenerateIonProof(data IonProofData) ([]byte, error) {
	txTrie, err := TxTrie(data.BlockTransactions)
	if err != nil {
		return nil, err
	}

	receiptTrie, err := ReceiptTrie(data.BlockReceipts)
	if err != nil {
		return nil, err
	}

	txPath := []byte{data.TransactionIndex}
	txRLP, _ := rlp.EncodeToBytes(data.Transaction)

	txProof, err := createProofPath(txTrie, txPath[:])
	if err != nil {
		return nil, err
	}

	receiptRLP, _ := rlp.EncodeToBytes(data.BlockReceipts[txPath[0]])
	receiptProof, err := createProofPath(receiptTrie, txPath[:])
	if err != nil {
		return nil, err
	}

	var decodedTx, decodedTxProof, decodedReceipt, decodedReceiptProof []interface{}

	err = rlp.DecodeBytes(txRLP, &decodedTx)
	if err != nil {
		return []byte{}, err
	}

	err = rlp.DecodeBytes(txProof, &decodedTxProof)
	if err != nil {
		return []byte{}, err
	}

	err = rlp.DecodeBytes(receiptRLP, &decodedReceipt)
	if err != nil {
		return []byte{}, err
	}

	fmt.Printf("RECEIPT RLP: %x\n", receiptRLP)

	err = rlp.DecodeBytes(receiptProof, &decodedReceiptProof)
	if err != nil {
		return []byte{}, err
	}

	proof := make([]interface{}, 0)
	proof = append(proof, txPath, decodedTx, decodedTxProof, decodedReceipt, decodedReceiptProof)

	return rlp.EncodeToBytes(proof)
}

// Proof creates an array of the proof path ordered
func createProofPath(trie *trie.Trie, path []byte) ([]byte, error) {
	proof, err := generateTrieProof(trie, path)
	if err != nil {
		return []byte{}, err
	}
	proofRLP, err := rlp.EncodeToBytes(proof)
	if err != nil {
		return []byte{}, err
	}

	return proofRLP, nil
}

func generateTrieProof(trie *trie.Trie, path []byte) ([]interface{}, error) {
	proof := memorydb.New()
	err := trie.Prove(path, 0, proof)
	if err != nil {
		return []interface{}{}, err
	}

	var proofArr []interface{}
	for nodeIt := trie.NodeIterator(nil); nodeIt.Next(true); {
		if val, err := proof.Get(nodeIt.Hash().Bytes()); val != nil && err == nil {
			var decodedVal interface{}
			err = rlp.DecodeBytes(val, &decodedVal)
			if err != nil {
				return []interface{}{}, err
			}
			proofArr = append(proofArr, decodedVal)
		}
	}

	return proofArr, nil
}
