// Copyright (c) 2018-2020 Clearmatics Technologies Ltd

package utils_test

import (
	"encoding/json"
	"github.com/clearmatics/ion-cli/utils"
	"gotest.tools/assert"
	"testing"
)

// Testing trie generation against Rinkeby block 2657422

func Test_TxTrie(t *testing.T) {
	var proofData utils.IonProofData
	err := json.Unmarshal([]byte(ProofDataStr), &proofData)
	assert.NilError(t, err)

	expectedTxTrieRoot := "0x07f36c7ad26564fa65daebda75a23dfa95d660199092510743f6c8527dd72586"

	txtrie, err := utils.TxTrie(proofData.BlockTransactions)

	assert.NilError(t, err)
	assert.Equal(t, txtrie.Hash().String(), expectedTxTrieRoot)
}

func Test_ReceiptTrie(t *testing.T) {
	var proofData utils.IonProofData
	err := json.Unmarshal([]byte(ProofDataStr), &proofData)
	assert.NilError(t, err)

	expectedReceiptTrieRoot := "0x907121bec78b40e8256fac47867d955c560b321e93fc9f046f919ffb5e3823ff"

	receiptTrie, err := utils.ReceiptTrie(proofData.BlockReceipts)

	assert.NilError(t, err)
	assert.Equal(t, receiptTrie.Hash().String(), expectedReceiptTrieRoot)
}
