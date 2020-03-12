package backend

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

func PersistObject(obj interface{}, file string) error {
	b, err := json.Marshal(obj)
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

// Takes path to a JSON and returns a struct of the contents
//func ReadSetup(config string) (Setup, error) {
//	setup := Setup{}
//	raw, err := ioutil.ReadFile(config)
//	if err != nil {
//		return setup, err
//	}
//
//	err = json.Unmarshal(raw, &setup)
//	if err != nil {
//		return setup, err
//	}
//
//	return setup, nil
//}

// Takes path to a JSON and returns a string of the contents
func ReadString(path string) (contents string) {
	raw, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Print(err, "\n")
	}

	contents = string(raw)

	return

}
