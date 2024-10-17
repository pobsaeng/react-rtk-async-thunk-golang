package config

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

var Db *sql.DB

func InitDB() {
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	name := os.Getenv("DB_NAME")
	
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		user, password, host, port, name)

	var err error
	Db, err = sql.Open("mysql", dataSourceName)
	if err != nil {
		log.Fatal(err)
	}

	// Set maximum number of open connections
	Db.SetMaxOpenConns(10)

	// Set maximum number of idle connections
	Db.SetMaxIdleConns(5)

	// Set the maximum lifetime of a connection
	Db.SetConnMaxLifetime(5 * time.Minute)

	err = Db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MySQL!")
}
