package cmd

import (
	"context"
	"fmt"

	gh "github.com/google/go-github/v33/github"
	"github.com/heycar-uk/semver/internal/github"
	"github.com/spf13/cobra"
	"golang.org/x/oauth2"
)

var SHA1 string
var URL string

func init() {
	rootCmd.AddCommand(prCmd)
	prCmd.PersistentFlags().StringVarP(&Repo, "repo", "r", "", "Repo name (required)")
	prCmd.PersistentFlags().StringVarP(&Owner, "owner", "o", "heycar-uk", "GitHub Owner")
	prCmd.PersistentFlags().StringVarP(&Token, "token", "t", "", "GitHub Token (required)")
	prCmd.PersistentFlags().StringVarP(&SHA1, "sha1", "s", "", "Git SHA1 ref (required)")
	prCmd.MarkPersistentFlagRequired("token")
	prCmd.MarkPersistentFlagRequired("repo")
	prCmd.MarkPersistentFlagRequired("sha1")
	prCmd.AddCommand(semverCmd)
	prCmd.AddCommand(nameCmd)
	prCmd.AddCommand(notesCmd)
	prCmd.AddCommand(githubReleasedCmd)
}

var prCmd = &cobra.Command{
	Use:   "release",
	Short: "Get information on the new release",
}

var semverCmd = &cobra.Command{
	Use:   "number",
	Short: "Get the new release semver number",
	Run: func(cmd *cobra.Command, args []string) {
		ctx := context.Background()
		ts := oauth2.StaticTokenSource(
			&oauth2.Token{AccessToken: Token},
		)
		tc := oauth2.NewClient(ctx, ts)
		client := gh.NewClient(tc)

		pr := github.GetRelease(ctx, client, Owner, Repo, SHA1)
		fmt.Print(pr.ReleaseNumber)
	},
}

var nameCmd = &cobra.Command{
	Use:   "name",
	Short: "Get the release name",
	Run: func(cmd *cobra.Command, args []string) {
		ctx := context.Background()
		ts := oauth2.StaticTokenSource(
			&oauth2.Token{AccessToken: Token},
		)
		tc := oauth2.NewClient(ctx, ts)
		client := gh.NewClient(tc)

		pr := github.GetRelease(ctx, client, Owner, Repo, SHA1)
		fmt.Print(pr.Name)
	},
}

var notesCmd = &cobra.Command{
	Use:   "notes",
	Short: "Get the release notes",
	Run: func(cmd *cobra.Command, args []string) {
		ctx := context.Background()
		ts := oauth2.StaticTokenSource(
			&oauth2.Token{AccessToken: Token},
		)
		tc := oauth2.NewClient(ctx, ts)
		client := gh.NewClient(tc)

		pr := github.GetRelease(ctx, client, Owner, Repo, SHA1)
		fmt.Print("Author: ", pr.Author, pr.ReleaseNotes)
	},
}

var githubReleasedCmd = &cobra.Command{
	Use:   "commit",
	Short: "Get the GitHub release for a commit",
	Run: func(cmd *cobra.Command, args []string) {
		ctx := context.Background()
		ts := oauth2.StaticTokenSource(
			&oauth2.Token{AccessToken: Token},
		)
		tc := oauth2.NewClient(ctx, ts)
		client := gh.NewClient(tc)

		fmt.Print(github.GetReleaseByCommit(ctx, client, Owner, Repo, SHA1))
	},
}
