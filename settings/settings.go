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

// ConnStr exposes the connection string to the database.
var ConnStr string = params.connStr

// Params struct holds the application settings parameters
type Params struct {
	// port exposes the port where the application is served.
	port string
	// tmplDir must be set to the address of the templates folder.
	tmplDir string
	// connStr must be set to the connection string to the database.
	connStr string
}

// NewParams reads env file then initialize a new Params object
func NewParams(f string) Params {
	err := godotenv.Load(f)
	if err != nil {
		log.Println("Error loading .env file")
	}
	port := getEnvParam("PORT")
	tmplDir := getEnvParam("TMPLDIR")
	connStr := getEnvParam("ConnectionString")

	return Params{port: port, tmplDir: tmplDir, connStr: connStr}
}

func getEnvParam(s string) string {
	p := os.Getenv(s)
	if p == "" {
		log.Fatalf("%s must be set. Consider verify your EnvFile.", p)
	}
	return p
}
