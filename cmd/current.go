package cmd

import (
	"fmt"
	"github.com/heycar-uk/semver/internal/utils"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(currentCmd)
	currentCmd.Flags().StringVarP(&URL, "URL", "u", "", "/status/ok URL (required)")
	currentCmd.MarkFlagRequired("URL")
}

var currentCmd = &cobra.Command{
	Use:   "current",
	Short: "Get the current deployed release",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Print(utils.GetCurrentDeployedRelease(URL))
	},
}
