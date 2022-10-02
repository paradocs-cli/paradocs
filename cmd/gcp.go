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
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
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
	gcpCmd.PersistentFlags().StringVarP(&OutDir, "out-dir", "o", ".", "Directory that markdown documentation should be exported to...")
}
