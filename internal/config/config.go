package config

import (
	"log"
	"sync"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Debug      bool   `env:"DEBUG" env-default:"false"`
	SecretKey  string `env:"SECRET_KEY" env-required:"true"`
	HTTPConfig struct {
		Port string `env:"PORT" env-default:"8000"`
	}
	AMQPConfig struct {
		URI string `env:"AMQP_URI" env-required:"true"`
	}
	MongoConfig struct {
		HOST     string `env:"MONGO_HOST" env-required:"true"`
		PORT     string `env:"MONGO_PORT" env-required:"true"`
		User     string `env:"MONGO_USER"`
		Password string `env:"MONGO_PASSWORD"`
		DBName   string `env:"MONGO_DB_NAME" env-required:"true"`
	}
	LogLevel string `env:"LOG_LEVEL" env-default:"trace"`
}

var instance *Config
var once sync.Once

func GetConfig() *Config {
	once.Do(func() {
		log.Print("gather config")

		instance = &Config{}

		if err := cleanenv.ReadEnv(instance); err != nil {
			helpText := "~ GOtrika backend ~"
			help, _ := cleanenv.GetDescription(instance, &helpText)
			log.Print(help)
			log.Fatal(err)
		}
	})
	return instance
}
