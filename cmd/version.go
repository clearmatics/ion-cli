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
		fmt.Println("versin called ")


		fmt.Println("Using default configs:", Configs.Get("rpc-to"))
	},
}

func init(){
	fmt.Println("version init called")

	// add command to root
	//rootCmd.AddCommand(versionCmd)
}