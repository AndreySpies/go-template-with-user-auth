package config

import (
	"strings"

	"github.com/spf13/viper"
)

type Config struct {
	Version string `mapstructure:"-"`

	ApiKey            string `mapstructure:"api_key"`
	Port              int    `mapstructure:"port"`
	HTTPBodySizeLimit string `mapstructure:"http_body_size_limit"`

	DBUser     string `mapstructure:"mysql_db_user"`
	DBPassword string `mapstructure:"mysql_db_password"`
	DBHost     string `mapstructure:"mysql_db_host"`
	DBPort     int    `mapstructure:"mysql_db_port"`
	DBName     string `mapstructure:"mysql_db_name"`

	APIs map[string]map[string]string `mapstructure:"apis"`
}

func Read() (*Config, error) {
	v := viper.New()

	v.SetDefault("api_key", "")
	v.SetDefault("port", 5050)
	v.SetDefault("http_body_size_limit", "4MB")

	v.SetDefault("mysql_db_user", "")
	v.SetDefault("mysql_db_password", "")
	v.SetDefault("mysql_db_host", "")
	v.SetDefault("mysql_db_port", 3306)
	v.SetDefault("mysql_db_name", "")

	v.AddConfigPath(".")
	v.AddConfigPath("..")
	v.AddConfigPath("../..")
	v.SetConfigName("config")
	v.SetEnvKeyReplacer(strings.NewReplacer("-", "_"))
	v.SetEnvPrefix("app")
	v.AutomaticEnv()
	v.SetTypeByDefaultValue(true)

	v.ReadInConfig()

	var config Config
	err := v.Unmarshal(&config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}
