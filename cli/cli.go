// Copyright (c) 2018 Clearmatics Technologies Ltd
package cli

import (
	"fmt"
	"github.com/abiosoft/ishell"
	"github.com/clearmatics/ion-cli/cli/cmd"
	"github.com/clearmatics/ion-cli/cli/core"
	"github.com/clearmatics/ion-cli/config"
)

func printWelcome() {
	// display welcome info.
	fmt.Println("===============================================================")
	fmt.Print("Ion Command Line Interface\n\n")
	fmt.Println("Use 'help' to list commands")
	fmt.Println("===============================================================")
}

// Launch - definition of commands and creates the interface
func Launch(setup *config.Setup) {
	// by default, new shell includes 'exit', 'help' and 'clear' commands.
	shell := ishell.New()
	session := core.InitSession()

	if setup != nil {
		// Add all accounts in config to memory
		for _, account := range setup.Accounts {
			user, err := config.InitUser(account.Keyfile, account.Password)
			if err != nil {
				fmt.Printf("Setup Failed: Adding Account %s from configuration failed %s", account.Name, err.Error())
				return
			}
			session.Accounts[account.Name] = &user
		}

		// Compile and add all contract instances to memory
		for _, configContract := range setup.Contracts {
			compiledContract, err := core.AddCompilerAndCompileContract(session, configContract.File)
			if err != nil {
				fmt.Printf("Setup Failed: Compiling contract %s from configuration failed: %s", configContract.Name, err.Error())
				return
			}
			session.Contracts[configContract.Name] = compiledContract
		}

		// Compile and add all contract instances to memory
		for _, configNetwork := range setup.Networks {
			client, err := core.GetClient(configNetwork.Uri)
			if err != nil {
				fmt.Printf("Could not connect to client %s\n", configNetwork.Name)
				return
			}

			session.Networks[configNetwork.Name] = client
		}
	}

	// Add commands
	for _, command := range cmd.CoreCommands(session) {
		shell.AddCmd(command)
	}

	for _, command := range cmd.CliqueCommands(session) {
		shell.AddCmd(command)
	}

	printWelcome()
	shell.Run()
	session.Close()
	shell.Close()
}
