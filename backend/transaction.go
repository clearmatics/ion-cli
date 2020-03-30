package backend

import "encoding/json"

type Transaction struct {
	Tx    json.RawMessage `json:"tx"`
	Interface TransactionInterface `json:"-"`
}

type TransactionInterface interface {
	// assign the rlp encoded ion proof for a specific tx hash ready for submission
	GenerateIonProof(rpcURL string, hash string) error
	Marshal() ([]byte, error)
}