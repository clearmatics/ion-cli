package backend

import (
	"encoding/json"
	//"github.com/clearmatics/ion-cli/backend/ethereum"
)

type Block struct {
	Header json.RawMessage `json:"header"` // this is the polymorphic bit
	Interface BlockInterface `json:"-"`
}

type Blocks map[string]Block

type BlockInterface interface {
	RlpEncode() (err error)
	GetByNumber(rpcURL string, number string) (err error)
	GetByHash(rpcURL string, hash string) (err error)
	Marshal() (header []byte, err error)
	Print()
}
