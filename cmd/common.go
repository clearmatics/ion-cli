package cmd

import (
	"fmt"
	"os"
)

// utils for cli commands these functions will be available to all commands

func returnIfError(err error) {
	if err != nil {
		fmt.Println(err)
		fmt.Println("Quitting..")
		os.Exit(-1)
	}
}