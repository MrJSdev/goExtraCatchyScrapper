package helpers

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// GetEnv returns the value of the environment variable named by the key.
// It returns the value, which will be empty if the variable is not present.
func GetEnv(key string) string {
	return os.Getenv(key)
}

// LoadEnvVariableFromFile loads the environment variables from the .env file
func LoadEnvVariableFromFile() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
}
