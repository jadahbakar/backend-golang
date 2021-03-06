package config

import (
	"github.com/spf13/viper"
)

//  exported
type Config struct {
	ApiURLGroup       string `mapstructure:"API_URL_GROUP"`
	ApiURLVersion     string `mapstructure:"API_URL_VERSION"`
	LogFolder         string `mapstructure:"LOG_FOLDER"`
	HttpPort          int    `mapstructure:"HTTP_PORT"`
	ServerReadTimeOut int    `mapstructure:"SERVER_READ_TIMEOUT"`
	Prefork           bool   `mapstructure:"PREFORK"`
	CaseSensitive     bool   `mapstructure:"CASE_SENSITIVE"`
	DBURL             string `mapstructure:"DATABASE_URL"`
}

// LoadConfig reads configuration from file or environment variables.
func LoadConfig() (config Config, err error) {
	// viper.SetConfigFile(".env")
	viper.AddConfigPath("./")
	viper.SetConfigName("app")
	viper.SetConfigType("env")
	viper.AutomaticEnv()
	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}
