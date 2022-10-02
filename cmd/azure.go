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
	Short: "Documents state for the azure storage account backend:",
	Long: `Documents state for the azure storage account backend:		
		--> Documents state of resources
		--> Usage: paradocs tfstate backend azure --storage-account <storage account> --container-name <container name> --blob-name <blob name> --sassy-token <token> 
		--> Usage Short: paradocs tfstate backend azure -s <storage account> -c <container name> -b <blob name> -t <token> 
		--> Make contributions at paradocs-cli: https://github.com/paradocs-cli`,
	Run: func(cmd *cobra.Command, args []string) {
		log.Printf("azure called")
		cloud, err := generatedocs.WhichCloudState("azure", ProviderConfigs)
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
	azureCmd.PersistentFlags().StringVarP(&OutDir, "out-dir", "o", ".", "Directory that markdown documentation should be exported to...")
}
