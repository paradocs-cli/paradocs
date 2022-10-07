package cmd

import (
	generatedocs "github.com/paradocs-cli/paradocs/generatedocs"
	"log"
	"os"

	"github.com/spf13/cobra"
)

// gcpCmd represents the gcp command
var gcpCmd = &cobra.Command{
	Use:   "gcp",
	Short: "Documents state for backend in google cloud:",
	Long: `Documents state for backend in google cloud:
		--> Documents state of resources
		--> Usage: paradocs tfstate backend gcp --gcp-bucket-name <bucket name> --gcp-object-name <object name> --oauth2-token <token>
		--> Usage Short: paradocs tfstate backend gcp -b <bucket name> -o <object name> -t <token>
		--> Make contributions at paradocs-cli: https://github.com/paradocs-cli
`,
	Run: func(cmd *cobra.Command, args []string) {
		log.Printf("gcp called")
		cloud, err := generatedocs.WhichCloudState("gcp", ProviderConfigs)
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
	backendCmd.AddCommand(gcpCmd)
	gcpCmd.PersistentFlags().StringVarP(&ProviderConfigs.GoogleCloud.BucketName, "gcp-bucket-name", "b", "", "the name of the AWS/GCP bucket with terraform state file")
	err := gcpCmd.MarkPersistentFlagRequired("gcp-bucket-name")
	if err != nil {
		log.Printf(err.Error())
		os.Exit(1)
	}
	gcpCmd.PersistentFlags().StringVarP(&ProviderConfigs.GoogleCloud.ObjectName, "gcp-object", "o", "", "the name of the AWS/GCP object with terraform state file")
	err = gcpCmd.MarkPersistentFlagRequired("gcp-object")
	if err != nil {
		log.Printf(err.Error())
		os.Exit(1)
	}
	gcpCmd.PersistentFlags().StringVarP(&ProviderConfigs.GoogleCloud.Oauth2Token, "oauth2-token", "t", "", "OAuth 2 token for authentication for getting state stored in GCP")
	err = gcpCmd.MarkPersistentFlagRequired("oauth2-token")
	if err != nil {
		log.Printf(err.Error())
		os.Exit(1)
	}
	gcpCmd.PersistentFlags().StringVar(&OutDir, "out-dir", ".", "Directory that markdown documentation should be exported to...")
}
