package database

import (
	"database/sql"
	"fmt"
	"log"
	"time"

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

func CreateFollowersTable(db *sql.DB) {
	sqlStmt := `
	CREATE TABLE IF NOT EXISTS followers (
		id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		follower_id INTEGER NOT NULL,
		followed_id INTEGER NOT NULL,
		FOREIGN KEY(follower_id) REFERENCES users(id) ON DELETE CASCADE ON UPDATE CASCADE,
		FOREIGN KEY(followed_id) REFERENCES users(id) ON DELETE CASCADE ON UPDATE CASCADE
	);`
	_, err := db.Exec(sqlStmt)
	if err != nil {
		log.Printf("%q: %s\n", err, sqlStmt)
		return
	}
}

func CreatePostsTable(db *sql.DB) {
	sqlStmt := `
	CREATE TABLE IF NOT EXISTS posts (
		id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		user_id INTEGER,
		title TEXT,
		created_at TEXT,
		content TEXT,
		FOREIGN KEY(user_id) REFERENCES users(id) ON DELETE CASCADE ON UPDATE CASCADE
	);`
	_, err := db.Exec(sqlStmt)
	if err != nil {
		log.Printf("%q: %s\n", err, sqlStmt)
		return
	}
}

func CreateUpvotesTable(db *sql.DB) {
	sqlStmt := `
	CREATE TABLE IF NOT EXISTS post_upvotes (
		id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		user_id INTEGER NOT NULL,
		post_id INTEGER NOT NULL,
		FOREIGN KEY(user_id) REFERENCES users(id) ON DELETE CASCADE ON UPDATE CASCADE,
		FOREIGN KEY(post_id) REFERENCES posts(id) ON DELETE CASCADE ON UPDATE CASCADE,
		UNIQUE(user_id, post_id)
	);`
	_, err := db.Exec(sqlStmt)
	if err != nil {
		log.Printf("%q: %s\n", err, sqlStmt)
		return
	}
}

func CreateCommentsTable(db *sql.DB) {
	sqlStmt := `
	CREATE TABLE IF NOT EXISTS comments (
		id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		user_id INTEGER,
		post_id INTEGER,
		comment TEXT,
		created_at TEXT,
		FOREIGN KEY(user_id) REFERENCES users(id) ON DELETE CASCADE ON UPDATE CASCADE,
		FOREIGN KEY(post_id) REFERENCES posts(id) ON DELETE CASCADE ON UPDATE CASCADE
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

func IsPostAuthor(db *sql.DB, userID, postID string) (bool, error) {
	var exists bool
	err := db.QueryRow(`
		SELECT EXISTS(
			SELECT 1 FROM posts WHERE id = ? AND user_id = ?
		)
	`, postID, userID).Scan(&exists)
	if err != nil {
		return false, fmt.Errorf("IsPostAuthor error: %w", err)
	}
	return exists, nil
}

func DeletePost(db *sql.DB, postID string) error {
	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
		return fmt.Errorf("Begin transaction error: %w", err)
	}

	// Delete the post
	_, err = tx.Exec(`
		DELETE FROM posts WHERE id = ?
	`, postID)
	if err != nil {
		tx.Rollback()
		log.Fatal(err)
		return fmt.Errorf("Delete post error: %w", err)
	}

	// Commit the transaction
	err = tx.Commit()
	if err != nil {
		log.Fatal(err)
		return fmt.Errorf("Transaction commit error: %w", err)
	}

	return nil
}

func IsUserFollowing(db *sql.DB, followerID, followedID string) (bool, error) {
	var exists bool
	err := db.QueryRow(`
		SELECT EXISTS(
			SELECT 1 FROM followers WHERE follower_id = ? AND followed_id = ?
		)
	`, followerID, followedID).Scan(&exists)
	if err != nil {
		return false, fmt.Errorf("IsUserFollowing error: %w", err)
	}
	return exists, nil
}

func FollowUser(db *sql.DB, followerID, followedID string) (int, error) {
	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
		return 0, fmt.Errorf("Begin transaction error: %w", err)
	}

	// Check if the follow relationship already exists
	var exists bool
	err = tx.QueryRow(`
		SELECT EXISTS(
			SELECT 1 FROM followers WHERE follower_id = ? AND followed_id = ?
		)
	`, followerID, followedID).Scan(&exists)
	if err != nil {
		tx.Rollback()
		log.Fatal(err)
		return 0, fmt.Errorf("QueryRow error: %w", err)
	}

	if exists {
		// If it exists, delete the follow relationship
		_, err = tx.Exec(`
			DELETE FROM followers WHERE follower_id = ? AND followed_id = ?
		`, followerID, followedID)
		if err != nil {
			tx.Rollback()
			log.Fatal(err)
			return 0, fmt.Errorf("Delete follow error: %w", err)
		}
	} else {
		// If it does not exist, create the follow relationship
		_, err = tx.Exec(`
			INSERT INTO followers (follower_id, followed_id) VALUES (?, ?)
		`, followerID, followedID)
		if err != nil {
			tx.Rollback()
			log.Fatal(err)
			return 0, fmt.Errorf("Insert follow error: %w", err)
		}
	}

	// Commit the transaction
	err = tx.Commit()
	if err != nil {
		log.Fatal(err)
		return 0, fmt.Errorf("Transaction commit error: %w", err)
	}

	// Get the current followers count for the followed user
	var followersCount int
	err = db.QueryRow(`
		SELECT COUNT(*) FROM followers WHERE followed_id = ?
	`, followedID).Scan(&followersCount)
	if err != nil {
		log.Fatal(err)
		return 0, fmt.Errorf("Count followers error: %w", err)
	}

	return followersCount, nil
}

func CreateNewPost(db *sql.DB, post *types.Post) (int64, error) {
	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
		return 0, fmt.Errorf("Begin transaction error: %w", err)
	}

	stmt, err := tx.Prepare("INSERT INTO posts(user_id, title, created_at, content) values(?, ?, ?, ?)")
	if err != nil {
		log.Fatal(err)
		return 0, fmt.Errorf("Prepare statement error: %w", err)
	}
	defer stmt.Close()

	result, err := stmt.Exec(post.UserId, post.Title, time.Now().Unix(), post.Content)
	if err != nil {
		tx.Rollback()
		log.Fatal(err)
		return 0, fmt.Errorf("Statement execute error: %w", err)
	}

	postID, err := result.LastInsertId()
	if err != nil {
		tx.Rollback()
		log.Fatal(err)
		return 0, fmt.Errorf("Retrieve last insert ID error: %w", err)
	}

	err = tx.Commit()
	if err != nil {
		log.Fatal(err)
		return 0, fmt.Errorf("Transaction commit error: %w", err)
	}

	return postID, nil
}

func GetPostsByFollowedUsers(db *sql.DB, userID string) ([]types.Post, error) {
	rows, err := db.Query(`
		SELECT posts.id, posts.user_id, posts.title, posts.created_at, posts.content, users.username 
		FROM posts
		INNER JOIN users ON posts.user_id = users.id
		INNER JOIN followers ON followers.followed_id = posts.user_id
		WHERE followers.follower_id = ?
		ORDER BY posts.created_at DESC
	`, userID)

	if err != nil {
		fmt.Println(err)
		return nil, fmt.Errorf("GetPostsByFollowedUsers Error: %v", err)
	}

	defer rows.Close()

	var posts []types.Post

	for rows.Next() {
		var post types.Post
		if err := rows.Scan(&post.ID, &post.UserId, &post.Title, &post.CreatedAt, &post.Content, &post.Author); err != nil {
			fmt.Println(err)
			return nil, fmt.Errorf("GetPostsByFollowedUsers Error %v: %v", post, err)
		}

		// Get the upvotes for the post
		err = db.QueryRow(`
			SELECT COUNT(*) FROM post_upvotes WHERE post_id = ?
		`, post.ID).Scan(&post.Upvotes)
		if err != nil {
			fmt.Println(err)
			return nil, fmt.Errorf("GetPostsByFollowedUsers Upvotes Error %v: %v", post, err)
		}

		posts = append(posts, post)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("GetPostsByFollowedUsers Error %v: %v", posts, err)
	}

	return posts, nil
}

func UpvotePost(db *sql.DB, userID, postID string) (int, error) {
	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
		return 0, fmt.Errorf("Begin transaction error: %w", err)
	}

	// Check if the upvote already exists
	var exists bool
	err = tx.QueryRow(`
		SELECT EXISTS(
			SELECT 1 FROM post_upvotes WHERE user_id = ? AND post_id = ?
		)
	`, userID, postID).Scan(&exists)
	if err != nil {
		tx.Rollback()
		log.Fatal(err)
		return 0, fmt.Errorf("QueryRow error: %w", err)
	}

	if exists {
		// If it exists, delete the upvote
		_, err = tx.Exec(`
			DELETE FROM post_upvotes WHERE user_id = ? AND post_id = ?
		`, userID, postID)
		if err != nil {
			tx.Rollback()
			log.Fatal(err)
			return 0, fmt.Errorf("Delete upvote error: %w", err)
		}
	} else {
		// If it does not exist, create the upvote
		_, err = tx.Exec(`
			INSERT INTO post_upvotes (user_id, post_id) VALUES (?, ?)
		`, userID, postID)
		if err != nil {
			tx.Rollback()
			log.Fatal(err)
			return 0, fmt.Errorf("Insert upvote error: %w", err)
		}
	}

	// Commit the transaction
	err = tx.Commit()
	if err != nil {
		log.Fatal(err)
		return 0, fmt.Errorf("Transaction commit error: %w", err)
	}

	// Get the current upvote count for the post
	var upvoteCount int
	err = db.QueryRow(`
		SELECT COUNT(*) FROM post_upvotes WHERE post_id = ?
	`, postID).Scan(&upvoteCount)
	if err != nil {
		log.Fatal(err)
		return 0, fmt.Errorf("Count upvotes error: %w", err)
	}

	return upvoteCount, nil
}

func CreateNewComment(db *sql.DB, comment *types.Comment) error {
	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
		return fmt.Errorf("Begin transaction error: %w", err)
	}

	stmt, err := tx.Prepare("INSERT INTO comments(user_id, post_id, comment, created_at) values(?, ?, ?, ?)")
	if err != nil {
		log.Fatal(err)
		return fmt.Errorf("Prepare statement error: %w", err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(comment.UserId, comment.PostId, comment.Comment, time.Now().Unix())
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

func GetPosts(db *sql.DB) ([]types.Post, error) {
	rows, err := db.Query(`
		SELECT posts.id, posts.user_id, posts.title, posts.created_at, posts.content, users.username 
		FROM posts 
		INNER JOIN users ON posts.user_id = users.id 
		ORDER BY posts.created_at DESC
	`)

	if err != nil {
		fmt.Println(err)
		return nil, fmt.Errorf("GetPosts Error: %v", err)
	}

	defer rows.Close()

	var posts []types.Post

	for rows.Next() {
		var post types.Post
		if err := rows.Scan(&post.ID, &post.UserId, &post.Title, &post.CreatedAt, &post.Content, &post.Author); err != nil {
			fmt.Println(err)
			return nil, fmt.Errorf("GetPosts Error %v: %v", post, err)
		}

		// Get the upvotes for the post
		err = db.QueryRow(`
			SELECT COUNT(*) FROM post_upvotes WHERE post_id = ?
		`, post.ID).Scan(&post.Upvotes)
		if err != nil {
			fmt.Println(err)
			return nil, fmt.Errorf("GetPosts Upvotes Error %v: %v", post, err)
		}

		posts = append(posts, post)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("GetPosts Error %v: %v", posts, err)
	}

	return posts, nil
}

func GetPostsByUserID(db *sql.DB, userID string) ([]types.Post, error) {
	rows, err := db.Query(`
		SELECT posts.id, posts.user_id, posts.title, posts.created_at, posts.content, users.username 
		FROM posts
		INNER JOIN users ON posts.user_id = users.id 
		WHERE posts.user_id = ?
		ORDER BY posts.created_at DESC
	`, userID)

	if err != nil {
		fmt.Println(err)
		return nil, fmt.Errorf("GetPostsByUserID Error: %v", err)
	}

	defer rows.Close()

	var posts []types.Post

	for rows.Next() {
		var post types.Post
		if err := rows.Scan(&post.ID, &post.UserId, &post.Title, &post.CreatedAt, &post.Content, &post.Author); err != nil {
			fmt.Println(err)
			return nil, fmt.Errorf("GetPostsByUserID Error %v: %v", post, err)
		}

		err = db.QueryRow(`
			SELECT COUNT(*) FROM post_upvotes WHERE post_id = ?
		`, post.ID).Scan(&post.Upvotes)
		if err != nil {
			fmt.Println(err)
			return nil, fmt.Errorf("GetPostsByUserID Upvotes Error %v: %v", post, err)
		}

		posts = append(posts, post)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("GetPostsByUserID Error %v: %v", posts, err)
	}

	return posts, nil
}

func GetPostByID(db *sql.DB, postId string) (*types.Post, error) {
	row := db.QueryRow(`
		SELECT posts.id, posts.user_id, posts.title, posts.created_at, posts.content, users.username 
		FROM posts 
		INNER JOIN users ON posts.user_id = users.id 
		WHERE posts.id = ?
	`, postId)

	var post types.Post
	if err := row.Scan(&post.ID, &post.UserId, &post.Title, &post.CreatedAt, &post.Content, &post.Author); err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("no post found with id %s", postId)
		}
		return nil, fmt.Errorf("GetPostByID Error: %v", err)
	}

	// Get the upvotes for the post
	err := db.QueryRow(`
		SELECT COUNT(*) FROM post_upvotes WHERE post_id = ?
	`, postId).Scan(&post.Upvotes)
	if err != nil {
		return nil, fmt.Errorf("GetPostByID Upvotes Error: %v", err)
	}

	return &post, nil
}

func GetCommentsForPost(db *sql.DB, postId int) ([]types.Comment, error) {
	rows, err := db.Query(`
		SELECT comments.id, comments.user_id, comments.post_id, comments.comment, comments.created_at, users.username 
		FROM comments 
		INNER JOIN users ON comments.user_id = users.id 
		WHERE comments.post_id = ? 
		ORDER BY comments.created_at DESC
	`, postId)
	if err != nil {
		fmt.Println(err)
		return nil, fmt.Errorf("GetCommentsForPost Error: %v", err)
	}
	defer rows.Close()

	var comments []types.Comment

	for rows.Next() {
		var comment types.Comment
		if err := rows.Scan(&comment.ID, &comment.UserId, &comment.PostId, &comment.Comment, &comment.CreatedAt, &comment.Author); err != nil {
			fmt.Println(err)
			return nil, fmt.Errorf("GetCommentsForPost Error %v: %v", comment, err)
		}
		comments = append(comments, comment)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("GetCommentsForPost Error %v: %v", comments, err)
	}

	return comments, nil
}

func UpdatePost(db *sql.DB, postID string, title string, content string) error {
	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
		return fmt.Errorf("Begin transaction error: %w", err)
	}

	_, err = tx.Exec(`
		UPDATE posts
		SET title = ?, content = ?
		WHERE id = ?
	`, title, content, postID)
	if err != nil {
		tx.Rollback()
		log.Fatal(err)
		return fmt.Errorf("Update post error: %w", err)
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
	ByID
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

	var rows *sql.Rows
	var err error
	switch searchBy {
	case ByEmail:
		rows, err = db.Query("SELECT * FROM users WHERE email = ?", usr.Email)
	case ByID:
		rows, err = db.Query("SELECT * FROM users WHERE id = ?", usr.ID)
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

	if err != nil {
		fmt.Println(err)
		return nil, fmt.Errorf("SearchForUsers Error %v: %v", usr, err)
	}
	defer rows.Close()

	for rows.Next() {
		var usr types.User
		if err := rows.Scan(&usr.ID, &usr.Name, &usr.Email, &usr.Password); err != nil {
			fmt.Println(err)
			return nil, fmt.Errorf("SearchForUsers %v: %v", usr, err)
		}

		// Get the number of followers
		err = db.QueryRow(`
			SELECT COUNT(*) FROM followers WHERE followed_id = ?
		`, usr.ID).Scan(&usr.Followers)
		if err != nil && err != sql.ErrNoRows {
			fmt.Println("Error fetching followers:", err)
		}

		// Get the number of following
		err = db.QueryRow(`
			SELECT COUNT(*) FROM followers WHERE follower_id = ?
		`, usr.ID).Scan(&usr.Following)
		if err != nil && err != sql.ErrNoRows {
			fmt.Println("Error fetching following:", err)
		}

		foundList = append(foundList, usr)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("SearchForUsers %v: %v", usr, err)
	}
	return foundList, nil
}
