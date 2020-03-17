package backend

import (
	"encoding/hex"
	"fmt"
	"github.com/ethereum/go-ethereum/rlp"
)

// manages all the functions to encode/decode block headers of different types (consensus) as well

// calculate and assign the rlp form of the header
func (b *BlockHeader) RlpEncode() (err error) {
	// Encode the orginal block header
	rlpHead, err := rlp.EncodeToBytes(&b.Header)
	if err != nil {
		fmt.Println("can't RLP encode requested block:", err)
		return
	}

	b.RlpEncoded = hex.EncodeToString(rlpHead)
	return
}
