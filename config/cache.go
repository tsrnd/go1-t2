package config

import (
	"os"
	"log"

	"github.com/joho/godotenv"
	"goweb2/services/cache"
	"goweb2/services/cache/redis"
)

// Cache func
func Cache() cache.Cache {
	er := godotenv.Load()
	if er != nil {
		log.Print("Error loading .env file")
	}
	return redis.Connect(
		os.Getenv("REDIS_ADDR"),
		os.Getenv("REDIS_PASSWORD"),
		0,
	)
}
