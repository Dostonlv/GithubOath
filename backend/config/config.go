package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/spf13/cast"
	"os"
)

type Config struct {
	ClientId     string
	ClientSecret string
}

func Load() Config {
	if err := godotenv.Load("/backend/.env"); err != nil {
		fmt.Println("No .env file found")
	}
	cfg := Config{}
	cfg.ClientSecret = cast.ToString(getOrReturnDefaultValue("CLIENT_SECRET", ""))
	cfg.ClientId = cast.ToString(getOrReturnDefaultValue("CLIENT_ID", ""))
	return cfg
}

func getOrReturnDefaultValue(key string, defaultValue interface{}) interface{} {
	val, exists := os.LookupEnv(key)

	if exists {
		return val
	}

	return defaultValue
}
