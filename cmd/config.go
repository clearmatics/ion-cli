package cmd

import (
	"github.com/spf13/cobra"
)

// configsCmd represents the root command to manage configs
var (

	/* CONFIGS */
	configsCmd = &cobra.Command{
		Use:   "config",
		Short: "",
		// this will be run before each config subcommand to avoid loading the configs in multiple places
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			// load configs
			returnIfError(loadConfig(configPath))
			return nil
		},
	}

)

func init() {
	rootCmd.AddCommand(configsCmd)
}
