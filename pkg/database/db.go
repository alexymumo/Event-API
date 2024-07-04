package database

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

var db *sql.DB
var err error

func Connect() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Failed to load env values")
	}
	var (
		dbuser     = os.Getenv("DB_USER")
		dbpassword = os.Getenv("DB_PASS")
		dbhost     = os.Getenv("DB_HOST")
		dbport     = os.Getenv("DB_PORT")
		dbname     = os.Getenv("DB_NAME")
	)

	dburl := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", dbuser, dbpassword, dbhost, dbport, dbname)
	d, err := sql.Open("mysql", dburl)
	if err != nil {
		fmt.Println("Failed to connect")
	} else {
		fmt.Println("Connected successfully")
	}
	db = d
}

func GetDb() *sql.DB {
	return db
}
