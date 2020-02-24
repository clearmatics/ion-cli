package core

import (
	contract "github.com/clearmatics/ion-cli/contracts"
	"github.com/ethereum/go-ethereum/common/compiler"
	"gotest.tools/assert"
	"io/ioutil"
	"os"
	"testing"
)

const testContractCode = "pragma solidity ^0.4.12;contract Contract{}"
const badTestContractCode = "pragma solidity ^0.4.12contract Contract{}"

func Test_GetDefaultSolidityCompiler(t *testing.T) {
	file, err := GetDefaultSolidityCompiler()
	if err != nil {
		t.Error(err)
	}
	defer DestroyTempFile(file.Name())

	solidity, err := compiler.SolidityVersion(file.Name())
	assert.Equal(t, err, nil)
	assert.Assert(t, solidity != nil)

	assert.Equal(t, solidity.Version, defaultSolidityVersion)
}

func Test_GetVersionedSolidityCompilerFromContract(t *testing.T) {
	file, err := HelperWritetemptestcontract(testContractCode)
	if err != nil {
		t.Error(err)
	}
	defer DestroyTempFile(file.Name())
	assert.Assert(t, err == nil)

	version, err := GetSolidityContractVersion(file.Name())
	assert.Assert(t, err == nil)
	assert.Equal(t, version, "0.4.12")

	solc, err := getSolidityCompilerLinux(version)
	if err != nil {
		t.Error(err)
	}
	defer DestroyTempFile(solc.Name())

	solidity, err := compiler.SolidityVersion(solc.Name())
	assert.Equal(t, err, nil)
	assert.Assert(t, solidity != nil)
	assert.Equal(t, solidity.Version, version)

	compiledContract, err := contract.CompileContractAt(file.Name(), solc.Name())
	if err != nil {
		t.Error(err)
	}
	assert.Assert(t, compiledContract != nil)
	assert.Equal(t, compiledContract.Info.LanguageVersion, "0.4.12")
	assert.Equal(t, compiledContract.Info.CompilerVersion, "0.4.12")

	file, err = HelperWritetemptestcontract(badTestContractCode)
	if err != nil {
		t.Error(err)
	}
	defer DestroyTempFile(file.Name())

	version, err = GetSolidityContractVersion(file.Name())
	assert.Assert(t, version == "")
	assert.Assert(t, err != nil)
}

func HelperWritetemptestcontract(code string) (*os.File, error) {
	file, err := CreateTempFile("contracttest.sol")
	if err != nil {
		return nil, err
	}

	err = ioutil.WriteFile(file.Name(), []byte(code), 0700)
	if err != nil {
		return nil, err
	}

	err = file.Close()
	if err != nil {
		return nil, err
	}

	return file, nil
}
