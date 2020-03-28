package backend

// a chain holds network and wallet configs of a profile, plus cached data for subsequent calls
type Chain struct {
	Network NetworkInfo `json:"network"`
	Accounts Accounts `json:"accounts"`

	// cache
	Block Block `json:"block"`
	Transaction Transaction `json:"transaction"`

	// ion proofs

	// contract addresses
}

type Chains map[string]*Chain

// tells if a chain with id exists
func (c Chains) Exist(id string) bool {
	return c[id].Network != NetworkInfo{}
}

// add a chain object with id id
func (c Chains) Add (id string, network NetworkInfo) {
	c[id] = &Chain{
		Network:     network,
		Accounts:    Accounts{},
		Transaction: Transaction{},
		Block:Block{},
	}
}

// remove a chain object with id id
func (c Chains) Remove (id string) {
	delete(c, id)
}
