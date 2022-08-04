package common

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	DBUser            string `mapstructure:"DB_USER"`
	DBPassword        string `mapstructure:"DB_PASSWORD"`
	DBHost            string `mapstructure:"DB_HOST"`
	DBName            string `mapstructure:"DB_NAME"`
	DBPort            string `mapstructure:"DB_PORT"`
	GRPCServerAddress string `mapstructure:"GRPC_SERVER_ADDRESS"`
}

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
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

func NewConfig() *Config {
	config, err := LoadConfig(".")
	if err != nil {
		log.Fatal("Cannot load config:", err)
	}

	return &config
}
