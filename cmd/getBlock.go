package cmd

import (
	"fmt"
	"github.com/clearmatics/ion-cli/backend"
	"github.com/spf13/viper"

	"github.com/spf13/cobra"
)

// TODO getBlock -number/hash -encoded
var getBlockCmd = &cobra.Command{
	Use:   "getBlock",
	Short: "Allow to retrieve a block through a rpc call",
	Long: `Allow to retrieve a block through a rpc call, either by number or by hash, rlp encoded or as object`,
	Run: func(cmd *cobra.Command, args []string) {

		eth, err := backend.GetClient(viper.GetString("rpc"))
		if err != nil {
			fmt.Println(err)
			return
		}

		block, _, err := eth.GetBlockByNumber("latest")
		fmt.Println(block)
	},
}

func init() {
	// TODO user pass only flags related to this specific command (block number). all the needed params are in viper
	rootCmd.AddCommand(getBlockCmd)

}
