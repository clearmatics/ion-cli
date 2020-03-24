package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/clearmatics/ion-cli/backend"
	"github.com/spf13/cobra"
	"github.com/spf13/cobra/doc"
	"github.com/spf13/viper"
	"io/ioutil"
	"os"
	"time"
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
	deleteFlag bool

	// global variable to all commands
	activeProfile backend.Profile
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

			// choose profile to use
			initProfile()
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
	rootCmd.PersistentFlags().StringVarP(&configPath, "config", "c", "./config/config-test.json", "Configs file path")
	rootCmd.PersistentFlags().StringVarP(&sessionPath, "session", "s", "./config/session-test.json", "Session file path")
	rootCmd.PersistentFlags().StringVarP(&profilesPath, "profiles", "", "./config/profiles-test.json", "Profiles file path")

	rootCmd.Flags().StringVarP(&outputDir, "outputDir", "o", "./docs", "The output directory the docs will be written into")
	rootCmd.Flags().BoolVarP(&docFlag, "docgen", "", false, "Generate documentation of the whole command tree")

	rootCmd.Flags().StringVarP(&profileName, "profile", "p", "", "The profile name the configs will be taken from")
}

func initProfile() {

	returnIfError(loadProfiles(profilesPath))

	// if profile flag is set use that if it's a valid profile
	if profiles.Exist(profileName){
		fmt.Println("Using profile", profileName, "from the flag")

		activeProfile = profiles[profileName]
	} else {
		// check if a session is active
		b, _ := ioutil.ReadFile(sessionPath)
		returnIfError(json.Unmarshal(b, &session))

		if session.IsValid(timeoutSec) && profiles.Exist(session.Profile) {
			fmt.Println("Loading profile", session.Profile, "from the session")

			session.LastAccess = int(time.Now().Unix())
			session.Save(sessionPath)

			activeProfile = profiles[session.Profile]
		}
	}

	// TODO how about no profiles?
	fmt.Println(activeProfile)
}


