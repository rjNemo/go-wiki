package controller

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var Port string = readEnv(".env")
var tmplDir string = "templates/"

func readEnv(f string) string {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("PORT must be set")
	}
	return port
}
