/*
Copyright © 2022 John Hession
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

// gitlabCmd represents the gitlabdocs command
var gitlabCmd = &cobra.Command{
	Use:   "gitlab",
	Short: "this command generates docs for gitlab",
	Long: `This sub command generates your documentation for gitlab:
			--Currently only terraform is supported
`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("gitlab docs being generated, standby....")
		err := docbuilder.BuildGitLabDocs(builder)
		if err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	platformCmd.AddCommand(gitlabCmd)
	gitlabCmd.PersistentFlags().StringSliceVarP(&builder.ProjectIds, "project-ids", "p", []string{}, "comma separated list of project ids for gitlab documentation")
	err := gitlabCmd.MarkPersistentFlagRequired("project-ids")
	if err != nil {
		log.Printf(err.Error())
		os.Exit(1)
	}
	gitlabCmd.PersistentFlags().StringVarP(&builder.UserName, "user-name", "u", "", "username for authentication with gitlab")
	err = gitlabCmd.MarkPersistentFlagRequired("user-name")
	if err != nil {
		log.Printf(err.Error())
		os.Exit(1)
	}
	gitlabCmd.PersistentFlags().StringVarP(&builder.Token, "token", "t", "", "your PAT for authentication with gitlab")
	err = gitlabCmd.MarkPersistentFlagRequired("token")
	if err != nil {
		log.Printf(err.Error())
		os.Exit(1)
	}

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// gitlabCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// gitlabCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
