package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

// descriptions
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of Ion",
	Run: func(cmd *cobra.Command, args []string) {
		// TODO set it dynamic
		fmt.Println("Ion version 1.0.0")
	},
}

func init(){
	// add command to root
	rootCmd.AddCommand(versionCmd)
}