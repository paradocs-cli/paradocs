package cmd

import (
	generatedocs "github.com/paradocs-cli/paradocs/generatedocs"
	"log"

	"github.com/spf13/cobra"
)

var (
	ProviderConfigs generatedocs.StateProviders
)

// backendCmd represents the cloudNative command
var backendCmd = &cobra.Command{
	Use:   "backend",
	Short: "A brief description of your command",
	Long: `cloud-native sets a cloud provider as the data source for the state pull
			for documentation:
			-Pulls in state data stored in cloud provided storage
			-Creates new data of just terraform state
'`,
	Run: func(cmd *cobra.Command, args []string) {
		log.Printf("cloud native backend designated, executing doc generation...")
	},
}

func init() {
	tfstateCmd.AddCommand(backendCmd)
}
