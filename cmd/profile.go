package cmd

import (
	"fmt"
	"github.com/clearmatics/ion-cli/backend"
	"github.com/spf13/cobra"
)

// profileCmd represents the profile command
var (

	deleteFlag bool //flag to remove object

	/* PROFILES */
	profileCmd = &cobra.Command{
		Use:   "profiles",
		Short: "Manage the profiles on the cli",
		Long: `Add a profile`,
		Run: func(cmd *cobra.Command, args []string) {

			if len(args) != 1 {
				fmt.Println("Specify what profile name are you referring to")
				return
			}

			profileId := args[0]

			returnIfError(loadProfiles(profilesPath))
			returnIfError(loadConfig(configPath))

			if profiles.Exist(profileId) {

				if deleteFlag {
					fmt.Println("Deleting profile", profileId)
					delete (profiles, profileId)
				} else {
					fmt.Println("This profile already exist:\n", profiles[profileId])
				}

				profiles.Save(profilesPath)

			} else {
				fmt.Println("Creating a profile named", profileId)
				profiles[profileId] = *backend.InitProfile(profileId)
			}
		},
	}

	/* CHAINS OF A PROFILE */
	chainsCmd = &cobra.Command{
		Use:   "chains",
		Short: "Manage the chains of a profile",
		Long: ``,
		Run: func(cmd *cobra.Command, args []string) {

			if len(args) != 2 {
				fmt.Println("Specify profile and network identifiers you are referring to")
				return
			}

			profileId := args[0]
			networkId := args[1]

			returnIfError(loadProfiles(profilesPath))

			if profiles.Exist(profileId) {

				if deleteFlag {
					fmt.Println(fmt.Sprintf("Deleting chain %v in profile %v", networkId, profileId))

					delete (profiles[profileId].Chains, networkId)
				} else {
					fmt.Println(fmt.Sprintf("Creating chain %v in profile %v", networkId, profileId))
					returnIfError(loadConfig(configPath))

					network := backend.Network{}
					returnIfError(configs.UnmarshalKey("networks." + networkId, &network))

					profiles[profileId].Chains[networkId] = backend.Chain{
						Network:     network,
						Accounts:    make(map[string]backend.AccountInfo),
						Blocks:      nil,
						Transaction: backend.Transaction{},
					}
				}

				profiles.Save(profilesPath)
			} else {
				fmt.Println("This profile does not exists yet! Create it first")
				// or we can initialize it here
			}
		},
	}

	/* WALLETS OF A CHAIN OF A PROFILE */
	walletsCmd = &cobra.Command{
		Use:   "wallets",
		Short: "Add or delete wallet configuration to the profile",
		Long: `Add or delete wallet configurations to be used for a specific chain and profile`,
		Run: func(cmd *cobra.Command, args []string) {

			if len(args) != 3 {
				fmt.Println("Specify profile, chain and wallet identifier you are referring to")
				return
			}

			profileId := args[0]
			chainId := args[1]
			walletId := args[2]

			returnIfError(loadProfiles(profilesPath))

			if profiles.Exist(profileId){

				if profiles[profileId].Chains.Exist(chainId) {

					if deleteFlag {

						fmt.Println(fmt.Sprintf("Deleting wallet %v of chain %v in profile %v", walletId, chainId, profileId))
						delete (profiles[profileId].Chains[chainId].Accounts, walletId)

					} else {
						fmt.Println(fmt.Sprintf("Creating wallet %v of chain %v in profile %v", walletId, chainId, profileId))

						returnIfError(loadConfig(configPath))
						wallet := backend.AccountInfo{}
						returnIfError(configs.UnmarshalKey("accounts."+ walletId, &wallet))

						profiles[profileId].Chains[chainId].Accounts[walletId] = wallet
					}

					profiles.Save(profilesPath)

				} else {
					fmt.Println("This chain does not exists yet! Add it to the profile first")
				}

			} else {
				fmt.Println("This profile does not exists yet! Create it first")
			}

		},
	}
)

func init() {
	profileCmd.Flags().BoolVarP(&deleteFlag, "delete", "d", false, "Delete the specified profile")
	// TODO find a way to inherit parents flags
	chainsCmd.Flags().BoolVarP(&deleteFlag, "delete", "d", false, "Delete the specified chain")
	walletsCmd.Flags().BoolVarP(&deleteFlag, "delete", "d", false, "Delete the specified wallet")

	// tree of commands
	rootCmd.AddCommand(profileCmd)
	profileCmd.AddCommand(chainsCmd)
	profileCmd.AddCommand(walletsCmd)

}


