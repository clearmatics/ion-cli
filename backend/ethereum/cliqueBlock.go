package ethereum

import (
	"encoding/hex"
	"fmt"
	"github.com/clearmatics/ion-cli/utils"
	"github.com/ethereum/go-ethereum/rlp"
	"github.com/ethereum/go-ethereum/core/types"
)

// Implements the standard Block Interface of the session object
type CliqueBlockHeader struct {
	Header     *types.Header 	 `json:"header"`
	RlpSignedEncoded []byte        `json:"rlp_signed"`
	RlpUnsignedEncoded []byte        `json:"rlp_unsigned"`
}

// calculate and assign the rlp form of the header
func (b *CliqueBlockHeader) RlpEncode() (err error) {
	// Encode the orginal block header
	_, err = rlp.EncodeToBytes(&b)
	if err != nil {
		fmt.Println("can't RLP encode requested block:", err)
		return
	}

	// Generate an interface to encode the blockheader without the signature in the extraData
	b.RlpSignedEncoded, err = utils.RlpEncodeBlock(b.Header)
	if err != nil {
		return
	}

	b.RlpUnsignedEncoded, err = utils.RlpEncodeUnsignedBlock(b.Header)
	if err != nil {
		return
	}
	return nil
}

func (b *CliqueBlockHeader) GetByNumber(rpcURL string, number string) (err error) {
	fmt.Println("Connecting to the RPC client..")

	eth, err := GetClient(rpcURL)
	//returnIfError(err)

	b.Header, _, err = eth.GetBlockByNumber(number)
	if err != nil {
		return err
	}

	return nil
}

func (b *CliqueBlockHeader) GetByHash(rpcURL string, hash string) (err error) {
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
func (b CliqueBlockHeader) String() string{
	return fmt.Sprintf(hex.EncodeToString(b.Header.Extra))
}
