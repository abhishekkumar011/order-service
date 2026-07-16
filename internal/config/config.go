package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port           string
	MongoURI       string
	DatabaseName   string
	CollectionName string
}

func LoadConfig() Config {

	err := godotenv.Load()

	if err != nil {
		log.Println(".env file not found, using system environment")
	}

	return Config{
		Port:           os.Getenv("PORT"),
		MongoURI:       os.Getenv("MONGO_URI"),
		DatabaseName:   os.Getenv("DATABASE_NAME"),
		CollectionName: os.Getenv("COLLECTION_NAME"),
	}
}

/*
	1. package config - This file belongs to the config package
	2. struct - A struct is Go's primary way of grouping related data. (like TypeScript interface plus an object)
	3. func LoadConfig() Config - function name: LoadConfig, Return type: Config
	4. os.Getenv("PORT") - Reading environment variables (like process.env.PORT)
	5. os.Getenv - like process.env
*/