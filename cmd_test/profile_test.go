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
const profilePath = "/home/andreadinenno/go/src/github.com/clearmatics/ion-cli/config/test.json"
const rootPath = "/home/andreadinenno/go/src/github.com/clearmatics/ion-cli/"

func Test_Add_Profile_Success(t *testing.T) {
	// clean test file
	err := cleanFile(profilePath)
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

func cleanFile(path string) (error) {
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	_, err = file.Write([]byte("{}"))
	if err != nil {
		return err
	}

	return nil
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

