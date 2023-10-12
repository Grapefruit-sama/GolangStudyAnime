package config

// todo:сделать конфиг под бд
type Config struct {
	ServerPort string
}

func NewConfig(port string) *Config {
	if port == "" {
		port = "3001"
	}

	return &Config{
		ServerPort: port,
	}
}
