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

// MustLoad loads config file and returns config struct.
func MustLoad(cfgFile string) *Config {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Use the user's home directory to search for the config file
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		// Search config with name "cli-github-issues.cobra.yaml" in home directory
		viper.AddConfigPath(home)
		viper.SetConfigType("yaml")
		viper.SetConfigName("cli-github-issues.cobra")
	}

	// Read the config file
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("cannot read config: %s", err)
	}

	// Unmarshal config
	var cfg Config
	err := viper.Unmarshal(&cfg)
	if err != nil {
		log.Fatalf("unable to decode into struct, %v", err)
	}

	return &cfg
}
