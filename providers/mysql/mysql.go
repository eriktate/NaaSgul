package mysql

import (
	"log"

	//Need to do a blank import for sqlx
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var db *sqlx.DB

//Connect establishes communication with the mysql data source.
func Connect(connectionString string) error {
	var err error
	log.Println("Connecting to mysql...")
	db, err = sqlx.Connect("mysql", connectionString)
	if err != nil {
		return err
	}

	log.Println("Successfully connected to mysql!")
	return nil
}

//GetDatabaseConnection returns the database struct created upon connection.
func GetDatabaseConnection() *sqlx.DB {
	return db
}
