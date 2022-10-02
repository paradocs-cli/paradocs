package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var (
	WorkingDir string
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "paradocs",
	Short: "paradocs is an automation tool for generating documentation for cloud native technologies",
	Long: `paradocs is a tool for generating documentation for cloud native technologies:
		--Use cloud docs and it's flags to specify supported doc creation languages
		--For example: paradocs code terraform, would generate documention for terraform code
		--All docs generated are in Markdown
		--Additional flags such as 'azdo push' will push documentation to specified repositories
		`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {}
