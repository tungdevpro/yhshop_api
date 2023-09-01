package configs

import "github.com/joho/godotenv"

type Configuration struct {
	HashSalt int `"env:"SALT" envDefault:"3001"`
}


func NewConfiguration() *Configuration {
	err := godotenv.Load()
	if err != nil {

	}

	return &Configuration{
	}
}
