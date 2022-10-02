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
	Short: "Subcommand for preparing the backend to be documented",
	Long:  `Subcommand for preparing the backend to be documented`,
	Run: func(cmd *cobra.Command, args []string) {
		log.Printf("cloud native backend designated, executing doc generation...")
	},
}

func init() {
	tfstateCmd.AddCommand(backendCmd)
}
