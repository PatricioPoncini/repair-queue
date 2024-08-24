// Package config provides configuration management for the application.
package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

// Config contains the environment keys for the project.
type Config struct {
	PublicHost             string
	Port                   string
	DBUser                 string
	DBPassword             string
	DBAdress               string
	DBName                 string
	JWTExpirationInSeconds int64
	JWTSecret              string
}

// Envs is a global variable that holds the configuration settings.
var Envs = initConfig()

func initConfig() Config {
	_ = godotenv.Load()

	return Config{
		PublicHost:             getEnv("PUBLIC_HOST", "http://localhost"),
		Port:                   getEnv("PORT", "8080"),
		DBUser:                 getEnv("DB_USER", "root"),
		DBPassword:             getEnv("DB_PASSWORD", "myfakepassword"),
		DBAdress:               fmt.Sprintf("%s:%s", getEnv("DB_HOST", "127.0.0.1"), getEnv("DB_PORT", "3306")),
		DBName:                 getEnv("DB_NAME", "repair_queue"),
		JWTExpirationInSeconds: getEnvAsInt("JWT_EXPIRATION_TIME", 3600*24*7),
		JWTSecret:              getEnv("JWT_SECRET", "jwt_secret"),
	}
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

func getEnvAsInt(key string, fallback int64) int64 {
	if value, ok := os.LookupEnv(key); ok {
		i, err := strconv.ParseInt(value, 10, 64)
		if err != nil {
			return fallback
		}
		return i
	}
	return fallback
}
