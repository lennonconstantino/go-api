package config

/*
import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	// see the container
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

// Load initialize environment variables
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
*/

import (
	"strings"
	"sync"

	"github.com/spf13/viper"
)

type (
	Config struct {
		Server *Server
		Db     *Db
	}

	Server struct {
		Port      int
		SecretKey string
	}

	Db struct {
		Host     string
		Port     int
		User     string
		Password string
		DBName   string
		SSLMode  string
		TimeZone string
	}
)

var (
	once           sync.Once
	ConfigInstance *Config
)

func GetConfig() *Config {
	once.Do(func() {
		viper.SetConfigName("config")
		viper.SetConfigType("yaml")
		viper.AddConfigPath("./")
		viper.AutomaticEnv()
		viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

		if err := viper.ReadInConfig(); err != nil {
			panic(err)
		}

		if err := viper.Unmarshal(&ConfigInstance); err != nil {
			panic(err)
		}
	})

	return ConfigInstance
}
