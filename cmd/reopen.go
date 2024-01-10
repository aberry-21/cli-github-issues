package cmd

import (
	"cli-github-issues/internal/github"
	"fmt"
	"github.com/spf13/cobra"
	"log"
	"net/http"
)

var reopenCmd = &cobra.Command{
	Use:   "reopen",
	Short: "Reopen an issue on the specified repository.",
	Run: func(cmd *cobra.Command, args []string) {
		// get issue number from cli
		number := flagMustExist(cmd.Flags().GetInt("number"))

		// reopen issue
		editedIssue := &github.IssueRequest{
			State: github.String("open"),
		}
		issue, resp, err := client.Issues.Update(cfg.Owner, cfg.Repo, number, editedIssue)
		if err != nil {
			log.Fatal(err)
		}

		// check and print result
		if resp.StatusCode == http.StatusOK {
			fmt.Printf("#%-5d %9.9s %.55s %q\n", *issue.Number, *issue.User.Login, *issue.Title, *issue.Body)
		} else {
			log.Fatalf("Invalid status code: %d", resp.StatusCode)
		}
	},
}

func init() {
	rootCmd.AddCommand(reopenCmd)

	// set required flag
	reopenCmd.Flags().Int("number", 0, "issue number")
	reopenCmd.MarkFlagRequired("number")
}
