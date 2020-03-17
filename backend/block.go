package backend

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/ethereum/go-ethereum/rlp"
	"github.com/ethereum/go-ethereum/core/types"
)
// Every block type must implement the standard Block Interface of the session object

type EthBlockHeader struct {
	Header     *types.Header `json:"header"`
	RlpEncoded string        `json:"rlp_encoded"`
}

// calculate and assign the rlp form of the header
func (b EthBlockHeader) RlpEncode() (err error) {
	// Encode the orginal block header
	rlpHead, err := rlp.EncodeToBytes(&b.Header)
	if err != nil {
		fmt.Println("can't RLP encode requested block:", err)
		return
	}

	b.RlpEncoded = hex.EncodeToString(rlpHead)
	return
}

func (b EthBlockHeader) Marshal() ([]byte, error) {
	by, err := json.MarshalIndent(b, "", "	")
	if err != nil {
		fmt.Errorf("error marshaling the Eth Block Header")
		return []byte{}, err
	}

	return by, nil
}

func (b EthBlockHeader) Print() {
	//json.Unmarshal(b, &b)

	fmt.Println(b.Header, b.RlpEncoded)
}
