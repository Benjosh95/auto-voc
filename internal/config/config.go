package config

import "os"

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
	return Config{
		ServerConfig: ServerConfig{
			Address: getEnv("SERVER_ADDRESS", "127.0.0.1:8080"),
		},
		DBConfig: DBConfig{
			ConnectionString: getEnv("DB_CONNECTION_STRING", "postgres://vocmaster:vocmasterpassword@voc-db.cbeug4q2kw1j.eu-central-1.rds.amazonaws.com:5432/voc-db?sslmode=require"),
		},
	}
}

func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}
