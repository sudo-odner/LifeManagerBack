package config

import (
	"errors"
	"github.com/ilyakaznacheev/cleanenv"
	"github.com/joho/godotenv"
	"log"
	"os"
)

type Config struct {
	Logger     string     `yaml:"logger" env-default:"info"` // Прочитаь про struct tag
	HTTPServer HTTPServer `yaml:"http_server"`
	DB         DB         `yaml:"db"`
}

type HTTPServer struct {
	Host string `yaml:"host" env-default:"localhost"`
	Port string `yaml:"port" env-default:"8080"`
}

type DB struct {
	RebuildMode bool    `yaml:"rebuild_mode"`
	MongoDB     MongoDB `yaml:"mongo_db"`
}

type MongoDB struct {
	Uri string `yaml:"uri"`
}

func New() (*Config, error) {
	// Провекрка существования .env
	if err := godotenv.Load(); err != nil {
		log.Print(".env file found")
		return nil, err
	}

	configPath, exists := os.LookupEnv("CONFIG_PATH")
	if !exists {
		log.Fatal("CONFIG_PATH is not found")
		return nil, errors.New("CONFIG_PATH is not found")
	}

	// Провека на существования файла
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		log.Fatalf("config file does not exist: %s", configPath)
		return nil, err
	}

	var cfg Config
	if err := cleanenv.ReadConfig(configPath, &cfg); err != nil {
		log.Fatalf("cannot read config: %s", err)
		return nil, err
	}

	return &cfg, nil
}
