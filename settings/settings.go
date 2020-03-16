package settings

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// EnvFile must be set to the address of the main .env file.
const EnvFile string = ".env"

var params = NewParams(EnvFile)

// Port exposes the port where the application is served.
var Port = params.port

// TmplDir exposes the address of the templates folder.
var TmplDir string = params.tmplDir

// Params struct holds the application settings parameters
type Params struct {
	// port exposes the port where the application is served.
	port string
	// tmplDir must be set to the address of the templates folder.
	tmplDir string
}

// NewParams reads env file then initialize a new Params object
func NewParams(f string) Params {
	err := godotenv.Load(f)
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	port := getEnvParam("PORT")
	tmplDir := getEnvParam("TMPLDIR")

	return Params{port: port, tmplDir: tmplDir}
}

func getEnvParam(s string) string {
	p := os.Getenv(s)
	if p == "" {
		log.Fatalf("%s must be set. Consider verify your EnvFile.", p)
	}
	return p
}
