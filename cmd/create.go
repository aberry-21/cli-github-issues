package cmd

import (
	"cli-github-issues/internal/editor"
	"cli-github-issues/internal/github"
	"fmt"
	"log"
	"net/http"

	"github.com/spf13/cobra"
)

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new issue on the specified repository",
	Run: func(cmd *cobra.Command, args []string) {
		// save title of issue
		title := flagMustExist(cmd.Flags().GetString("title"))

		// save body of issue
		bodyData, err := editor.SaveInputDataToFile(cfg.Editor)
		if err != nil {
			return
		}

		// create issue
		req := &github.IssueRequest{
			Title: &title,
			Body:  github.String(string(bodyData)),
		}
		issue, resp, err := client.Issues.Create(cfg.Owner, cfg.Repo, req)
		if err != nil {
			log.Fatal(err)
		}

		// check and print result
		if resp.StatusCode == http.StatusCreated {
			fmt.Printf("#%-5d %9.9s %.55s %q\n", *issue.Number, *issue.User.Login, *issue.Title, *issue.Body)
		} else {
			log.Fatalf("Invalid status code: %d", resp.StatusCode)
		}
	},
}

func init() {
	rootCmd.AddCommand(createCmd)

	// set required flag
	createCmd.Flags().String("title", "", "Issue title")
	createCmd.MarkFlagRequired("title")
}
