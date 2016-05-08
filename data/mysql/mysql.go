package mysql

import (
	"fmt"
	"log"

	"github.com/eriktate/NaaSgul/config"
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

//NotificationProvider is a struct that provides access to Noti
type NotificationProvider struct {
	db *sqlx.DB
}

//NewNotificationProvider returns a notification provider
func NewNotificationProvider(user, password, host, port, database string) (*NotificationProvider, error) {
	if db != nil {
		return &NotificationProvider{db}, nil
	}

	username := config.GetMysqlUserName()
	pword := config.GetMysqlPassword()
	hostname := config.GetMysqlHost()
	listenPort := config.GetMysqlPort()
	dbname := config.GetMysqlDatabase()

	connectionString := fmt.Sprintf("%s:%s@(%s:%s)/%s", username, pword, hostname, listenPort, dbname)
	err := Connect(connectionString)

	if err != nil {
		return nil, err
	}

	return &NotificationProvider{db}, nil
}
