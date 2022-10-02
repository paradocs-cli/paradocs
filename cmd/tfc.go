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
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		log.Printf("tfc called")
		cloud, err := generatedocs.WhichCloudState("tfc", ProviderConfigs)
		if err != nil {
			log.Printf(err.Error())
			os.Exit(1)
		}
		generatedocs.WriteMarkdownCloudState(cloud)

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
	tfcCmd.PersistentFlags().BoolP("generate-workspace", "g", false, "optional flag to generate documentation for the terraform cloud workspace as well")
}
