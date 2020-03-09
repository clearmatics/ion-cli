package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)


var(
	// flags can be persistent or local - more down below
	Verbose bool
	Configs *viper.Viper
	Session *viper.Viper

	ConfigPath string

	rootCmd = &cobra.Command{
		Use:   "ion-cli",
		Short: "Cross-chain framework tool",
		Long: "Ion is a system and function-agnostic framework for building cross-interacting smart contracts between blockchains and/or systems",

		Run: func(cmd *cobra.Command, args []string) {
			//fmt.Println("Hi from the ION cli. Type ion-cli help to display the help")

			fmt.Println("root called")

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
	fmt.Println("root init called")

	// add global flags
	rootCmd.PersistentFlags().BoolVarP(&Verbose, "verbose", "v", false, "verbose output")
	rootCmd.Flags().StringVarP(&ConfigPath, "config", "c", "./config/test.json", "Config file to populate the session with")

	// add local flags with rootCmd.Flags()

	// initialize empty Session viper object
	Session = viper.New()

	// initialize default configs from file
	Configs = viper.New()

	Configs.SetConfigFile(ConfigPath)
	err := Configs.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}

	// add all commands
	rootCmd.AddCommand(versionCmd)
	//rootCmd.AddCommand(sessionCmd)
}


