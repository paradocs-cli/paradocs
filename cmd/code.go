/*
Copyright © 2022 John Hession
*/
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
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
		fmt.Printf("Executing doc build for code")
	},
}

func init() {
	rootCmd.AddCommand(codeCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// codeCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// codeCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func NewCodeCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "code",
		Short: "testing for code ",
		Long:  `testing for functionality of code subcommand`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("Executing doc build for code")
		},
	}
}
