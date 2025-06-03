package postgres

import (
	"database/sql"
	"fmt"
	"go-api/config"

	_ "github.com/lib/pq"
)

func Connect() *sql.DB {
	stringConnectionDB := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%d sslmode=%s",
		config.ConfigInstance.Db.Host,
		config.ConfigInstance.Db.User,
		config.ConfigInstance.Db.Password,
		config.ConfigInstance.Db.DBName,
		config.ConfigInstance.Db.Port,
		config.ConfigInstance.Db.SSLMode,
		//config.ConfigInstance.Db.TimeZone,
	)

	db, err := sql.Open("postgres", stringConnectionDB)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Connected to " + config.ConfigInstance.Db.Host)

	return db
}
