package backend

import (
	"encoding/hex"
	"fmt"
	"github.com/ethereum/go-ethereum/rlp"
)

// Implements the standard Block Interface of the session object

// calculate and assign the rlp form of the header
func (b *EthBlockHeader) RlpEncode() (err error) {
	rlpHead, err := rlp.EncodeToBytes(&b.Header)
	if err != nil {
		fmt.Println("can't RLP encode requested block:", err)
		return
	}

	b.RlpEncoded = hex.EncodeToString(rlpHead)
	return
}


// TODO if needed
func (b EthBlockHeader) String() string{
	return fmt.Sprintf(hex.EncodeToString(b.Header.Extra))
}
