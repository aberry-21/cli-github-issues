package config

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"log"
	"os"
)

type Config struct {
	Editor string `mapstructure:"editor"`
	Github `mapstructure:"github"`
}

type Github struct {
	Owner string `mapstructure:"owner"`
	Repo  string `mapstructure:"repo"`
	Token string `mapstructure:"token"`
}

func MustLoad(cfgFile string) *Config {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		// Search config in home directory with name "cli-github-issues.cobra" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigType("yaml")
		viper.SetConfigName("cli-github-issues.cobra")
	}

	//viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("cannot read config: %s", err)
	}

	var cfg Config
	err := viper.Unmarshal(&cfg)
	if err != nil {
		log.Fatalf("unable to decode into struct, %v", err)
	}

	return &cfg
}
