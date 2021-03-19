package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var Repo string
var RepoType string
var Owner string
var Token string

var rootCmd = &cobra.Command{
	Use:   "semver",
	Short: "heycar tool for GitHub SemVer releases and Slack release notes",
	Long:  ``,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
