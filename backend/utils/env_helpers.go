package utils

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

func GetEnvVariable(key string) string {

	if value, exists := os.LookupEnv("SQLAZURECONNSTR_" + key); exists {
		return value
	}

	value := os.Getenv(key)
	if value != "" {
		return value
	}

	// If not set, try to load from .env file (for local development)
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Warning: Error loading .env file: %v", err)
		// Return empty string or handle this case as needed
	}

	return os.Getenv(key)
}
