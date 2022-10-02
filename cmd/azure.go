package cmd

import (
	generatedocs "github.com/paradocs-cli/paradocs/generatedocs"
	"log"
	"os"

	"github.com/spf13/cobra"
)

// azureCmd represents the azure command
var azureCmd = &cobra.Command{
	Use:   "azure",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		log.Printf("azure called")
		cloud, err := generatedocs.WhichCloudState("azure", ProviderConfigs)
		if err != nil {
			log.Printf(err.Error())
			os.Exit(1)
		}
		generatedocs.WriteMarkdownCloudState(cloud)

		log.Printf("state pulled, if no errors were received markdown should be generated at %s\n", os.Getenv("PWD"))
	},
}

func init() {
	backendCmd.AddCommand(azureCmd)
	azureCmd.PersistentFlags().StringVarP(&ProviderConfigs.Azure.StorageAccountName, "storage-account", "s", "", "the name of the Azure storage account with terraform state file")
	err := azureCmd.MarkPersistentFlagRequired("storage-account")
	if err != nil {
		log.Printf(err.Error())
		os.Exit(1)
	}
	azureCmd.PersistentFlags().StringVarP(&ProviderConfigs.Azure.ContainerName, "container-name", "c", "", "the name of the Azure storage account container with terraform state file")
	err = azureCmd.MarkPersistentFlagRequired("container-name")
	if err != nil {
		log.Printf(err.Error())
		os.Exit(1)
	}
	azureCmd.PersistentFlags().StringVarP(&ProviderConfigs.Azure.BlobName, "blob-name", "b", "", "the name of the Azure storage account blob with terraform state file")
	err = azureCmd.MarkPersistentFlagRequired("blob-name")
	if err != nil {
		log.Printf(err.Error())
		os.Exit(1)
	}
	azureCmd.PersistentFlags().StringVarP(&ProviderConfigs.Azure.SasToken, "sassy-token", "t", "", "Azure storage account SAS token for blob download via GET request")
	err = azureCmd.MarkPersistentFlagRequired("sassy-token")
	if err != nil {
		log.Printf(err.Error())
		os.Exit(1)
	}
}
