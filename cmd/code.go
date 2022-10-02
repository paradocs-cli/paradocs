package cmd

import (
	"github.com/spf13/cobra"
	"log"
)

// codeCmd represents the code command
var codeCmd = &cobra.Command{
	Use:   "code",
	Short: "subcommand for calling code documentation",
	Long: `subcommand code is used to direct paradocs to create
	documentation for code via terraform, helm, etc:
	-- Currently functional for HCL/terraform only 
`,
	Run: func(cmd *cobra.Command, args []string) {
		log.Printf("Executing doc build for code")
	},
}

func init() {
	rootCmd.AddCommand(codeCmd)
}

func NewCodeCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "code",
		Short: "testing for code ",
		Long:  `testing for functionality of code subcommand`,
		Run: func(cmd *cobra.Command, args []string) {
			log.Printf("Executing doc build for code")
		},
	}
}
