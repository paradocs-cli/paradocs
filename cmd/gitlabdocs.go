/*
Copyright © 2022 John Hession <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"github.com/paradocs-cli/docbuilder"
	"github.com/spf13/cobra"
	"log"
	"os"
)

var (
	builder docbuilder.GitlabData
)

// gitlabdocsCmd represents the gitlabdocs command
var gitlabdocsCmd = &cobra.Command{
	Use:   "gitlabdocs",
	Short: "this command generates docs for gitlab",
	Long: `This sub command generates your documentation for gitlab:
			--Currently only terraform is supported
`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("gitlabdocs being generated, standby....")
		err := docbuilder.BuildGitLabDocs(builder)
		if err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(gitlabdocsCmd)
	gitlabdocsCmd.PersistentFlags().StringSliceVarP(&builder.ProjectIds, "project-ids", "p", []string{}, "comma separated list of project ids for gitlab documentation")
	err := gitlabdocsCmd.MarkPersistentFlagRequired("project-ids")
	if err != nil {
		log.Printf(err.Error())
		os.Exit(1)
	}
	gitlabdocsCmd.PersistentFlags().StringVarP(&builder.UserName, "user-name", "u", "", "username for authentication with gitlab")
	err = gitlabdocsCmd.MarkPersistentFlagRequired("user-name")
	if err != nil {
		log.Printf(err.Error())
		os.Exit(1)
	}
	gitlabdocsCmd.PersistentFlags().StringVarP(&builder.Token, "token", "t", "", "your PAT for authentication with gitlab")
	err = gitlabdocsCmd.MarkPersistentFlagRequired("token")
	if err != nil {
		log.Printf(err.Error())
		os.Exit(1)
	}


	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// gitlabdocsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// gitlabdocsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
