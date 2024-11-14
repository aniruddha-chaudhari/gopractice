package config

import (
	"flag"
	"log"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
)

//this here struct with strucr tags is used to define the structure of the config file

type Config struct {
	Env string `yaml:"env" env:"ENV" env-required:"true" env-default:"production"`
	Storagepath string `yaml:"storagepath" env-required:"true"`
	HTTPServer `yaml:"http_server"`
}

type HTTPServer struct {
    Addr string `yaml:"address" env-required:"true"`
}

func MustLoad() *Config {
	var configPath string

	configPath = os.Getenv("CONFIG_PATH")

	if configPath == "" {
		flags := flag.String("config", "", "path to config file")
		flag.Parse()

		configPath = *flags

		if configPath == "" {
			log.Fatal("config path is required")
		}
		
	}
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		log.Fatalf("config file not found: %s", configPath)
	}

	var cfg Config 

	err := cleanenv.ReadConfig(configPath, &cfg)

	if err != nil {
		log.Fatalf("failed to read config: %v", err.Error())
	}

	return &cfg
}