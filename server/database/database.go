package database

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/MertJSX/forum-website/server/types"
)

func CreateUserTable(db *sql.DB) {
	sqlStmt := `CREATE TABLE IF NOT EXISTS users (
		id INTEGER NOT NULL PRIMARY KEY  AUTOINCREMENT,
		username TEXT,
		email TEXT,
		password TEXT
	);`
	_, err := db.Exec(sqlStmt)
	if err != nil {
		log.Printf("%q: %s\n", err, sqlStmt)
		return
	}
}

func CreateNewUser(db *sql.DB, usr types.User) error {
	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
		return fmt.Errorf("Begin transaction error: %w", err)
	}
	stmt, err := tx.Prepare("insert into users(username, email, password) values(?, ?, ?)")
	if err != nil {
		log.Fatal(err)
		return fmt.Errorf("Prepare statement error: %w", err)
	}
	defer stmt.Close()
	_, err = stmt.Exec(usr.Name, usr.Email, usr.Password)
	if err != nil {
		log.Fatal(err)
		return fmt.Errorf("Statement execute error: %w", err)
	}
	err = tx.Commit()
	if err != nil {
		log.Fatal(err)
		return fmt.Errorf("Transaction commit error: %w", err)
	}
	return nil
}
