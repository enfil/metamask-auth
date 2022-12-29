package config

import (
	"github.com/spf13/viper"
)

type Settings struct {
	JWT struct {
		TTL    int    `mapstructure:"JWT_TTL"`
		Secret string `mapstructure:"JWT_SECRET"`
		Issuer string `mapstructure:"JWT_ISSUER"`
	}
}

func LoadAndStoreConfig(configPath string) Settings {
	viper.SetConfigFile(configPath)
	err := viper.ReadInConfig()
	if err != nil {
		return Settings{}
	}

	var cfg Settings

	err = viper.Unmarshal(&cfg)
	err = viper.Unmarshal(&cfg.JWT)

	return cfg
}
