package config

import "github.com/caarlos0/env/v6"

var (
	config *Config
)

type Config struct {
	GinPort       string `env:"GIN_PORT" envDefault:"8080"`
	MySQLHost     string `env:"MYSQL_HOST"`
	MySQLPort     int    `env:"MYSQL_PORT" envDefault:"3306"`
	MySQLUser     string `env:"MYSQL_USER"`
	MySQLPassword string `env:"MYSQL_PASSWORD"`
	MySQLDatabase string `env:"MYSQL_DATABASE"`
}

func Init() {
	config = &Config{}
	if err := env.Parse(config); err != nil {
		panic(err)
	}
}

func GetConfig() *Config {
	return config
}
