package backend

import (
	"errors"
	"fmt"
	"github.com/clearmatics/ion-cli/backend/ethereum"
)

// a chain holds network and wallet configs of a profile, plus cached data for subsequent calls
type Chain struct {
	Network NetworkInfo `json:"network"`
	Accounts Accounts `json:"accounts"`
	Type string `json:"type"` // allow to identify block and transaction types of a chain

	// cache
	Block Block `json:"block"`
	Transaction Transaction `json:"transaction"`
}

type Chains map[string]Chain

// tells if a chain with id exists
func (c Chains) Exist(id string) bool {
	return c[id].Type != ""
}

// add a chain object with id id
// TODO enum checks on chainType
// TODO we may avoid adding chains with empty network info
func (c Chains) Add (id string, network NetworkInfo, chainType string) {
	c[id] = Chain{
		Network:     network,
		Accounts:    Accounts{},
		Type:	chainType,
		Transaction: Transaction{},
		Block:Block{},
	}
}

// remove a chain object with id id
func (c Chains) Remove (id string) {
	delete(c, id)
}

func (c *Chain) AssignImplementers() error {
	switch c.Type {
	case "eth":
		c.Transaction.Interface = &ethereum.EthTransaction{}
		c.Block.Interface = &ethereum.EthBlockHeader{}
	case "clique":
		c.Block.Interface = &ethereum.CliqueBlockHeader{}
		c.Transaction.Interface = &ethereum.EthTransaction{}
	case "ibft":
		c.Block.Interface = &ethereum.IBFTBlockHeader{}
		c.Transaction.Interface = &ethereum.EthTransaction{}
	default:
		return errors.New(fmt.Sprintf("The chain type %v is not recognised", c.Type))
	}

	return nil
}
