package cmd

import (
	"fmt"

	"github.com/spf13/viper"
)

type Redis struct {
	Host string `mapstructure:"host"`
	Port uint16 `mapstructure:"port"`
}

type Config struct {
	Redis Redis `mapstructure:"redis"`
}

func InitConfig() error {
	viper.SetConfigFile(configFile)

	viper.AutomaticEnv()

	viper.SetDefault(
		"redis",
		Redis{
			Host: "0.0.0.0",
			Port: 6379,
		},
	)

	if err := viper.ReadInConfig(); err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}

func ParseConfig(c *Config) error {
	if err := viper.Unmarshal(c); err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}
