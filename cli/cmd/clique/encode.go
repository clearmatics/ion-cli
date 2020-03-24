package clique

import (
	"fmt"
	"github.com/clearmatics/ion-cli/utils"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/rlp"
)

func rlpEncodeClique(blockHeader *types.Header) (rlpSignedBlock []byte, rlpUnsignedBlock []byte, err error) {

	_, err = rlp.EncodeToBytes(&blockHeader)
	if err != nil {
		fmt.Println("can't RLP encode requested block:", err)
		return
	}

	// Generate an interface to encode the blockheader without the signature in the extraData
	rlpSignedBlock, err = utils.RlpEncodeBlock(blockHeader)
	if err != nil {
		return
	}

	rlpUnsignedBlock, err = utils.RlpEncodeUnsignedBlock(blockHeader)
	if err != nil {
		return
	}

	return rlpSignedBlock, rlpUnsignedBlock, nil
}
