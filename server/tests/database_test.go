package tests

import (
	"database/sql"
	"fmt"
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

	database.CreatePostsTable(db)
	database.CreateUsersTable(db)
	database.CreateCommentsTable(db)
	database.CreateUpvotesTable(db)
	database.CreateFollowersTable(db)

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
		database.CreatePostsTable(db)
	})

	t.Run("Create users table", func(t *testing.T) {
		database.CreateUsersTable(db)
	})

	t.Run("Create followers table", func(t *testing.T) {
		database.CreateFollowersTable(db)
	})

	t.Run("Create comments table", func(t *testing.T) {
		database.CreateCommentsTable(db)
	})

	t.Run("Create upvotes table", func(t *testing.T) {
		database.CreateUpvotesTable(db)
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

	t.Run("Create new post with user", func(t *testing.T) {
		testPost := types.Post{
			Title:   "Test Post",
			Content: "This is a test post.",
			UserId:  1,
		}

		_, err := database.CreateNewPost(db, &testPost)
		if err != nil {
			t.Fatalf("Failed to create new post: %v", err)
		}

		createdPost, err := database.GetPostByID(db, fmt.Sprintf("%d", 1))
		if err != nil {
			t.Fatalf("Failed to retrieve created post: %v", err)
		}

		if createdPost.Title != testPost.Title || createdPost.Content != testPost.Content || createdPost.UserId != testPost.UserId {
			t.Fatalf("Created post does not match expected values")
		}
	})

	t.Run("Create new comment", func(t *testing.T) {
		testComment := types.Comment{
			Comment: "This is a test comment.",
			UserId:  "1",
			PostId:  "1",
		}

		err := database.CreateNewComment(db, &testComment)
		if err != nil {
			t.Fatalf("Failed to create new comment: %v", err)
		}
	})

	t.Run("Get post comments", func(t *testing.T) {
		comments, err := database.GetCommentsForPost(db, 1)
		if err != nil {
			t.Fatalf("Failed to retrieve post comments: %v", err)
		}

		if len(comments) == 0 {
			t.Fatalf("No comments found for post")
		}
	})

	t.Run("Edit post", func(t *testing.T) {
		testPost := types.Post{
			Title:   "Original Title",
			Content: "Original content.",
			UserId:  1,
		}

		postID, err := database.CreateNewPost(db, &testPost)
		if err != nil {
			t.Fatalf("Failed to create test post: %v", err)
		}

		newTitle := "Updated Title"
		newContent := "Updated content."
		err = database.UpdatePost(db, fmt.Sprintf("%d", postID), newTitle, newContent)
		if err != nil {
			t.Fatalf("Failed to update post: %v", err)
		}

		updatedPost, err := database.GetPostByID(db, fmt.Sprintf("%d", postID))
		if err != nil {
			t.Fatalf("Failed to retrieve updated post: %v", err)
		}

		if updatedPost.Title != newTitle || updatedPost.Content != newContent {
			t.Fatalf("Updated post does not match expected values")
		}
	})
}
