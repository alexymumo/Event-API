package db

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func Connect() *sql.DB {
	err := godotenv.Load()
	if err != nil {
		fmt.Printf("Failed to load env")
	} else {
		fmt.Println("Successfully loaded environment variables")
	}
	dburl := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", os.Getenv("DB_USER"), os.Getenv("DB_PASS"), os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_NAME"))
	DB, err := sql.Open("mysql", dburl)
	if err != nil {
		fmt.Println("Failed to connect to database")
	} else {
		fmt.Println("Connected successfully")
	}
	return DB
}
