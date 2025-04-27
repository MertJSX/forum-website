package tests

import (
	"database/sql"
	"log"
	"testing"

	"github.com/MertJSX/forum-website/server/database"
	"github.com/MertJSX/forum-website/server/types"

	_ "github.com/mattn/go-sqlite3"
)

func SetupTestDB() *sql.DB {
	var err error
	db, err = sql.Open("sqlite3", ":memory:")
	if err != nil {
		log.Fatalf("Failed to open test database: %v", err)
		return nil
	}

	database.CreateForumsTable(db)

	database.CreateUsersTable(db)

	testUser := types.User{
		Name:     "exampleexistinguser",
		Email:    "exampleexistingemail@gmail.com",
		Password: "password123",
	}

	database.CreateNewUser(db, testUser)

	return db
}

func TestDBFunctions(t *testing.T) {
	var err error
	db, err = sql.Open("sqlite3", ":memory:")
	if err != nil {
		t.Fatalf("Failed to open test database: %v", err)
		return
	}

	t.Run("Create forums table", func(t *testing.T) {
		database.CreateForumsTable(db)
	})

	t.Run("Create users table", func(t *testing.T) {
		database.CreateUsersTable(db)
	})

	t.Run("Create test user", func(t *testing.T) {
		testUser := types.User{
			Name:     "exampleexistinguser",
			Email:    "exampleexistingemail@gmail.com",
			Password: "password123",
		}
		err := database.CreateNewUser(db, testUser)
		if err != nil {
			t.Fatalf("Failed to create test user: %v", err)
		}
	})

	t.Cleanup(func() {
		db.Close()
	})
}
