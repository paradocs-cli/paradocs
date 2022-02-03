/*
Copyright © 2022 John Hession johhess@cdw.com

*/
package cmd

import (
	"fmt"
	generatedocs "github.com/johhess40/generatedocs"
	"log"

	"github.com/spf13/cobra"
)

// terraformCmd represents the terraform command
var terraformCmd = &cobra.Command{
	Use:   "terraform",
	Short: "generates documentation for your Terraform configs",
	Long: `sub command terraform specifies that the documentation you 
	want generated is for terraform/hcl files:
	--Documents variables
	--Documents resources
	--Documents data resources
	--Documents providers
	--Documents outputs
	--Upcoming:
		--Added support for required providers 
		--Docs generated based on comments, similar to go.pkg.dev
		--Automated code snippets for examples
`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("generating Terraform docs...")
		data, err := generatedocs.GetData(WorkingDir)
		if err != nil {
			log.Fatalf(err.Error())
		}
		generatedocs.WriteMarkdownTerra(data)
	},
}

func init() {
	codeCmd.AddCommand(terraformCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// terraformCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// terraformCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	terraformCmd.PersistentFlags().StringVarP(&WorkingDir, "working-dir", "w", ".", "Working assets for running doc generation, defaults to '.'")
}

func NewTerraformCmd(w string) *cobra.Command {
	return &cobra.Command{
		Use:   "terraform",
		Short: "generates documentation for your Terraform configs",
		Long: `sub command terraform specifies that the documentation you 
	want generated is for terraform/hcl files:
	--Documents variables
	--Documents resources
	--Documents data resources
	--Documents providers
	--Documents outputs
	--Upcoming:
		--Added support for required providers 
		--Docs generated based on comments, similar to go.pkg.dev
		--Automated code snippets for examples
`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("generating Terraform docs...")
			data, err := generatedocs.GetData(w)
			if err != nil {
				log.Fatalf(err.Error())
			}
			generatedocs.WriteMarkdownTerra(data)
			fmt.Println("terraform docs generated....")
		},
	}
}
