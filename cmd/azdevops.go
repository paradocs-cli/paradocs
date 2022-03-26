/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// azdevopsCmd represents the azdevops command
var azdevopsCmd = &cobra.Command{
	Use:   "azdevops",
	Short: "azdevops designates Azure DevOps as the project platform for documentation",
	Long: `azdevops designates Azure DevOps as the project platform for documentation:
			-- Generates documentation for project 			
			-- Generates documentation for project repos  			
			-- Generates documentation for project pipelines 			
			-- Generates documentation for project sprints 			
`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("azdevops called...")
	},
}

func init() {
	platformCmd.AddCommand(azdevopsCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// azdevopsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// azdevopsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
