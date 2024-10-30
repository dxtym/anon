package utils

import "github.com/spf13/viper"

type Config struct {
	Address string `mapstructure:"ADDRESS"`
}

func LoadConfig(path string) (Config, error) {
	var cfg Config
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")
	if err := viper.ReadInConfig(); err != nil {
		return cfg, err
	}
	err := viper.Unmarshal(&cfg)
	return cfg, err
}