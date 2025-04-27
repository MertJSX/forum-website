package database

import (
	"database/sql"
	"fmt"
)

func CheckIfUsernameExists(db *sql.DB, username string) (bool, error) {
	var exists bool
	query := "SELECT EXISTS(SELECT 1 FROM users WHERE username = ?)"
	err := db.QueryRow(query, username).Scan(&exists)
	if err != nil {
		return false, fmt.Errorf("error checking if username exists: %v", err)
	}
	return exists, nil
}

func CheckIfEmailExists(db *sql.DB, email string) (bool, error) {
	var exists bool
	query := "SELECT EXISTS(SELECT 1 FROM users WHERE email = ?)"
	err := db.QueryRow(query, email).Scan(&exists)
	if err != nil {
		return false, fmt.Errorf("error checking if email exists: %v", err)
	}
	return exists, nil
}
