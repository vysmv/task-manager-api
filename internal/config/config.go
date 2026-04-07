package config

type Config struct {
	HTTPPort string
	DBHost   string
	DBPort   string
	DBUser   string
	DBPass   string
	DBName   string
	DBSSLMode string
}

func MustLoad() Config {
	return Config{
		HTTPPort:  "8080",
		DBHost:    "localhost",
		DBPort:    "5432",
		DBUser:    "task_manager",
		DBPass:    "task_manager",
		DBName:    "task_manager",
		DBSSLMode: "disable",
	}
}