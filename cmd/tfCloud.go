/*
Copyright © 2022 John Hession johhess@cdw.com
*/
package cmd

import (
	"fmt"
	"log"
	"os"

	generatedocs "github.com/paradocs-cli/generatedocs"

	"github.com/spf13/cobra"
)

var (
	ProviderConfigs generatedocs.StateProviders
)

// tfCloudCmd represents the tfCloud command
var tfCloudCmd = &cobra.Command{
	Use:   "tf-cloud",
	Short: "Pulls in state data from terraform cloud for documentation",
	Long: `tf-cloud sets terraform cloud as the data source for the state pull
			for documentation:
			-Pulls in 'current-state-version
			-Creates new data of just terraform state
			-Optionally can generate docs for the tfc workspace as well
'`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("terraform cloud state being pulled, docs being generated")
		cloud, err := generatedocs.WhichCloudState("tfc", ProviderConfigs)
		if err != nil {
			log.Printf(err.Error())
			os.Exit(1)
		}
		generatedocs.WriteMarkdownCloudState(cloud)

		fmt.Printf("state pulled, if now errors were received markdown should be generated at %s\n", os.Getenv("PWD"))

	},
}

func init() {
	tfstateCmd.AddCommand(tfCloudCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	tfCloudCmd.PersistentFlags().StringVarP(&ProviderConfigs.TerraformCloud.ApiToken, "api-token", "a", "", "token needed for authentication with TFC api")
	err := tfCloudCmd.MarkPersistentFlagRequired("api-token")
	if err != nil {
		log.Printf(err.Error())
		os.Exit(1)
	}
	tfCloudCmd.PersistentFlags().StringVarP(&ProviderConfigs.TerraformCloud.WorkspaceId, "workspace-id", "i", "", "id needed for authentication to workspace for TFC api")
	err = tfCloudCmd.MarkPersistentFlagRequired("workspace-id")
	if err != nil {
		log.Printf(err.Error())
		os.Exit(1)
	}
	tfCloudCmd.PersistentFlags().BoolP("generate-workspace", "g", false, "optional flag to generate documentation for the terraform cloud workspace as well")
	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// tfCloudCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
