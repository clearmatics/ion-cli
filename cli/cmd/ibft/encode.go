package ibft

import (
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/rlp"
)

// RlpEncodeIBFT returns rlp encoded block header from an IBFT consensus chain
func rlpEncodeIBFT(blockHeader *types.Header) (proposalBlock []byte, commitBlock []byte, err error) {
	// Generate an interface to encode the blockheader without the signature in the extraData

	commitBlock, err = encodeCommitBlock(blockHeader)
	if err != nil {
		return
	}

	proposalBlock, err = encodeProposalBlock(blockHeader)

	return
}

// encodeProposalBlock returns the block signed by the block proposer of an IBFT chain
func encodeProposalBlock(block *types.Header) ([]byte, error) {
	// extract istanbul extraData from the block header
	istanbul := block.Extra[IstanbulExtraVanity:]

	var istanbulExtra *IstanbulExtra
	err := rlp.DecodeBytes(istanbul, &istanbulExtra)
	if err != nil {
		return []byte{}, err
	}

	// remove proposal seal and commit seals
	istanbulExtra.Seal = make([]byte, 0)
	istanbulExtra.CommittedSeal = make([][]byte, 0)

	// Encode istanbulExtra
	encodedIstanbulExtra, err := rlp.EncodeToBytes(&istanbulExtra)
	if err != nil {
		return []byte{}, err
	}
	block.Extra = append(block.Extra[:IstanbulExtraVanity], encodedIstanbulExtra[:]...)

	return rlp.EncodeToBytes(&block)
}

// encodeCommitBlock returns the block signed by the validators of an IBFT chain
func encodeCommitBlock(block *types.Header) ([]byte, error) {
	// extract istanbul extraData from the block header
	istanbul := block.Extra[IstanbulExtraVanity:]

	var istanbulExtra *IstanbulExtra
	err := rlp.DecodeBytes(istanbul, &istanbulExtra)
	if err != nil {
		return []byte{}, err
	}

	// remove commit seals
	istanbulExtra.CommittedSeal = make([][]byte, 0)

	// Encode istanbulExtra
	encodedIstanbulExtra, err := rlp.EncodeToBytes(&istanbulExtra)
	if err != nil {
		return []byte{}, err
	}
	block.Extra = append(block.Extra[:32], encodedIstanbulExtra[:]...)
	// fmt.Printf("%x\n", block.Extra)

	return rlp.EncodeToBytes(&block)
}
