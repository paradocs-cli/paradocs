package cmd

import (
	generatedocs "github.com/paradocs-cli/paradocs/generatedocs"
	"log"
	"os"

	"github.com/spf13/cobra"
)

// tfcCmd represents the tfc command
var tfcCmd = &cobra.Command{
	Use:   "tfc",
	Short: "tfc calls for terraform cloud as the backend for state documentation",
	Long: `tfc calls for terraform cloud as the backend for state documentation:
		--> Documents state of resources
		--> Documents workspaces as well
		--> Usage: paradocs tfstate backend tfc --api-token <token> --workspace-id <workspace id> 
		--> Usage Short: paradocs tfstate backend tfc -t <token> -w <workspace id> 
		--> Make contributions at paradocs-cli: https://github.com/paradocs-cli`,
	Run: func(cmd *cobra.Command, args []string) {
		log.Printf("tfc called")
		cloud, err := generatedocs.WhichCloudState("tfc", ProviderConfigs)
		if err != nil {
			log.Printf(err.Error())
			os.Exit(1)
		}
		err = cloud.WriteMarkdownCloudState(OutDir)
		if err != nil {
			log.Printf(err.Error())
			os.Exit(1)
		}

		log.Printf("state pulled, if no errors were received markdown should be generated at %s\n", os.Getenv("PWD"))
	},
}

func init() {
	backendCmd.AddCommand(tfcCmd)
	tfcCmd.PersistentFlags().StringVarP(&ProviderConfigs.TerraformCloud.ApiToken, "api-token", "t", "", "token needed for authentication with TFC api")
	err := tfcCmd.MarkPersistentFlagRequired("api-token")
	if err != nil {
		log.Printf(err.Error())
		os.Exit(1)
	}
	tfcCmd.PersistentFlags().StringVarP(&ProviderConfigs.TerraformCloud.WorkspaceId, "workspace-id", "w", "", "id needed for authentication to workspace for TFC api")
	err = tfcCmd.MarkPersistentFlagRequired("workspace-id")
	if err != nil {
		log.Printf(err.Error())
		os.Exit(1)
	}
	tfcCmd.PersistentFlags().StringVarP(&OutDir, "out-dir", "o", ".", "Directory that markdown documentation should be exported to...")
	tfcCmd.PersistentFlags().BoolP("generate-workspace", "g", false, "optional flag to generate documentation for the terraform cloud workspace as well")
}
