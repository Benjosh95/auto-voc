package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	ServerConfig
	DBConfig
}
type ServerConfig struct {
	Address string
}
type DBConfig struct {
	ConnectionString string
}

// TODO: maybe use library for env-variables
func LoadConfig() Config {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found")
	}

	return Config{
		ServerConfig: ServerConfig{
			Address: getEnv("SERVER_ADDRESS", "localhost:8080"),
		},
		DBConfig: DBConfig{
			ConnectionString: getEnv("DATABASE_URL", ""),
		},
	}
}

func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}
