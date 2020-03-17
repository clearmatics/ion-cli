package utils_test

import (
	contract "github.com/clearmatics/ion-cli/contracts"
	"github.com/clearmatics/ion-cli/utils"
	"github.com/ethereum/go-ethereum/common/compiler"
	"gotest.tools/assert"
	"io/ioutil"
	"os"
	"testing"
)

const testContractCode = "pragma solidity ^0.4.12;contract Contract{}"
const badTestContractCode = "pragma solidity ^0.4.12contract Contract{}"

func Test_GetDefaultSolidityCompiler(t *testing.T) {
	file, err := utils.GetDefaultSolidityCompiler()
	assert.NilError(t, err)
	defer utils.DestroyTempFile(file.Name())

	solidity, err := compiler.SolidityVersion(file.Name())
	assert.NilError(t, err)
	assert.Assert(t, solidity != nil)

	assert.Equal(t, solidity.Version, utils.DefaultSolidityVersion)
}

func Test_GetVersionedSolidityCompilerFromContract(t *testing.T) {
	file, err := HelperWritetemptestcontract(testContractCode)
	assert.NilError(t, err)
	defer utils.DestroyTempFile(file.Name())

	version, err := utils.GetSolidityContractVersion(file.Name())
	assert.NilError(t, err)
	assert.Equal(t, version, "0.4.12")

	solc, err := utils.GetSolidityCompilerLinux(version)
	assert.NilError(t, err)
	defer utils.DestroyTempFile(solc.Name())

	solidity, err := compiler.SolidityVersion(solc.Name())
	assert.NilError(t, err)
	assert.Assert(t, solidity != nil)
	assert.Equal(t, solidity.Version, version)

	compiledContract, err := contract.CompileContractAt(file.Name(), solc.Name())
	assert.NilError(t, err)
	assert.Assert(t, compiledContract != nil)
	assert.Equal(t, compiledContract.Info.LanguageVersion, "0.4.12")
	assert.Equal(t, compiledContract.Info.CompilerVersion, "0.4.12")

	file, err = HelperWritetemptestcontract(badTestContractCode)
	assert.NilError(t, err)
	defer utils.DestroyTempFile(file.Name())

	version, err = utils.GetSolidityContractVersion(file.Name())
	assert.Assert(t, version == "")
	assert.ErrorContains(t, err, "pragma solidity ^0.4.12contract Contract{}")
}

func HelperWritetemptestcontract(code string) (*os.File, error) {
	file, err := utils.CreateTempFile("contracttest.sol")
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
