package db

import (
	"database/sql"
	"fmt"
	"go-api/config"

	_ "github.com/lib/pq"
)

func ConnectDB() (*sql.DB, error) {
	psqlInfo := config.StringConnectionDB

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Connected to " + config.Dbname)

	return db, nil
}
