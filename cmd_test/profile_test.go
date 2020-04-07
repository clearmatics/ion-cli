package cmd_test

import (
	"encoding/json"
	"github.com/clearmatics/ion-cli/backend"
	"github.com/clearmatics/ion-cli/cmd"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"os"
	"testing"
)

// TODO don-t hardcode
const profilePath = "/home/andreadinenno/go/src/github.com/clearmatics/ion-cli/cmd_test/profiles.json"
const configPath = "/home/andreadinenno/go/src/github.com/clearmatics/ion-cli/cmd_test/config.json"
const rootPath = "/home/andreadinenno/go/src/github.com/clearmatics/ion-cli/"

const c = "chain"
const n = "network"
const hTy = "header type"
const u = "url"
const p = "andrea"

var testNet = backend.NetworkInfo{
	Name:   n,
	Url:    "url",
	Header: hTy,
}

var testChain = &backend.Chain{
	Network:     testNet,
	Accounts:    nil,
	Type:        hTy,
	Block:       backend.Block{},
	Transaction: backend.Transaction{},
}

func Test_Add_Profile_Success(t *testing.T) {
	// clean test file
	err := cmd.CleanJSONFile(profilePath)
	if err != nil {
		assert.Fail(t, err.Error())
	}

	// set working dir
	err = os.Chdir(rootPath)
	if err != nil {
		assert.Fail(t, err.Error())
	}

	//command tree - always start from root cmd
	rootCmd := cmd.NewRootCmd()
	rootCmd.AddCommand(cmd.NewProfileCmd())

	// input
	id := "andrea"
	rootCmd.SetArgs([]string{"profile", "add", id, "--profiles", profilePath})

	// trigger
	err = rootCmd.Execute()
	if err != nil {
		assert.Fail(t, err.Error())
	}

	// test
	profiles, err := loadProfiles(profilePath)
	if err != nil {
		assert.Fail(t, err.Error())
	}

	assert.True(t, profiles.Exist(id))
}

func Test_Delete_Profile(t *testing.T) {
	// clean test file
	err := cmd.CleanJSONFile(profilePath)
	if err != nil {
		assert.Fail(t, err.Error())
	}

	// set working dir
	err = os.Chdir(rootPath)
	if err != nil {
		assert.Fail(t, err.Error())
	}

	//command tree - always start from root cmd
	rootCmd := cmd.NewRootCmd()
	rootCmd.AddCommand(cmd.NewProfileCmd())

	// add
	id := "andrea"
	rootCmd.SetArgs([]string{"profile", "add", id, "--profiles", profilePath})

	// trigger
	err = rootCmd.Execute()
	if err != nil {
		assert.Fail(t, err.Error())
	}

	// delete
	rootCmd.SetArgs([]string{"profile", "del", id, "--profiles", profilePath})
	err = rootCmd.Execute()
	if err != nil {
		assert.Fail(t, err.Error())
	}

	// test
	profiles, err := loadProfiles(profilePath)
	if err != nil {
		assert.Fail(t, err.Error())
	}

	assert.False(t, profiles.Exist(id))
}

func Test_Add_Chain(t *testing.T) {

	assert.NoError(t, cmd.CleanJSONFile(profilePath))
	assert.NoError(t, cmd.CleanJSONFile(configPath))

	// set working dir
	assert.NoError(t, os.Chdir(rootPath))

	//command tree - always start from root cmd
	rootCmd := cmd.NewRootCmd()
	rootCmd.AddCommand(cmd.NewProfileCmd())

	// add network in configs
	rootCmd.SetArgs([]string{"config", "networks", "add", n, u, hTy, "--config", configPath})
	assert.NoError(t, rootCmd.Execute())

	// add profile
	rootCmd.SetArgs([]string{"profile", "add", p, "--profiles", profilePath})
	assert.NoError(t, rootCmd.Execute())

	// add chain with bad networkID
	rootCmd.SetArgs([]string{"profile", "chains", "add", p, c, "net", hTy, "--profiles", profilePath, "--config", configPath})
	assert.NoError(t, rootCmd.Execute())

	profiles, err := loadProfiles(profilePath)
	assert.NoError(t, err)
	assert.False(t, profiles[p].Chains.Exist(c))

	// add proper chain
	rootCmd.SetArgs([]string{"profile", "chains", "add", p, c, n, hTy, "--profiles", profilePath, "--config", configPath})
	assert.NoError(t, rootCmd.Execute())

	profiles, err = loadProfiles(profilePath)
	assert.NoError(t, err)

	assert.True(t, profiles.Exist(p))
	assert.True(t, profiles[p].Chains.Exist(c))
	assert.Equal(t, profiles[p].Chains[c].Type, hTy)
	assert.Equal(t, profiles[p].Chains[c].Network, testNet)

}


func loadProfiles(profilesPath string) (backend.Profiles, error) {
	jsonFile, err := os.Open(profilesPath)
	if err != nil {
		return backend.Profiles{}, err
	}
	defer jsonFile.Close()

	profiles := backend.Profiles{}

	b, err := ioutil.ReadAll(jsonFile)

	err = json.Unmarshal(b, &profiles)
	return profiles, err
}

