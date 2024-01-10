package cmd

import (
	"cli-github-issues/internal/editor"
	"cli-github-issues/internal/github"
	"fmt"
	"github.com/spf13/cobra"
	"log"
	"net/http"
)

var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update an issue on the specified repository",
	Run: func(cmd *cobra.Command, args []string) {
		// get issue request params
		number := flagMustExist(cmd.Flags().GetInt("number"))
		title := flagMustExist(cmd.Flags().GetString("title"))
		bodyData, err := editor.SaveInputDataToFile(cfg.Editor)
		if err != nil {
			return
		}

		// update issue
		editedIssue := &github.IssueRequest{
			Title: &title,
			Body:  github.String(string(bodyData)),
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
	rootCmd.AddCommand(updateCmd)

	// set required flag
	updateCmd.Flags().Int("number", 0, "issue number")
	updateCmd.Flags().String("title", "", "Issue title")

	updateCmd.MarkFlagRequired("title")
	updateCmd.MarkFlagRequired("number")
}
