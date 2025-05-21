package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	// vai enxergar o container
	postDB             = ""
	portDB             = 0
	userDB             = ""
	passwordDB         = ""
	Dbname             = ""
	StringConnectionDB = ""

	PortAPI = 0
	// SecretKey is the key that will be used to sign the token
	SecretKey []byte
)

// Carregar vai inicializar as variaveis de ambiente
func Load() {
	var err error

	if err = godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	PortAPI, err = strconv.Atoi(os.Getenv("API_PORT"))
	if err != nil {
		PortAPI = 8000
	}

	postDB = os.Getenv("DB_HOST")
	portDB, err = strconv.Atoi(os.Getenv("DB_PORT"))
	if err != nil {
		portDB = 5432
	}
	userDB = os.Getenv("DB_USER")
	passwordDB = os.Getenv("DB_PASSWORD")
	Dbname = os.Getenv("DB_NAME")

	StringConnectionDB =
		fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
			postDB, portDB, userDB, passwordDB, Dbname)

	SecretKey = []byte(os.Getenv("SECRET_KEY"))
}
