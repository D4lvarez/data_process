package utils

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func GetDotEnvVar(key string) string {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %s", err)
	}

	v, ok := os.LookupEnv(key)

	if !ok {
		log.Fatalf("Env %s don't set", key)
	}

	return v
}
