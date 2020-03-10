// Copyright (c) 2018 Clearmatics Technologies Ltd

package utils

import (
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/rlp"
)

// EncodePrefix calculate prefix of the entire signed block
func RlpEncodeUnsignedBlock(block *types.Header) ([]byte, error) {
	block.Extra = block.Extra[:len(block.Extra)-65]

	encodedBlock, err := rlp.EncodeToBytes(&block)
	if err != nil {
		return []byte{}, err
	}

	return encodedBlock, nil

}

// EncodePrefix calculate prefix of the entire signed block
func RlpEncodeBlock(block *types.Header) ([]byte, error) {
	encodedBlock, err := rlp.EncodeToBytes(&block)
	if err != nil {
		return []byte{}, err
	}

	return encodedBlock, nil
}

// EncodePrefix calculate prefix of the entire signed block
func EncodePrefix(blockHeader types.Header) ([]byte, error) {
	blockHeader.Extra = blockHeader.Extra[:len(blockHeader.Extra)-65]
	encodedPrefixBlock, err := rlp.EncodeToBytes(blockHeader)
	if err != nil {
		return []byte{}, err
	}

	return encodedPrefixBlock[1:3], nil
}

// EncodeExtraData calculate prefix of the extraData with the signature
func EncodeExtraData(blockHeader types.Header) ([]byte, error) {
	extraData := blockHeader.Extra[:len(blockHeader.Extra)-65]
	encodedExtraData, err := rlp.EncodeToBytes(extraData)
	if err != nil {
		return []byte{}, err
	}

	return encodedExtraData[0:1], nil
}
