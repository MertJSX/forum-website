package tests

import (
	"database/sql"
	"log"
	"os"
	"testing"
)

var db *sql.DB

func TestMain(m *testing.M) {
	log.Print("Setting up test database")
	db = SetupTestDB()
	code := m.Run()
	defer db.Close()
	defer os.Exit(code)
}

func GetTestDB() *sql.DB {
	if db == nil {
		log.Fatal("DB is nil")
	}
	return db
}
