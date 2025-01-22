package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	Database struct {
		Host     string `mapstructure:"host"`
		Port     int    `mapstructure:"port"`
		User     string `mapstructure:"user"`
		Password string `mapstructure:"password"`
		DBName   string `mapstructure:"db_name"`
		SSL      bool   `mapstructure:"ssl"`
	} `mapstructure:"database"`

	AWS struct {
		AccessKeyID     string `mapstructure:"access_key_id"`
		SecretAccessKey string `mapstructure:"secret_access_key"`
		Region          string `mapstructure:"region"`
	} `mapstructure:"aws"`

	Service struct {
		Host string `mapstructure:"host"`
		Port int    `mapstructure:"port"`
	} `mapstructure:"service"`

	Logging struct {
		Level  string `mapstructure:"level"`
		Format string `mapstructure:"format"`
	} `mapstructure:"logging"`
}

var AppConfig Config

func LoadConfig(env string) {
	viper.AddConfigPath("./config/env")
	viper.SetConfigType("yaml")

	configFile := "./config/env/" + env + "/hotel-booking.yaml"
	viper.SetConfigFile(configFile)

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}

	if err := viper.Unmarshal(&AppConfig); err != nil {
		log.Fatalf("Unable to decode into struct, %v", err)
	}
}
