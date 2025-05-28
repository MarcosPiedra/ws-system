package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type (
	RedisConfig struct {
		Addr    string `mapstructure:"ADDR"`
		Channel string `mapstructure:"CHANNEL"`
	}

	WebConfig struct {
		Port   string `mapstructure:"PORT"`
		Secret string `mapstructure:"SECRET"`
		User   string `mapstructure:"USER"`
		Pass   string `mapstructure:"PASS"`
	}

	AppConfig struct {
		Environment string      `mapstructure:"ENVIRONMENT"`
		LogLevel    string      `mapstructure:"LOGLEVEL"`
		Web         WebConfig   `mapstructure:"WEB"`
		Redis       RedisConfig `mapstructure:"REDIS"`
	}
)

func Setup() (cfg AppConfig, err error) {

	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.AddConfigPath("/ws")

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	viper.SetDefault("LogLevel", "Debug")

	err = viper.Unmarshal(&cfg)

	return
}

func (c WebConfig) Address() string {
	return fmt.Sprintf("%s", c.Port)
}
