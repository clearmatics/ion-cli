package ethereum

import (
	"encoding/hex"
	"fmt"
	"github.com/ethereum/go-ethereum/rlp"
	"github.com/ethereum/go-ethereum/core/types"
)

// Implements the standard Block Interface of the session object
type EthBlockHeader struct {
	Header     *types.Header 	 `json:"header"`
	RlpEncoded string        `json:"rlp_encoded"`
}

// calculate and assign the rlp form of the header
func (b *EthBlockHeader) RlpEncode() (err error) {
	rlpHead, err := rlp.EncodeToBytes(&b.Header)
	if err != nil {
		fmt.Println("can't RLP encode requested block:", err)
		return err
	}

	b.RlpEncoded = hex.EncodeToString(rlpHead)
	return nil
}

func (b *EthBlockHeader) GetByNumber(rpcURL string, number string) (err error) {
	fmt.Println("Connecting to the RPC client..")

	eth, err := GetClient(rpcURL)
	//returnIfError(err)

	b.Header, _, err = eth.GetBlockByNumber(number)
	if err != nil {
		return err
	}

	return nil
}

func (b *EthBlockHeader) GetByHash(rpcURL string, hash string) (err error) {
	fmt.Println("Connecting to the RPC client..")

	eth, err := GetClient(rpcURL)
	//returnIfError(err)

	b.Header, _, err = eth.GetBlockByHash(hash)
	if err != nil {
		return err
	}

	return nil
}



// TODO if needed
func (b EthBlockHeader) String() string{
	return fmt.Sprintf(hex.EncodeToString(b.Header.Extra))
}
