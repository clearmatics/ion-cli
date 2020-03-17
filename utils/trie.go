// Copyright (c) 2018 Clearmatics Technologies Ltd
package utils

import (
	"log"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethdb/memorydb"
	"github.com/ethereum/go-ethereum/rlp"
	"github.com/ethereum/go-ethereum/trie"
)

// ReceiptTrie generate trie for receipts
// TODO: the argument should be of type interface so that this is a generic function
func ReceiptTrie(receipts []*types.Receipt) (*trie.Trie, error) {
	var receiptRLPidxArr, receiptRLPArr [][]byte
	for idx, receipt := range receipts {
		idxRLP, err := rlp.EncodeToBytes(uint(idx))
		if err != nil {
			return nil, err
		}
		txRLP, err := rlp.EncodeToBytes(receipt)
		if err != nil {
			return nil, err
		}

		receiptRLPidxArr = append(receiptRLPidxArr, idxRLP)
		receiptRLPArr = append(receiptRLPArr, txRLP)
	}

	trieObj := generateTrie(receiptRLPidxArr, receiptRLPArr)

	return trieObj, nil
}

// TxTrie generated Trie out of transaction array
// TODO: the argument should be of type interface so that this is a generic function
func TxTrie(transactions []*types.Transaction) (*trie.Trie, error) {
	var txRLPIdxArr, txRLPArr [][]byte
	for idx, tx := range transactions {
		idxRLP, err := rlp.EncodeToBytes(uint(idx))
		if err != nil {
			return nil, err
		}
		txRLP, err := rlp.EncodeToBytes(tx)
		if err != nil {
			return nil, err
		}

		txRLPIdxArr = append(txRLPIdxArr, idxRLP)
		txRLPArr = append(txRLPArr, txRLP)
	}

	trieObj := generateTrie(txRLPIdxArr, txRLPArr)

	return trieObj, nil
}

func generateTrie(paths [][]byte, values [][]byte) *trie.Trie {
	if len(paths) != len(values) {
		log.Fatal("Paths array and Values array have different lengths when generating Trie")
	}

	trieDB := trie.NewDatabase(memorydb.New())
	trieObj, _ := trie.New(common.Hash{}, trieDB) // empty trie

	for idx := range paths {
		p := paths[idx]
		v := values[idx]

		trieObj.Update(p, v) // update trie with the rlp encode index and the rlp encoded transaction
	}

	_, err := trieObj.Commit(nil) // commit to database (which in this case is stored in memory)
	if err != nil {
		log.Fatalf("commit error: %v", err)
	}

	return trieObj
}
