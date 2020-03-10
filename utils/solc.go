package utils

import (
	"errors"
	"fmt"
	contract "github.com/clearmatics/ion-cli/contracts"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"runtime"
	"strings"
)

const DefaultSolidityVersion = "0.5.12" // Match Ion Release contract versions

// Hacky way to find correct contract version
// Attempts to compile using default solc version and parses solc error to retrieve correct version
func GetSolidityContractVersion(contractFilePath string) (string, error) {
	solc, err := GetDefaultSolidityCompiler()
	if err != nil {
		return "", err
	}
	defer DestroyTempFile(solc.Name())

	compiledContract, err := contract.CompileContractAt(contractFilePath, solc.Name())
	if compiledContract != nil {
		return compiledContract.Info.LanguageVersion, nil
	}

	if err != nil {
		version, err := extractSolidityVersionFromCompilationError(err)
		if err != nil {
			return "", err
		}
		return version, nil
	}

	return "", errors.New("could not retrieve contract version")
}

func extractSolidityVersionFromCompilationError(err error) (string, error) {
	errMsg := err.Error()
	compilationErrorString := "Source file requires different compiler version"

	if strings.Contains(errMsg, compilationErrorString) {
		index := strings.Index(errMsg, "pragma solidity ")
		if index != -1 {
			version := strings.ReplaceAll(strings.Split(strings.Split(errMsg, "pragma solidity ")[1], ";")[0], "^", "")
			return version, nil
		} else {
			return "", errors.New("compiler version error, but can't find solidity version")
		}
	}

	return "", err
}

func GetDefaultSolidityCompiler() (*os.File, error) {
	switch runtime.GOOS {
	case "darwin":
		return nil, nil
	case "linux":
		return GetSolidityCompilerLinux(DefaultSolidityVersion)
	}

	return nil, nil
}

func GetSolidityCompilerVersion(version string) (*os.File, error) {
	switch runtime.GOOS {
	case "darwin":
		return nil, nil
	case "linux":
		fmt.Printf("System is Linux\n")
		return GetSolidityCompilerLinux(version)
	}

	return nil, nil
}

func GetSolidityCompilerLinux(version string) (*os.File, error) {
	fmt.Printf("Getting Solidity compiler for version %s\n", version)
	url := fmt.Sprintf("https://github.com/ethereum/solidity/releases/download/v%s/solc-static-linux", version)

	file, err := downloadSolidityCompiler(url)
	if err != nil {
		return nil, err
	}

	err = os.Chmod(file.Name(), 0700)
	if err != nil {
		return nil, err
	}

	return file, nil
}

func downloadSolidityCompiler(url string) (*os.File, error) {
	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	file, err := CreateTempFile("ioncli-solc")

	// Write the body to file
	_, err = io.Copy(file, resp.Body)
	if err != nil {
		return nil, err
	}

	err = file.Close()
	if err != nil {
		return nil, err
	}

	return file, nil
}

func CreateTempFile(name string) (*os.File, error) {
	dir, err := ioutil.TempDir("", "")
	file, err := ioutil.TempFile(dir, name)
	if err != nil {
		return nil, err
	}

	return file, nil
}

func DestroyTempFile(filepath string) error {
	return os.Remove(filepath)
}
