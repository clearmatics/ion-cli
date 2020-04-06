package ethereum

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/clearmatics/ion-cli/utils"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

type EthTransaction struct {
	Tx    *types.Transaction `json:"tx"`
	Proof string             `json:"proof"`
}

func (tx *EthTransaction) GenerateIonProof(rpcURL string, hash string) error {
	eth, err := GetClient(rpcURL)

	tx.Tx, _, err = eth.GetTransactionByHash(hash)
	if err != nil {
		return err
	}

	txHash := common.HexToHash(hash)
	data, err := utils.FetchProofData(eth.Client, txHash)
	if err != nil {
		return err
	}

	proof, err := utils.GenerateIonProof(*data)
	if err != nil {
		return err
	}

	tx.Proof = hex.EncodeToString(proof)

	return nil
}

func (tx *EthTransaction) Marshal() ([]byte, error) {
	return json.Marshal(tx)
}

func (tx *EthTransaction) Print() {
	fmt.Println("Tx hash:", tx.Tx.Hash().String(), "\nION proof:", tx.Proof)

}

