package cmd

import (
	"fmt"
	"github.com/clearmatics/ion-cli/backend"
	"github.com/spf13/cobra"
	"github.com/spf13/cobra/doc"
	"github.com/spf13/viper"
	"os"
)

// TODO add log with verbosity level
// TODO we might move also root configs command to backend
var(
	// flags
	Verbose bool
	docFlag bool
	outputDir string
	sessionPath string
	configPath string
	profilesPath string
	profileName string
	rpcURL string
	chainType string
	forceFlag bool


	activeProfile backend.Profile
	activeChain backend.Chain
	profiles backend.Profiles

	session backend.Session
	configs *viper.Viper
	timeoutSec =  3600

	rootCmd = &cobra.Command{
		Use:   "ion-cli",
		Short: "Cross-chain framework tool",
		Long: "Ion is a system and function-agnostic framework for building cross-interacting smart contracts between blockchains and/or systems",

		Run: func(cmd *cobra.Command, args []string) {

			// generate docs
			if docFlag {
				fmt.Println("Generating documentation at", outputDir)

				if _, err := os.Stat(outputDir); err != nil {
					fmt.Println("Output path didn't exist. Creating the folder..")
					returnIfError(os.Mkdir(outputDir, 0777))
				}

				returnIfError(doc.GenMarkdownTree(cmd, outputDir))
				fmt.Println("Success!")
				return
			}

		},

		Args: func(cmd *cobra.Command, args []string) error {
			// this to validate positional arguments
			return nil
		},
	}
)

func Execute() {
	// start the cli app
	rootCmd.Execute()
}

func init(){
	// flags
	rootCmd.PersistentFlags().BoolVarP(&Verbose, "verbose", "v", false, "verbose output")
	rootCmd.PersistentFlags().StringVarP(&configPath, "config", "", "./config/config-test.json", "Configs file path")
	rootCmd.PersistentFlags().StringVarP(&sessionPath, "session", "", "./config/session-test.json", "Session file path")
	rootCmd.PersistentFlags().StringVarP(&profilesPath, "profiles", "", "./config/profiles-test.json", "Profiles file path")

	rootCmd.Flags().StringVarP(&outputDir, "outputDir", "o", "./docs", "The output directory the docs will be written into")
	rootCmd.Flags().BoolVarP(&docFlag, "docgen", "", false, "Generate documentation of the whole command tree")

	rootCmd.PersistentFlags().StringVarP(&profileName, "profile", "p", "", "The profile name the configs will be taken from")

	rootCmd.PersistentFlags().BoolVarP(&forceFlag, "force", "f", false, "Overwrites objects that already exist")

	return
}

func NewRootCmd() *cobra.Command {
	// internally calls also the init function of the command
	return rootCmd
}


