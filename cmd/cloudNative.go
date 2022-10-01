/*
Copyright © 2022 John Hession johhess@cdw.com

*/
package cmd

import (
	"fmt"
	"log"
	"os"

	generatedocs "github.com/paradocs-cli/paradocs/generatedocs"

	"github.com/spf13/cobra"
)

var (
	Provider string
)

// cloudNativeCmd represents the cloudNative command
var cloudNativeCmd = &cobra.Command{
	Use:   "cloud-native",
	Short: "A brief description of your command",
	Long: `cloud-native sets a cloud provider as the data source for the state pull
			for documentation:
			-Pulls in state data stored in cloud provided storage
			-Creates new data of just terraform state
'`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("cloud native backend designated, executing doc generation...")
		cloud, err := generatedocs.WhichCloudState(Provider, ProviderConfigs)
		if err != nil {
			log.Printf(err.Error())
			os.Exit(1)
		}
		generatedocs.WriteMarkdownCloudState(cloud)

		fmt.Printf("state pulled, if no errors were received markdown should be generated at %s\n", os.Getenv("PWD"))

	},
}

func init() {
	tfstateCmd.AddCommand(cloudNativeCmd)
	cloudNativeCmd.PersistentFlags().StringVarP(&Provider, "provider", "p", "", "cloud provider that stores your terraform backend for tfstate")
	cloudNativeCmd.PersistentFlags().StringVarP(&ProviderConfigs.Azure.StorageAccountName, "storage-account", "s", "", "the name of the Azure storage account with terraform state file")
	cloudNativeCmd.PersistentFlags().StringVarP(&ProviderConfigs.Azure.ContainerName, "container-name", "c", "", "the name of the Azure storage account container with terraform state file")
	cloudNativeCmd.PersistentFlags().StringVarP(&ProviderConfigs.Azure.BlobName, "blob-name", "b", "", "the name of the Azure storage account blob with terraform state file")
	cloudNativeCmd.PersistentFlags().StringVarP(&ProviderConfigs.Azure.SasToken, "sassy-token", "k", "", "Azure storage account SAS token for blob download via GET request")
	cloudNativeCmd.PersistentFlags().StringVarP(&ProviderConfigs.Aws.BucketName, "aws-bucket-name", "m", "", "the name of the AWS/GCP bucket with terraform state file")
	cloudNativeCmd.PersistentFlags().StringVarP(&ProviderConfigs.GoogleCloud.BucketName, "gcp-bucket-name", "w", "", "the name of the AWS/GCP bucket with terraform state file")
	cloudNativeCmd.PersistentFlags().StringVarP(&ProviderConfigs.Aws.Object, "aws-object", "o", "", "the name of the AWS/GCP object with terraform state file")
	cloudNativeCmd.PersistentFlags().StringVarP(&ProviderConfigs.GoogleCloud.ObjectName, "gcp-object", "i", "", "the name of the AWS/GCP object with terraform state file")
	cloudNativeCmd.PersistentFlags().StringVarP(&ProviderConfigs.Aws.Region, "region", "g", "", "the name of the AWS S3 region that contains the S3 bucket")
	cloudNativeCmd.PersistentFlags().StringVarP(&ProviderConfigs.Aws.AccessKey, "access-key", "y", "", "AWS access key for authenticating with developer credentials")
	cloudNativeCmd.PersistentFlags().StringVarP(&ProviderConfigs.Aws.SecretAccessKey, "secret-access-key", "v", "", "AWS secret access key for authenticating with developer credentials")
	cloudNativeCmd.PersistentFlags().StringVarP(&ProviderConfigs.Aws.SessionToken, "session-token", "u", "", "AWS session token for authenticating with developer credentials")
	cloudNativeCmd.PersistentFlags().StringVarP(&ProviderConfigs.GoogleCloud.Oauth2Token, "oauth2-token", "q", "", "OAuth 2 token for authentication for getting state stored in GCP")

	if Provider == "azure" {
		err := tfCloudCmd.MarkPersistentFlagRequired("storage-account")
		if err != nil {
			log.Printf(err.Error())
			os.Exit(1)
		}
		err = tfCloudCmd.MarkPersistentFlagRequired("container-name")
		if err != nil {
			log.Printf(err.Error())
			os.Exit(1)
		}
		err = tfCloudCmd.MarkPersistentFlagRequired("blob-name")
		if err != nil {
			log.Printf(err.Error())
			os.Exit(1)
		}
		err = tfCloudCmd.MarkPersistentFlagRequired("sassy-token")
		if err != nil {
			log.Printf(err.Error())
			os.Exit(1)
		}
	} else if Provider == "aws" {
		err := tfCloudCmd.MarkPersistentFlagRequired("aws-bucket-name")
		if err != nil {
			log.Printf(err.Error())
			os.Exit(1)
		}
		err = tfCloudCmd.MarkPersistentFlagRequired("aws-object")
		if err != nil {
			log.Printf(err.Error())
			os.Exit(1)
		}
		err = tfCloudCmd.MarkPersistentFlagRequired("region")
		if err != nil {
			log.Printf(err.Error())
			os.Exit(1)
		}
		err = tfCloudCmd.MarkPersistentFlagRequired("access-key")
		if err != nil {
			log.Printf(err.Error())
			os.Exit(1)
		}
		err = tfCloudCmd.MarkPersistentFlagRequired("secret-access-key")
		if err != nil {
			log.Printf(err.Error())
			os.Exit(1)
		}
		err = tfCloudCmd.MarkPersistentFlagRequired("session-token")
		if err != nil {
			log.Printf(err.Error())
			os.Exit(1)
		}
	} else if Provider == "gcp" {
		err := tfCloudCmd.MarkPersistentFlagRequired("gcp-bucket-name")
		if err != nil {
			log.Printf(err.Error())
			os.Exit(1)
		}
		err = tfCloudCmd.MarkPersistentFlagRequired("gcp-object")
		if err != nil {
			log.Printf(err.Error())
			os.Exit(1)
		}
		err = tfCloudCmd.MarkPersistentFlagRequired("oauth2-token")
		if err != nil {
			log.Printf(err.Error())
			os.Exit(1)
		}
	}
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// cloudNativeCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// cloudNativeCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
