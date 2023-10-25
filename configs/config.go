package configs

import (
	"coffee_api/helpers"

	"github.com/caarlos0/env/v9"
	"github.com/joho/godotenv"
)

type Configuration struct {
	Port                string `env:"PORT" envDefault:"3001"`
	HashSalt            string `env:"HASH_SALT,required"`
	DBConnectionURL     string `env:"DB_CONNECTION_URL,required"`
	SecretKey           string `env:"SECRET_KEY,required"`
	S3BucketName        string `env:"S3_BUCKET_NAME,required"`
	S3Region            string `env:"S3_REGION,required"`
	S3ApiKey            string `env:"S3_API_KEY,required"`
	S3SecretKey         string `env:"s3_SECRET_KEY,required"`
	S3Domain            string `env:"S3_DOMAIN,required"`
	EmailSenderName     string `env:"EMAIL_SENDER_NAME,required"`
	EmailSenderAddress  string `env:"EMAIL_SENDER_ADDRESS,required"`
	EmailSenderPassword string `env:"EMAIL_SENDER_PASSWORD,required"`
	ApplicationName     string `env:"APPLICATION_NAME,required"`
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
