package backend

import "encoding/json"

type Transaction struct {
	Tx    json.RawMessage `json:"tx"`
	Interface TransactionInterface `json:"-"`
}

type TransactionInterface interface {
	GenerateIonProof(rpcURL string, hash string) error
	Marshal() ([]byte, error)
	Print()
}