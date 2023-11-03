package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/yoshi-zen/sea-turtle/backend/api"
)

var (
	db         *sql.DB
	dbUser     string
	dbPassword string
	dbHost     string
	dbPort     string
	dbDatabase string
	dbConfig   string
)

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
		return
	}

	dbUser = os.Getenv("DB_USER")
	dbPassword = os.Getenv("DB_PASSWORD")
	dbHost = os.Getenv("DB_HOST")
	dbPort = os.Getenv("DB_PORT")
	dbDatabase = os.Getenv("DB_DATABASE")
	dbConfig = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", dbUser, dbPassword, dbHost, dbPort, dbDatabase)
}

func connectDB() (*sql.DB, error) {
	var err error
	db, err = sql.Open("mysql", dbConfig)
	if err != nil {
		return nil, err
	}

	return db, nil
}

func main() {
	db, err := connectDB()
	if err != nil {
		fmt.Println(err)
		return
	}

	r := api.NewRouter(db)

	log.Println("server start at port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
