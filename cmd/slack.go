package cmd

import (
	"github.com/heycar-uk/semver/internal/slack"
	"github.com/spf13/cobra"
)

var Webhook string
var ReleaseNotes string

func init() {
	rootCmd.AddCommand(slackCmd)
	slackCmd.Flags().StringVarP(&ReleaseNotes, "notes", "n", "", "Release notes (required)")
	slackCmd.Flags().StringVarP(&Webhook, "webhook", "w", "", "Webhook URL (required)")
	slackCmd.MarkFlagRequired("notes")
	slackCmd.MarkFlagRequired("webhook")
}

var slackCmd = &cobra.Command{
	Use:   "slack",
	Short: "Create a Slack release message",
	Run: func(cmd *cobra.Command, args []string) {
		slack.SendMessage(Webhook, ReleaseNotes)
	},
}
