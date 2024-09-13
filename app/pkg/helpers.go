package pkg

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string
}

func (c *Config) GetConnectionString() string {
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", c.DBHost, c.DBPort, c.DBUser, c.DBPassword, c.DBName)
}

func LoadConfig() (*Config, error) {
	// Load .env file if it exists (for local development)
	err := godotenv.Load()
	if err != nil {
		log.Fatal("No .env file found")
	}

	config := &Config{
		DBHost:     os.Getenv("DB_HOST"),
		DBPort:     os.Getenv("DB_PORT"),
		DBUser:     os.Getenv("DB_USER"),
		DBPassword: os.Getenv("DB_PASSWORD"),
		DBName:     os.Getenv("DB_NAME"),
	}

	// Validate required configuration
	if config.DBHost == "" || config.DBPort == "" || config.DBUser == "" || config.DBPassword == "" || config.DBName == "" {
		return nil, fmt.Errorf("missing required database configuration")
	}

	return config, nil
}
