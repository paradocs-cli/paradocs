package cmd

import (
	"github.com/spf13/cobra"
	"log"
)

// tfstateCmd represents the tfstate command
var tfstateCmd = &cobra.Command{
	Use:   "tfstate",
	Short: "Tells paradocs to start initiating documentation for terraform state",
	Long:  `Tells paradocs to start initiating documentation for terraform state `,
	Run: func(cmd *cobra.Command, args []string) {
		log.Printf("tfstate called")
	},
}

func init() {
	rootCmd.AddCommand(tfstateCmd)
}
