package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func LoadEnv(fileLocation string) (string, string) {
	if err := godotenv.Load(fileLocation); err != nil {
		log.Fatalf("failed to load .env, err:%v", err)
	}
	dbUrl := os.Getenv("DATABASE_URL")
	port := os.Getenv("PORT")
	return dbUrl, port
}
