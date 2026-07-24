package utils

import (
	"log"
	"os"
	"sync"

	"github.com/joho/godotenv"
)

var loadDotEnvOnce sync.Once

func GetEnvVariable(key string) string {
	value := os.Getenv(key)
	if value != "" {
		return value
	}

	// Not set in the process environment — fall back to a .env file, which is
	// how local development supplies config. Containers get real environment
	// variables and ship no .env, so a missing file is expected there and must
	// not take the process down. Loaded once: this path runs on every lookup
	// of an unset key, including per-request ones.
	loadDotEnvOnce.Do(func() {
		if err := godotenv.Load(".env"); err != nil {
			log.Printf("no .env file loaded (%v); using the process environment only", err)
		}
	})

	return os.Getenv(key)
}
