package main

import (
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/rjNemo/go-wiki/controller"
	"github.com/rjNemo/go-wiki/settings"
)

var port string = readEnv(settings.EnvFile)

func main() {
	startServer(port, controller.Router)
}

func startServer(p string, r func()) {
	log.Printf("Start Go-wiki server on http://localhost:%s\n", port)
	port := ":" + p
	r()
	log.Fatal(http.ListenAndServe(port, nil))
}

func readEnv(f string) string {
	err := godotenv.Load(f)
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("PORT must be set. Consider verify your EnvFile.")
	}
	return port
}
