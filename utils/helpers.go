// Copyright (c) 2018 Clearmatics Technologies Ltd

package utils

import (
	"context"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/ethclient"
)

func PersistObject(obj interface{}, file string) error {
	b, err := json.MarshalIndent(obj, "", " ")
	if err != nil {
		fmt.Errorf("error marshaling the object")
		return err
	}

	err = ioutil.WriteFile(file, b, 0644)
	if err != nil {
		fmt.Errorf("error writing to file")
		return err
	}

	return nil
}

// Takes path to a JSON and returns a string of the contents
func ReadString(path string) (contents string) {
	raw, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Print(err, "\n")
	}

	contents = string(raw)

	return

}

func GetNonce(client *ethclient.Client, auth *bind.TransactOpts) {
	// Find the correct tx nonce
	nonce, err := client.PendingNonceAt(context.Background(), auth.From)
	if err != nil {
		log.Fatalf("Failed to calculate nonce: %v", err)
	}

	auth.Nonce = big.NewInt(int64(nonce))
}

// Needed to convert transaction strings into the correct format
func StringToBytes32(input string) (output [32]byte, err error) {
	// Check string length is correct 64
	if len(input) == 64 {
		inputBytes, err := hex.DecodeString(input)
		if err != nil {
			log.Fatalf("Failed to encode string as bytes: %v", err)
		}

		copy(output[:], inputBytes[:len(output)])

		return output, nil
	} else if len(input) == 66 && input[:2] == "0x" {
		inputBytes, err := hex.DecodeString(input[2:])
		if err != nil {
			log.Fatalf("Failed to encode string as bytes: %v", err)
		}

		copy(output[:], inputBytes[:len(output)])

		return output, nil
	} else {
		return [32]byte{}, fmt.Errorf("Failed to encode string as bytes32, incorrect string input")
	}

}
