package database

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/MertJSX/forum-website/server/types"
)

func CreateUsersTable(db *sql.DB) {
	sqlStmt := `
	CREATE TABLE IF NOT EXISTS users (
		id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		username TEXT NOT NULL UNIQUE,
		email TEXT NOT NULL UNIQUE,
		password TEXT NOT NULL
	);`
	_, err := db.Exec(sqlStmt)
	if err != nil {
		log.Printf("%q: %s\n", err, sqlStmt)
		return
	}
}

func CreateForumsTable(db *sql.DB) {
	sqlStmt := `
	CREATE TABLE IF NOT EXISTS forums (
		id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		user_id TEXT,
		title TEXT,
		description TEXT,
		created_at TEXT,
		content TEXT
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

	ifUsernameExists, _ := CheckIfUsernameExists(db, usr.Name)
	if ifUsernameExists {
		fmt.Println("Username already exists")
		return fmt.Errorf("username already exists")
	}

	ifEmailExists, _ := CheckIfEmailExists(db, usr.Email)
	if ifEmailExists {
		fmt.Println("Email already exists")
		return fmt.Errorf("email already exists")
	}

	stmt, err := tx.Prepare("INSERT INTO users(username, email, password) values(?, ?, ?)")
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

type SearchForUsersBy int

const (
	ByEmail SearchForUsersBy = iota
	ByUsername
	ByPassword
	ByAll
	ByUsernameAndPassword
	ByUsernameAndEmail
	ByEmailAndPassword
)

func SearchForUsers(
	db *sql.DB,
	usr types.User,
	searchBy SearchForUsersBy) ([]types.User, error) {
	var foundList []types.User

	// rows, err := db.Query("SELECT * FROM users WHERE username = ?", searchByItem)
	var rows *sql.Rows
	var err error
	switch searchBy {
	case ByEmail:
		rows, err = db.Query("SELECT * FROM users WHERE email = ?", usr.Email)
	case ByUsername:
		rows, err = db.Query("SELECT * FROM users WHERE username = ?", usr.Name)
	case ByPassword:
		rows, err = db.Query("SELECT * FROM users WHERE password = ?", usr.Password)
	case ByAll:
		rows, err = db.Query("SELECT * FROM users WHERE username = ? AND email = ? AND password = ?",
			usr.Name, usr.Email, usr.Password)
	case ByUsernameAndEmail:
		rows, err = db.Query("SELECT * FROM users WHERE username = ? AND email = ?",
			usr.Name, usr.Email)
	case ByEmailAndPassword:
		rows, err = db.Query("SELECT * FROM users WHERE email = ? AND password = ?",
			usr.Email, usr.Password)
	case ByUsernameAndPassword:
		rows, err = db.Query("SELECT * FROM users WHERE username = ? AND password = ?",
			usr.Name, usr.Password)
	}

	fmt.Println("First Possible Error")

	if err != nil {
		fmt.Println(err)
		return nil, fmt.Errorf("SearchForUsers Error %v: %v", usr, err)
	}
	defer rows.Close()

	fmt.Println("Second Possible Error")

	fmt.Println(rows.Columns())

	for rows.Next() {
		var usr types.User
		if err := rows.Scan(&usr.ID, &usr.Name, &usr.Email, &usr.Password); err != nil {
			fmt.Println(err)
			return nil, fmt.Errorf("SearchForUsers %v: %v", usr, err)
		}
		foundList = append(foundList, usr)
	}

	fmt.Println("Third Possible Error")

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("SearchForUsers %v: %v", usr, err)
	}
	return foundList, nil
}
