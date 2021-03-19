package cmd

import (
	"context"
	"fmt"
	gh "github.com/google/go-github/v33/github"
	"github.com/heycar-uk/semver/internal/github"
	"github.com/spf13/cobra"
	"golang.org/x/oauth2"
)

var CurrentDeployed string

func init() {
	rootCmd.AddCommand(backlogCmd)
	backlogCmd.PersistentFlags().StringVarP(&Repo, "repo", "r", "", "Repo name (required)")
	backlogCmd.PersistentFlags().StringVarP(&Owner, "owner", "o", "heycar-uk", "GitHub Owner")
	backlogCmd.PersistentFlags().StringVarP(&Token, "token", "t", "", "GitHub Token (required)")
	backlogCmd.Flags().StringVarP(&CurrentDeployed, "deployed", "d", "", "Current deployed version (required)")
	backlogCmd.Flags().StringVarP(&RepoType, "code", "c", "", "Repo type [frontend/backend/devops]")
	backlogCmd.MarkFlagRequired("deployed")
	backlogCmd.MarkFlagRequired("code")
	backlogCmd.MarkPersistentFlagRequired("token")
	backlogCmd.MarkPersistentFlagRequired("repo")
}

var backlogCmd = &cobra.Command{
	Use:   "backlog",
	Short: "Get all the backlogged release notes",
	Run: func(cmd *cobra.Command, args []string) {
		ctx := context.Background()
		ts := oauth2.StaticTokenSource(
			&oauth2.Token{AccessToken: Token},
		)
		tc := oauth2.NewClient(ctx, ts)
		client := gh.NewClient(tc)

		fmt.Print(github.GetReleaseNotesBacklog(ctx, client, Owner, Repo, CurrentDeployed, RepoType))
	},
}
