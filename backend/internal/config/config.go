package config

import (
	"log"
	"os"

	"github.com/spf13/viper"
)

type Config struct {
	Server   *Server  `mapstructure:"server"`
	Database *Storage `mapstructure:"db"`
}

type Server struct {
	Host string `mapstructure:"host"`
	Port string `mapstructure:"port"`
}

type Storage struct {
	Host     string `mapstructure:"host"`
	Port     string `mapstructure:"port"`
	Sslmode  string `mapstructure:"sslmode"`
	DbName   string
	Username string
	Password string
}

func MustLoadConfig() *Config {
	viper.AddConfigPath("./internal/config")
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatal(err)
	}

	cfg := &Config{}

	if err := viper.Unmarshal(cfg); err != nil {
		log.Fatal(err)
	}

	cfg.Database.Password = os.Getenv("DB_PASSWORD")
	cfg.Database.Username = os.Getenv("POSTGRES_USER")
	cfg.Database.DbName = os.Getenv("POSTGRES_DB")

	return cfg
}
