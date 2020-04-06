package backend_test

import (
	"encoding/json"
	"github.com/clearmatics/ion-cli/backend"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"os"
	"testing"
)

func Test_AddProfile(t *testing.T) {
	profiles := backend.Profiles{}
	id := "andrea"

	profiles.Add(id)
	assert.True(t, profiles.Exist(id))
}

func Test_RemoveProfile(t *testing.T) {
	profiles := backend.Profiles{}
	id := "andrea"

	profiles.Add(id)
	assert.True(t, profiles.Exist(id))

	profiles.Remove(id)
	assert.False(t, profiles.Exist(id))
}

func Test_ActiveProfile(t *testing.T) {
	profiles := backend.Profiles{}
	id := "andrea"
	chainId := "test"

	profiles.Add(id)
	// not active yet
	assert.False(t, profiles[id].IsActive())

	profiles[id].Chains.Add(chainId, backend.NetworkInfo{}, "eth")
	assert.True(t, profiles[id].IsActive())
}

func Test_SaveProfilesFile_Success(t *testing.T) {
	profiles := backend.Profiles{}
	filePath := "./profiles.json"
	id := "andrea"

	defer os.Remove(filePath)

	profiles.Add(id)
	err := profiles.Save(filePath)
	if err != nil {
		assert.Fail(t, err.Error())
	}

	jsonFile, err := os.Open(filePath)
	if err != nil {
		assert.Fail(t, err.Error())
	}
	defer jsonFile.Close()

	fileProfiles := backend.Profiles{}

	b, _ := ioutil.ReadAll(jsonFile)
	err = json.Unmarshal(b, &fileProfiles)

	assert.True(t, fileProfiles.Exist(id))
}

func Test_SaveProfilesFile_WrongPath(t *testing.T) {
	profiles := backend.Profiles{}
	filePath := "./wrong/profiles.json"
	id := "andrea"

	defer os.Remove(filePath)

	profiles.Add(id)
	err := profiles.Save(filePath)
	assert.Error(t, err)

}