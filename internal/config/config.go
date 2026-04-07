package config

type Config struct {
	HTTPPort string
}

func MustLoad() Config {
	return Config{
		HTTPPort: "8080",
	}
}