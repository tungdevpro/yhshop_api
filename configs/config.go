package configs

import (
	"coffee_api/helpers"

	"github.com/caarlos0/env/v9"
	"github.com/joho/godotenv"
)

type Configuration struct {
	Port            string `env:"PORT" envDefault:"3001"`
	HashSalt        string `env:"HASH_SALT,required"`
	DBConnectionURL string `env:"DB_CONNECTION_URL,required"`
	SecretKey       string `env:"SECRET_KEY,required"`
}

func NewConfiguration() *Configuration {
	err := godotenv.Load()
	if err != nil {
		helpers.Fatal("No .env file could be found")
	}

	cfg := Configuration{}
	err = env.Parse(&cfg)
	if err != nil {
		helpers.Fatal(err)
	}

	return &cfg
}
