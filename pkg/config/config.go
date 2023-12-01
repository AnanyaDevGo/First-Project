package config

import (
	"github.com/go-playground/validator/v10"
	"github.com/spf13/viper"
)

type Config struct {
	DBHost     string `mapstructure:"DB_HOST"`
	DBName     string `mapstructure:"DB_NAME"`
	DBUser     string `mapstructure:"DB_USER"`
	DBPort     string `mapstructure:"DB_PORT"`
	DBPassword string `mapstructure:"DB_PASSWORD"`
	AUTHTOKEN  string `mapstructure:"DB_AUTHTOKEN"`
	ACCOUNTSID string `mapstructure:"DB_ACCOUNTSID"`
	SERVICESID string `mapstructure:"DB_SERVICESID"`

	Access_key_ID     string `mapstructure:"Access_key_ID"`
	Secret_access_key string `mapstructure:"Secret_access_key"`
	AWSRegion         string `mapstructure:"AWSRegion"`

	KEY_ID_FOR_PAY     string `mapstructure:"KEY_ID_FOR_PAY"`
	SECRET_KEY_FOR_PAY string `mapstructure:"SECRET_KEY_FOR_PAY"`
}

var envs = []string{
	"DB_HOST", "DB_NAME", "DB_USER", "DB_PORT", "DB_PASSWORD", "DB_AUTHTOKEN", "DB_ACCOUNTSID", "DB_SERVICESID", "Access_key_ID", "Secret_access_key", "AWSRegion", "KEY_ID_FOR_PAY", "SECRET_KEY_FOR_PAY",
}

func LoadConfig() (Config, error) {
	var config Config

	viper.AddConfigPath("./")
	viper.SetConfigFile(".env")
	viper.ReadInConfig()

	for _, env := range envs {
		if err := viper.BindEnv(env); err != nil {
			return config, err
		}
	}

	if err := viper.Unmarshal(&config); err != nil {
		return config, err
	}

	if err := validator.New().Struct(&config); err != nil {
		return config, err
	}
	return config, nil
}
