package lib

import (
	"os"

	"github.com/joho/godotenv"
)

func GetenvFromSecretfile(key string) string {
	godotenv.Load()

	if os.Getenv(key) != "" {
		return os.Getenv(key)
	}
	path := os.Getenv(key + "_FILE")
	if path == "" {
		return ""
	}
	file, err := os.ReadFile(path)
	if err != nil {
		return ""
	}
	return string(file)
}
