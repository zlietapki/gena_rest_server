// start name:top
package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	//start name:struct type:add
	HTTPListen string
	//start name:body
	Env string
}

func New() *Config {
	godotenv.Load()

	return &Config{
		//start name:return type:add
		HTTPListen: getEnv("HTTP_LISTEN", ":8080"),
		//start name:post_return
		Env: getEnv("ENV", "local"),
	}
}

func getEnv(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}

	return fallback
}
