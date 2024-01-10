package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"log"
	"net/http"
)

var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Get an issue from the specified repository",
	Run: func(cmd *cobra.Command, args []string) {
		// get issue number from cli
		number := flagMustExist(cmd.Flags().GetInt("number"))

		// get issue
		issue, resp, err := client.Issues.Get(cfg.Owner, cfg.Repo, number)
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
	rootCmd.AddCommand(getCmd)

	// set required flag
	getCmd.Flags().Int("number", 0, "issue number")
	getCmd.MarkFlagRequired("number")
}
