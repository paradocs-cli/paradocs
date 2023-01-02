/*
Copyright © 2022 John J. Hession

*/
package cmd

import (
	"github.com/fatih/color"
	"log"

	"github.com/spf13/cobra"
)

// packerCmd represents the packer command
var packerCmd = &cobra.Command{
	Use:   "packer",
	Short: "packer flags tells paradocs to document packer HCL code",
	Long: `paradocs sub command packer specifies that the documentation you 
	want generated is for packer/hcl files:
		--> Documents Variables
		--> Documents Sources
		--> Documents Builds
		--> Make contributions at paradocs-cli: https://github.com/paradocs-cli/paradocs`,
	Run: func(cmd *cobra.Command, args []string) {
		if OutDir == "." {
			log.Printf(color.GreenString("[INFO] 🧙‍generating Packer docs for output in directory: root🧙‍"))
		} else {
			log.Printf(color.GreenString("[INFO] 🧙‍generating Packer docs for output in directory: %s🧙‍", OutDir))
		}

		if OutDir == "." {
			log.Printf(color.GreenString("[INFO] 🧙‍writing Packer docs to directory: root🧙‍"))
		} else {
			log.Printf(color.GreenString("[INFO] 🧙‍writing Packer docs to directory: %s🧙‍", OutDir))
		}
	},
}

func init() {
	packerCmd.PersistentFlags().StringVar(&WorkingDir, "working-dir", ".", "Working assets for running doc generation, defaults to '.'")
	packerCmd.PersistentFlags().StringVar(&OutDir, "out-dir", ".", "Directory that markdown documentation should be exported to...")
}
