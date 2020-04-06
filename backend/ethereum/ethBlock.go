package ethereum

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/ethereum/go-ethereum/rlp"
	"github.com/ethereum/go-ethereum/core/types"
)

// Implements the standard Block Interface of the profile object
type EthBlockHeader struct {
	Header     *types.Header 	 `json:"header"`
	RlpEncoded string        `json:"rlp_encoded"`
}

// marshal itself so that the
func (b *EthBlockHeader) Marshal() (header []byte, err error) {
	return json.Marshal(b)
}

// calculate and assign the rlp form of the header
func (b *EthBlockHeader) RlpEncode() (err error) {
	rlpH, err := rlp.EncodeToBytes(&b.Header)
	if err != nil {
		fmt.Println("can't RLP encode requested block:", err)
		return err
	}

	b.RlpEncoded = hex.EncodeToString(rlpH)

	return nil
}

func (b *EthBlockHeader) GetByNumber(rpcURL string, number string) (err error) {
	fmt.Println("Connecting to the RPC client..")

	eth, _ := GetClient(rpcURL)

	b.Header, _, err = eth.GetBlockByNumber(number)
	if err != nil {
		return err
	}

	return nil
}

func (b *EthBlockHeader) GetByHash(rpcURL string, hash string) (err error) {
	fmt.Println("Connecting to the RPC client..", rpcURL)

	eth, _ := GetClient(rpcURL)

	b.Header, _, err = eth.GetBlockByHash(hash)

	if err != nil {
		return err
	}

	return nil
}
