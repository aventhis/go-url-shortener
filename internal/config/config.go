package config

import (
	"log"
	"os"
	"time"
)

type Config struct {
	Env         string     `yaml:"env" env-default:"local" env-required:"true"`
	StoragePath string     `yaml:"storage_path" env-required:"true"`
	HttpServer  HttpServer `yaml:"http_server"`
}

type HttpServer struct {
	Address     string        `yaml:"address" env-default:"localhost:8080"`
	Timeout     time.Duration `yaml:"timeout" env-default:"4s"`
	IdleTimeout time.Duration `yaml:"idle_timeout" env-default:"60s"`
}

func MustLoad() *Config {
	configPath := os.Getenv("CONFIG_PATH")
	if configPath == "" {
		log.Fatal("CONFIG_PATH environment variable not set")
	}

	// check if file exist
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		log.Fatalf("CONFIG_PATH does not exist %s", configPath)
	}

	var cfg Config

}
