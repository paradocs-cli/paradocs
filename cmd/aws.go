package cmd

import (
	generatedocs "github.com/paradocs-cli/paradocs/generatedocs"
	"log"
	"os"

	"github.com/spf13/cobra"
)

// awsCmd represents the aws command
var awsCmd = &cobra.Command{
	Use:   "aws",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		log.Printf("aws called")
		log.Printf("cloud native backend designated, executing doc generation...")
		cloud, err := generatedocs.WhichCloudState("aws", ProviderConfigs)
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
	backendCmd.AddCommand(awsCmd)
	awsCmd.PersistentFlags().StringVarP(&ProviderConfigs.Aws.BucketName, "aws-bucket-name", "b", "", "the name of the AWS/GCP bucket with terraform state file")
	err := awsCmd.MarkPersistentFlagRequired("aws-bucket-name")
	if err != nil {
		log.Printf(err.Error())
		os.Exit(1)
	}
	awsCmd.PersistentFlags().StringVarP(&ProviderConfigs.Aws.Object, "aws-object", "o", "", "the name of the AWS/GCP object with terraform state file")
	err = awsCmd.MarkPersistentFlagRequired("aws-object")
	if err != nil {
		log.Printf(err.Error())
		os.Exit(1)
	}
	awsCmd.PersistentFlags().StringVarP(&ProviderConfigs.Aws.Region, "region", "r", "", "the name of the AWS S3 region that contains the S3 bucket")
	err = awsCmd.MarkPersistentFlagRequired("region")
	if err != nil {
		log.Printf(err.Error())
		os.Exit(1)
	}
	awsCmd.PersistentFlags().StringVarP(&ProviderConfigs.Aws.AccessKey, "access-key", "a", "", "AWS access key for authenticating with developer credentials")
	err = awsCmd.MarkPersistentFlagRequired("access-key")
	if err != nil {
		log.Printf(err.Error())
		os.Exit(1)
	}
	awsCmd.PersistentFlags().StringVarP(&ProviderConfigs.Aws.SecretAccessKey, "secret-access-key", "k", "", "AWS secret access key for authenticating with developer credentials")
	err = awsCmd.MarkPersistentFlagRequired("secret-access-key")
	if err != nil {
		log.Printf(err.Error())
		os.Exit(1)
	}
	awsCmd.PersistentFlags().StringVarP(&ProviderConfigs.Aws.SessionToken, "session-token", "t", "", "AWS session token for authenticating with developer credentials")
	err = awsCmd.MarkPersistentFlagRequired("session-token")
	if err != nil {
		log.Printf(err.Error())
		os.Exit(1)
	}
	awsCmd.PersistentFlags().StringVarP(&OutDir, "out-dir", "o", ".", "Directory that markdown documentation should be exported to...")
}
