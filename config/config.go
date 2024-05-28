package config

import (
	"log"

	"github.com/ilyakaznacheev/cleanenv"
)

type ConfigDatabase struct {
	Port     string `env:"DB_PORT"`
	Host     string `env:"DB_HOST"`
	Name     string `env:"DB_NAME"`
	User     string `env:"DB_USER"`
	Password string `env:"DB_PASSWORD"`
}

func LoadConfig() (ConfigDatabase, error) {
	var cfg ConfigDatabase
	err := cleanenv.ReadEnv(&cfg)
	if err != nil {
		log.Println("Error reading environment variables")
	} else {
		log.Printf("Loaded config: %#v\n", cfg)
	}

	// Дополнительные проверки и вывод логов для диагностики
	if cfg.Port == "" || cfg.Host == "" || cfg.Name == "" || cfg.User == "" || cfg.Password == "" {
		log.Println("One or more required environment variables are empty")
	}

	return cfg, err
}

//debug

// func LoadConfig() (ConfigDatabase, error) {
// 	var cfg ConfigDatabase

// 	// Load .env file if it exists
// 	if err := godotenv.Load(); err != nil {
// 		log.Println("No .env file found")
// 	}

// 	err := cleanenv.ReadEnv(&cfg)
// 	if err != nil {
// 		log.Println("Error reading environment variables")
// 	} else {
// 		log.Printf("Loaded config: %#v\n", cfg)
// 	}
// 	return cfg, err
// }
//
