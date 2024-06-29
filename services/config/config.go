package config

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
)
var config *Config = nil
type Config struct {
	Port int
	APIInitialPath string

	LogFilePath string
	SwaggerFilePath string
	SwaggerHTMLFilePath string

	FrontendURL string
}

func init() {
	envFile, exists := os.LookupEnv("ENV_FILE")

	if exists {
		godotenv.Load(envFile)
	}
}

func Get() *Config{
	if config == nil {
		config = &Config{
			Port:           getEnvAsInt("PORT", 2020),
			APIInitialPath: getEnvOrDefault("APIInitialPath", "/api/v1"),

			LogFilePath: getEnvOrDefault("LOG_FILE_PATH", "logs/all.txt"),
			SwaggerFilePath: getEnvOrDefault("SWAGGER_FILE_PATH", "docs/swagger/swagger.json"),
			SwaggerHTMLFilePath: getEnvOrDefault("SWAGGER_HTML_FILE_PATH", "docs/swagger/swagger.html"),

			FrontendURL: getEnvOrDefault("FRONTEND_URL", "http://localhost:2022"),
		}
	}

	return config
}

func getEnvOrDefault(key string, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}

func getEnvAsInt(key string, defaultValue int) int {
	valueStr := os.Getenv(key)
	if valueStr == "" {
		return defaultValue
	}
	value, err := strconv.Atoi(valueStr)
	if err != nil {
		return defaultValue
	}
	return value
}