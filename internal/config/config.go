package config

type ServerConfig struct {
	Address string
}

type Config struct {
	ServerConfig
}

func LoadConfig() Config {
	return Config{
		ServerConfig: ServerConfig{
			Address: "127.0.0.1:8080",
		},
	}
}
