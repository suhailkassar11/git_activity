package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/suhailkassar11/github_activity/internal/api"
	// "github.com/suhailkassar11/github_activity/internal/models"
)

// RootCmd represents the base command when called without any subcommands.
var RootCmd = &cobra.Command{
	Use:   "github-activity",
	Short: "Fetch GitHub user activity",
	Long: `github-activity is a CLI tool to fetch and display recent activity 
for a specified GitHub user, such as commits, issues, and pull requests.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		username, err := cmd.Flags().GetString("username")
		if err != nil {
			return fmt.Errorf("error reading username flag: %w", err)
		}

		events, err := api.FetchUser(username)
		if err != nil {
			return fmt.Errorf("failed to fetch user events: %w", err)
		}

		// Process and display events (example output)
		if len(events) == 0 {
			fmt.Printf("No recent activity found for user %s\n", username)
			return nil
		}

		fmt.Printf("Recent activity for %s:\n", username)
		for i, event := range events {
			fmt.Printf("%d. Type: %s, Repo: %s, Created: %s\n",
				i+1, event.Type, event.Repo.Name, event.CreatedAt)
		}

		return nil
	},
}

// Execute runs the root command and handles errors.
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}

// init configures the root command’s flags and settings.
func init() {
	RootCmd.Flags().StringP("username", "u", "", "GitHub username (required)")
	RootCmd.MarkFlagRequired("username")
}
