package config

import (
	"log"

	"github.com/spf13/viper"
)

var viperContext *viper.Viper

func init() {
	log.Println("Initializing viper context...")
	viperContext = viper.New()

	viperContext.SetEnvPrefix("NAASGUL")
	viperContext.AutomaticEnv()
}

//GetViperContext provides access to the viper instance created during startup.
func GetViperContext() *viper.Viper {
	return viperContext
}

//GetMysqlUserName gets the configured username used for connecting to mysql.
func GetMysqlUserName() string {
	return viperContext.GetString("MYSQLUSERNAME")
}

//GetMysqlPassword gets the configured password used for connecting to mysql.
func GetMysqlPassword() string {
	return viperContext.GetString("MYSQLPASSWORD")
}

//GetMysqlHost gets the configured host currently serving the mysql database.
func GetMysqlHost() string {
	return viperContext.GetString("MYSQLHOST")
}

//GetMysqlPort gets the configured port that the mysql host is configured to communicate over.
func GetMysqlPort() string {
	return viperContext.GetString("MYSQLPORT")
}

//GetMysqlDatabase gets the configured database name that we want to connect to.
func GetMysqlDatabase() string {
	return viperContext.GetString("MYSQLDATABASE")
}

//GetServerHost gets the configured hostname to be used for the webserver.
func GetServerHost() string {
	return viperContext.GetString("HOST")
}

//GetServerPort gets the configured port to be used for the webserver.
func GetServerPort() string {
	return viperContext.GetString("PORT")
}
