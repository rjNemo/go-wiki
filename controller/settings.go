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
	return os.Getenv("PORT")
}
