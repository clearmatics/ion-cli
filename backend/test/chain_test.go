package backend_test

import (
	"github.com/clearmatics/ion-cli/backend"
	"github.com/clearmatics/ion-cli/backend/ethereum"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_AddChain(t *testing.T) {
	chains := backend.Chains{}
	id := "test"

	chains.Add(id, backend.NetworkInfo{}, "typeTest")
	assert.True(t, chains.Exist(id))
}

func Test_RemoveChain(t *testing.T) {
	chains := backend.Chains{}
	id := "test"

	chains.Add(id, backend.NetworkInfo{}, "typeTest")
	assert.True(t, chains.Exist(id))

	chains.Remove(id)
	assert.False(t, chains.Exist(id))
}

// TODO we may have all the possible block ypes encapsulated and this test will be automated
func Test_AssignImplementers(t *testing.T) {
	// ETHEREUM
	ethChain := backend.Chain{
		Network:     backend.NetworkInfo{},
		Accounts:    nil,
		Type:        "eth",
		Block:       backend.Block{},
		Transaction: backend.Transaction{},
	}
	err := ethChain.AssignImplementers()
	assert.Nil(t, err)
	assert.Equal(t, ethChain.Block.Interface, &ethereum.EthBlockHeader{})
	assert.Equal(t, ethChain.Transaction.Interface, &ethereum.EthTransaction{})

	// CLIQUE
	cliqueChain := backend.Chain{
		Network:     backend.NetworkInfo{},
		Accounts:    nil,
		Type:        "clique",
		Block:       backend.Block{},
		Transaction: backend.Transaction{},
	}

	err = cliqueChain.AssignImplementers()
	assert.Nil(t, err)
	assert.Equal(t, cliqueChain.Block.Interface, &ethereum.CliqueBlockHeader{})
	assert.Equal(t, cliqueChain.Transaction.Interface, &ethereum.EthTransaction{})

	// IBFT
	IBFTChain := backend.Chain{
		Network:     backend.NetworkInfo{},
		Accounts:    nil,
		Type:        "ibft",
		Block:       backend.Block{},
		Transaction: backend.Transaction{},
	}

	err = IBFTChain.AssignImplementers()
	assert.Nil(t, err)
	assert.Equal(t, IBFTChain.Block.Interface, &ethereum.IBFTBlockHeader{})
	assert.Equal(t, IBFTChain.Transaction.Interface, &ethereum.EthTransaction{})

	// NOT RECOGNIZED
	ErrorChain := backend.Chain{
		Network:     backend.NetworkInfo{},
		Accounts:    nil,
		Type:        "not a chain",
		Block:       backend.Block{},
		Transaction: backend.Transaction{},
	}

	err = ErrorChain.AssignImplementers()
	assert.Error(t, err)
}