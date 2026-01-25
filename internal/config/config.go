package config

type Config struct {
	AppName string
	Port    int
	DBHost  string
	DBPort  int
	DBName  string
}

func Load() *Config {
	return &Config{
		AppName: "TaskeFlow API",
		Port:    8080,
		DBHost:  "localhost",
		DBPort:  9000,
		DBName:  "taskeflow_db",
	}
}
