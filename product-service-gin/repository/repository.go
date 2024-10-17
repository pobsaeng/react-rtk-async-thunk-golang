package repository

import (
	"database/sql"
	"fmt"
	"log"
	"time"
)

var Db *sql.DB

func InitDB(dataSourceName string) {
	var err error
	Db, err = sql.Open("mysql", dataSourceName)
	if err != nil {
		log.Fatal(err)
	}

	Db.SetMaxOpenConns(10)
	Db.SetMaxIdleConns(5)
	Db.SetConnMaxLifetime(5 * time.Minute)

	err = Db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MySQL!")
}
