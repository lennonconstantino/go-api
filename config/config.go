package config

import (
	"os"
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
		Port               int
		SecretKey          string
		Environment        string
		CORSAllowedOrigins []string
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

		corsAllowedOrigins := viper.GetStringSlice("cors_allowed_origins")
		if len(corsAllowedOrigins) == 0 {
			corsAllowedOrigins = []string{"*"}
		}

		ConfigInstance.Server.CORSAllowedOrigins = corsAllowedOrigins
		ConfigInstance.Server.SecretKey = os.Getenv(viper.GetString(""))
	})

	return ConfigInstance
}
