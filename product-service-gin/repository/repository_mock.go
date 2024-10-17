package repository

import (
	"database/sql"
	"log"

	"github.com/DATA-DOG/go-sqlmock"
	_ "github.com/go-sql-driver/mysql"
	"pobsaeng.com/product-api/config"
)

var MockDB *sql.DB
var mock sqlmock.Sqlmock

func MockInit() {
	var err error
	MockDB, mock, err = sqlmock.New()
	if err != nil {
		log.Fatalf("Error initializing mock database: %s", err)
	}

	// Set the repository.Db to the mock database
	config.Db = MockDB

	// Define the expected queries and their respective responses
	// For example:
	// mock.ExpectQuery("SELECT id, name, image, price, store, type FROM products").
	//     WillReturnRows(sqlmock.NewRows([]string{"id", "name", "image", "price", "store", "type"}).
	//     AddRow("1", "Product A", "imageA.jpg", 10.0, 100, "Type A"))

	// Initialize tables or mock data if needed
	// For example:
	// MockDB.Exec(`CREATE TABLE products (id VARCHAR(255) PRIMARY KEY, name VARCHAR(255), image VARCHAR(255), price DOUBLE, store INT, type VARCHAR(255));`)
}

func MockClose() {
	MockDB.Close()
}
