package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	Currency  *CurrencyConfig  `mapstructure:"currency"`
	Root      *RootConfig      `mapstructure:"root"`
	Portfolio *PortfolioConfig `mapstructure:"portfolio"`
}

type CurrencyConfig struct {
	InitUnit string
}

type RootConfig struct {
	User string
}

type PortfolioConfig struct {
	User string
}

func setDefault() {
	viper.SetDefault("currency.initUnit", "united-states-dollar")
	viper.SetDefault("root.user", "default")
	viper.SetDefault("portfolio.user", "default")
}

func NewConfig() (*Config, error) {
	setDefault()

	var c Config
	if err := viper.Unmarshal(&c); err != nil {
		return nil, err
	}
	return &c, nil
}
