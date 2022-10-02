package cmd

import (
	"log"

	generatedocs "github.com/paradocs-cli/paradocs/generatedocs"

	"github.com/spf13/cobra"
)

var (
	OutDir string
)

// terraformCmd represents the terraform command
var terraformCmd = &cobra.Command{
	Use:   "terraform",
	Short: "generates documentation for your Terraform configs",
	Long: `sub command terraform specifies that the documentation you 
	want generated is for terraform/hcl files:
		--> Documents Variables
		--> Documents Resources
		--> Documents Modules
		--> Documents Outputs
		--> Documents Providers
		--> Make contributions at paradocs-cli: https://github.com/paradocs-cli
`,
	Run: func(cmd *cobra.Command, args []string) {
		log.Printf("generating Terraform docs...")
		data, err := generatedocs.GetData(WorkingDir)
		if err != nil {
			log.Fatalf(err.Error())
		}
		err = generatedocs.WriteMarkdownTerra(data, OutDir)
		if err != nil {
			log.Fatalf(err.Error())
		}
	},
}

func init() {
	codeCmd.AddCommand(terraformCmd)
	terraformCmd.PersistentFlags().StringVarP(&WorkingDir, "working-dir", "w", ".", "Working assets for running doc generation, defaults to '.'")
	terraformCmd.PersistentFlags().StringVarP(&OutDir, "out-dir", "o", ".", "Directory that markdown documentation should be exported to...")
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
			log.Printf("generating Terraform docs...")
			data, err := generatedocs.GetData(w)
			if err != nil {
				log.Fatalf(err.Error())
			}
			err = generatedocs.WriteMarkdownTerra(data, ".")
			if err != nil {
				log.Fatalf(err.Error())
			}
			log.Printf("terraform docs generated....")
		},
	}
}
