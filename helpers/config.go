package helpers

import (
	"os"

	"github.com/joho/godotenv"
)

var Env = map[string]string{}

func SetupConfig() {
	var err error
	Env, err = godotenv.Read(".env")
	if err != nil {
		Env = map[string]string{}
	}
}

func GetEnv(key string, val string) string {
	valFromOs := os.Getenv(key)
	if valFromOs != "" {
		return valFromOs
	}

	result := Env[key]
	if result == "" {
		result = val
	}

	return result
}
