package backend

// a chain holds network and wallet configs of a profile, plus cached data for subsequent calls
type Chain struct {
	Network Network `json:"network"`
	Accounts Accounts `json:"accounts"`

	// cache
	Blocks BlockMap `json:"blocks"`
	Transaction Transaction `json:"transaction"`

	// ion proofs

	// contract addresses
}

type Chains map[string]Chain

type Network struct {
	Name string `json:"name"`
	Url string `json:"url"`
}

func (c Chains) Exist(id string) bool {
	return c[id].Network != Network{}
}

func (c Chains) Add (id string, network Network) {
	c[id] = Chain{
		Network:     network,
		Accounts:    make(map[string]AccountInfo),
		Blocks:      nil,
		Transaction: Transaction{},
	}
}

func (c Chains) Remove (id string) {
	delete(c, id)
}
