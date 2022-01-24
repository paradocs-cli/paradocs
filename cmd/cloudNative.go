/*
Copyright © 2022 John Hession johhess@cdw.com

*/
package cmd

import (
	"fmt"
	generatedocs "github.com/johhess40/generatedocs"
	"log"
	"os"

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

		fmt.Printf("state pulled, if now errors were received markdown should be generated at %s\n", os.Getenv("PWD"))

	},
}

func init() {
	tfstateCmd.AddCommand(cloudNativeCmd)
	cloudNativeCmd.PersistentFlags().StringVarP(&Provider, "provider", "p", "", "cloud provider that stores your terraform backend for tfstate")
	cloudNativeCmd.PersistentFlags().String("storage-account", "", "the name of the Azure storage account with terraform state file")
	cloudNativeCmd.PersistentFlags().String("container-name", "", "the name of the Azure storage account container with terraform state file")
	cloudNativeCmd.PersistentFlags().String("blob-name", "", "the name of the Azure storage account blob with terraform state file")
	cloudNativeCmd.PersistentFlags().String("sassy-token", "", "Azure storage account SAS token for blob download via GET request")
	cloudNativeCmd.PersistentFlags().String("bucket-name", "", "the name of the AWS/GCP bucket with terraform state file")
	cloudNativeCmd.PersistentFlags().String("object", "", "the name of the AWS/GCP object with terraform state file")
	cloudNativeCmd.PersistentFlags().String("region", "", "the name of the AWS S3 region that contains the S3 bucket")
	cloudNativeCmd.PersistentFlags().String("access-key", "", "AWS access key for authenticating with developer credentials")
	cloudNativeCmd.PersistentFlags().String("secret-access-key", "", "AWS secret access key for authenticating with developer credentials")
	cloudNativeCmd.PersistentFlags().String("session-token", "", "AWS session token for authenticating with developer credentials")
	cloudNativeCmd.PersistentFlags().String("oauth2-token", "", "OAuth 2 token for authentication for getting state stored in GCP")

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
	ProviderConfigs.Azure.StorageAccountName = cloudNativeCmd.Flags().Lookup("storage-account").Value.String()
	ProviderConfigs.Azure.StorageAccountName = cloudNativeCmd.Flags().Lookup("container-name").Value.String()
	ProviderConfigs.Azure.StorageAccountName = cloudNativeCmd.Flags().Lookup("blob-name").Value.String()
	ProviderConfigs.Azure.StorageAccountName = cloudNativeCmd.Flags().Lookup("sassy-token").Value.String()
	} else if Provider == "aws" {
		err := tfCloudCmd.MarkPersistentFlagRequired("bucket-name")
		if err != nil {
			log.Printf(err.Error())
			os.Exit(1)
		}
		err = tfCloudCmd.MarkPersistentFlagRequired("object")
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
		ProviderConfigs.Aws.BucketName = cloudNativeCmd.Flags().Lookup("bucket-name").Value.String()
		ProviderConfigs.Aws.Object = cloudNativeCmd.Flags().Lookup("object").Value.String()
		ProviderConfigs.Aws.Region = cloudNativeCmd.Flags().Lookup("region").Value.String()
		ProviderConfigs.Aws.AccessKey = cloudNativeCmd.Flags().Lookup("access-key").Value.String()
		ProviderConfigs.Aws.SecretAccessKey = cloudNativeCmd.Flags().Lookup("secret-access-key").Value.String()
		ProviderConfigs.Aws.SessionToken = cloudNativeCmd.Flags().Lookup("session-token").Value.String()
	} else if Provider == "gcp"{
		err := tfCloudCmd.MarkPersistentFlagRequired("bucket-name")
		if err != nil {
			log.Printf(err.Error())
			os.Exit(1)
		}
		err = tfCloudCmd.MarkPersistentFlagRequired("object")
		if err != nil {
			log.Printf(err.Error())
			os.Exit(1)
		}
		err = tfCloudCmd.MarkPersistentFlagRequired("oauth2-token")
		if err != nil {
			log.Printf(err.Error())
			os.Exit(1)
		}
		ProviderConfigs.GoogleCloud.BucketName = cloudNativeCmd.Flags().Lookup("bucket-name").Value.String()
		ProviderConfigs.GoogleCloud.ObjectName = cloudNativeCmd.Flags().Lookup("object").Value.String()
		ProviderConfigs.GoogleCloud.Oauth2Token = cloudNativeCmd.Flags().Lookup("oauth2-token").Value.String()
	}
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// cloudNativeCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// cloudNativeCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
