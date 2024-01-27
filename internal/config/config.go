package config

import (
	"sync"

	"github.com/gotrika/gotrika_backend/pkg/logger"
	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Debug      bool   `env:"DEBUG" env-default:"false"`
	SecretKey  string `env:"SECRET_KEY" env-required:"true"`
	AccessTTL  int    `env:"ACCESS_TTL" env-default:"3600"`     // in seconds
	RefreshTTL int    `env:"REFRESH_TTL" env-default:"2592000"` // in seconds
	HTTPConfig struct {
		Port string `env:"PORT" env-default:"8000"`
	}
	CabbageConfig struct {
		BrokerURI        string `env:"CABBAGE_BROKER_URI" env-required:"true"`
		WorkerCount      int    `env:"CABBAGE_WORKER_COUNT" env-default:"10"`
		EventQueueName   string `env:"CABBAGE_EVENT_QUEUENAME" env-default:"gotrika_events"`
		SessionQueueName string `env:"CABBAGE_SESSION_QUEUENAME" env-default:"gotrika_sessions"`
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
		instance = &Config{}
		helpText := "~ GOtrika backend ~"
		help, _ := cleanenv.GetDescription(instance, &helpText)
		logger.Info(help)
		err := cleanenv.ReadConfig(".env", instance)
		if err != nil {
			if err := cleanenv.ReadEnv(instance); err != nil {
				logger.Fatal(err)
			}
		}

	})
	return instance
}
