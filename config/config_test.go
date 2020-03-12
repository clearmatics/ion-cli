// Copyright (c) 2018 Clearmatics Technologies Ltd

package config_test

import (
	"fmt"
	"runtime"
	"strings"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/stretchr/testify/assert"

	"github.com/clearmatics/ion-cli/config"
)

func Test_ReadValidKeystore(t *testing.T) {
	path := findPath() + "../keystore/UTC--2018-06-05T09-31-57.109288703Z--2be5ab0e43b6dc2908d5321cf318f35b80d0c10d"
	contents := config.ReadString(path)

	const val = "{\"address\":\"2be5ab0e43b6dc2908d5321cf318f35b80d0c10d\",\"crypto\":{\"cipher\":\"aes-128-ctr\",\"ciphertext\":\"0b11aa865046778a1b16a9b8cb593df704e3fe09f153823d75442ad1aab66caa\",\"cipherparams\":{\"iv\":\"4aa66b789ee2d98cf77272a72eeeaa50\"},\"kdf\":\"scrypt\",\"kdfparams\":{\"dklen\":32,\"n\":262144,\"p\":1,\"r\":8,\"salt\":\"b957fa7b7577240fd3791168bbe08903af4c8cc62c304f1df072dc2a59b1765e\"},\"mac\":\"197a06eb0449301d871400a6bdf6c136b6f7658ee41e3f2f7fd81ca11cd954a3\"},\"id\":\"a3cc1eae-3e36-4659-b759-6cf416216e72\",\"version\":3}"

	assert.Equal(t, val, contents)

}

func Test_InitUser(t *testing.T) {
	keystore := "../keystore/UTC--2018-06-05T09-31-57.109288703Z--2be5ab0e43b6dc2908d5321cf318f35b80d0c10d"
	password := "password1"
	expectedFrom := common.HexToAddress("2be5ab0e43b6dc2908d5321cf318f35b80d0c10d")
	expectedPrivateKey := "e176c157b5ae6413726c23094bb82198eb283030409624965231606ec0fbe65b"

	user, err := config.InitUser(keystore, password)
	assert.Equal(t, err, nil)

	assert.Equal(t, user.Auth.From, expectedFrom)
	privateKey := fmt.Sprintf("%x", crypto.FromECDSA(user.Key.PrivateKey))
	assert.Equal(t, privateKey, expectedPrivateKey)

}

func Test_ReadConfig(t *testing.T) {
	configuration := "configtest.json"

	expectedAccounts := []config.ConfigAccount{{"me", "keystore/UTC--2018-11-14T13-34-31.599642840Z--b8844cf76df596e746f360957aa3af954ef51605", "test"}}
	expectedContracts := []config.ConfigContracts{{"ion", "contracts/Ion.sol"}}

	setup, err := config.ReadSetup(configuration)
	assert.Equal(t, err, nil)

	assert.Equal(t, setup.Accounts, expectedAccounts)
	assert.Equal(t, setup.Contracts, expectedContracts)
}

func findPath() string {
	_, path, _, _ := runtime.Caller(0)
	pathSlice := strings.Split(path, "/")
	return strings.Trim(path, pathSlice[len(pathSlice)-1])
}
