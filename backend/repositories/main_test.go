package repositories_test

import (
	"database/sql"
	"fmt"
	"os"
	"os/exec"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

var (
	dbUser     = os.Getenv("DB_USER")
	dbPassword = os.Getenv("DB_PASSWORD")
	dbHost     = os.Getenv("DB_HOST")
	dbPort     = os.Getenv("DB_PORT")
	dbDatabase = os.Getenv("DB_DATABASE")
	dbConfig   = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", dbUser, dbPassword, dbHost, dbPort, dbDatabase)
)

func connectDB() error {
	var err error
	db, err = sql.Open("mysql", dbConfig)
	if err != nil {
		return err
	}

	return nil
}

func setupTestData() error {
	cmd := exec.Command("mysql", "-h", dbHost, "-u", dbUser, dbDatabase, fmt.Sprintf("--password=%s", dbPassword), "-e", "source ./testdata/setupDB.sql")

	err := cmd.Run()
	if err != nil {
		return err
	}

	return nil
}

func cleanupDB() error {
	//fmt.Println(os.Getwd())

	cmd := exec.Command("mysql", "-h", dbHost, "-u", dbUser, dbDatabase, fmt.Sprintf("--password=%s", dbPassword), "-e", "source ./testdata/cleanupDB.sql")

	// fmt.Println(cmd)

	var err error

	err = cmd.Run()
	if err != nil {
		return err
	}

	return nil
}

func setup() error {
	err := connectDB()
	if err != nil {
		fmt.Println("connect:", err)
		return err
	}

	err = cleanupDB()
	if err != nil {
		fmt.Println("cleanup:", err)
		return err
	}

	err = setupTestData()
	if err != nil {
		fmt.Println("setup:", err)
		return err
	}

	return nil
}

func teardown() error {
	err := cleanupDB()
	if err != nil {
		return err
	}

	db.Close()

	return nil
}

func TestMain(m *testing.M) {
	err := setup()
	if err != nil {
		os.Exit(1)
	}

	m.Run()

	err = teardown()
	if err != nil {
		os.Exit(1)
	}
}
