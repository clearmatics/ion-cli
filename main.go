// Copyright (c) 2018 Clearmatics Technologies Ltd

package main

import (
	"flag"
	"fmt"
	"github.com/clearmatics/ion-cli/cli"
	"github.com/clearmatics/ion-cli/config"
)

func main() {
	configFilePtr := flag.String("config", "", "File location of configuration file")
	flag.Parse()

	if *configFilePtr != "" {
		fmt.Printf("Configuring session...\n")
		configuration, err := config.ReadSetup(*configFilePtr)
		if err != nil {
			fmt.Printf("Could not read configuration file %s: %s", *configFilePtr, err.Error())
		} else {
			// Launch the CLI
			cli.Launch(&configuration)
		}
	} else {
		cli.Launch(nil)
	}
}
