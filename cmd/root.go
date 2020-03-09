package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)


var(
	// flags can be persistent or local - more down below
	Verbose bool

	rootCmd = &cobra.Command{
		Use:   "ion",
		Short: "Cross-chain framework tool",
		Long: "Ion is a system and function-agnostic framework for building cross-interacting smart contracts between blockchains and/or systems",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Hi from the ION cli. Type ion-cli help to display the help")
		},
		Args: func(cmd *cobra.Command, args []string) error {
			// this to validate positional arguments
			return nil
		},
	}
)

func Execute() {
	// start the app
	rootCmd.Execute()
}

func init(){
	// add global flags
	rootCmd.PersistentFlags().BoolVarP(&Verbose, "verbose", "v", false, "verbose output")

	// add local flags with rootCmd.Flags()

	// initialize session with viper and configs

}

func er(msg interface{}) {
	fmt.Println("Error:", msg)
	os.Exit(1)
}

