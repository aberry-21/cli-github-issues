package cmd

import (
	"cli-github-issues/internal/config"
	"cli-github-issues/internal/github"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"log"
	"net/http"
)

var (
	rootCmd = &cobra.Command{
		Use:   "cli-github-issues",
		Short: "A command line utility for working with github issues",
	}
	cfgFile string
	cfg     *config.Config
	client  *github.Client
)

func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	cobra.OnInitialize(func() {
		// load config
		cfg = config.MustLoad(cfgFile)

		// init GitHub client
		var err error
		if client, err = github.NewClient(http.DefaultClient, cfg.Token); err != nil {
			log.Fatal(err)
		}
	})

	// init global command line flags
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/cli-github-issues.cobra.yaml)")
	rootCmd.PersistentFlags().String("editor", "code", "issue editor")
	rootCmd.PersistentFlags().String("owner", "", "owner of repository")
	rootCmd.PersistentFlags().String("repo", "", "repository")
	rootCmd.PersistentFlags().String("token", "", "GitHub token")

	// bind cli flags with viper
	viper.BindPFlag("editor", rootCmd.PersistentFlags().Lookup("editor"))
	viper.BindPFlag("github.owner", rootCmd.PersistentFlags().Lookup("owner"))
	viper.BindPFlag("github.repo", rootCmd.PersistentFlags().Lookup("repo"))
	viper.BindPFlag("github.token", rootCmd.PersistentFlags().Lookup("token"))
}

func flagMustExist[T any](v T, err error) T {
	if err != nil {
		log.Fatal(err)
	}
	return v
}
