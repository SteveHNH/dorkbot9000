package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type BotConfig struct {
	Username string `mapstructure:"username"`
	OAuth    string `mapstructure:"oauth"`
	Channel  string `mapstructure:"channel"`
}

// Load configuration for service
func Load() *BotConfig {

	// Load configuration from file
	viper.SetConfigName("botconfig.yaml")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.AddConfigPath("$HOME/.config/dorkbot9000")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}

	return &BotConfig{
		Username: viper.GetString("username"),
		OAuth:    viper.GetString("password"),
		Channel:  viper.GetString("channel"),
	}
}
