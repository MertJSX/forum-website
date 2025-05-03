package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/MertJSX/forum-website/server/database"
	"github.com/MertJSX/forum-website/server/middleware"
	"github.com/MertJSX/forum-website/server/routes"
	"github.com/MertJSX/forum-website/server/types"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	_ "github.com/mattn/go-sqlite3"
	"gopkg.in/yaml.v3"
)

func main() {
	app := fiber.New()
	app.Use(cors.New())
	db, err := sql.Open("sqlite3", "./database/database.db")
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	database.CreateUsersTable(db)
	database.CreatePostsTable(db)
	database.CreateCommentsTable(db)
	database.CreateUpvotesTable(db)
	database.CreateFollowersTable(db)

	fileData, err := os.ReadFile("./config.yml")
	if err != nil {
		log.Fatal("Error reading config file:", err)
	}
	var config types.ConfigFile
	err = yaml.Unmarshal(fileData, &config)
	if err != nil {
		log.Fatalf("YAML parse error: %v", err)
	}
	var PORT string = fmt.Sprintf(":%d", config.Port)

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	app.Post("/register-user", func(c *fiber.Ctx) error {
		return routes.HandleRegisterUser(c, db)
	})

	app.Post("/login-user", func(c *fiber.Ctx) error {
		return routes.HandleLoginUser(c, db, config.JWTSecret)
	})

	app.Get("/posts", func(c *fiber.Ctx) error {
		return routes.HandleGetPosts(c, db)
	})

	app.Get("/posts/:id", func(c *fiber.Ctx) error {
		return routes.HandleGetPost(c, db)
	})

	app.Get("/posts/:id/comments", func(c *fiber.Ctx) error {
		return routes.HandleGetPostComments(c, db)
	})

	app.Get("/userposts/:id", func(c *fiber.Ctx) error {
		return routes.HandleGetUserPosts(c, db)
	})

	app.Use("/", func(c *fiber.Ctx) error {
		return middleware.CheckAuth(c, config.JWTSecret)
	})

	app.Get("/followed-users-posts", func(c *fiber.Ctx) error {
		return routes.HandleGetFollowedUsersPosts(c, db)
	})

	app.Get("/profile/:id", func(c *fiber.Ctx) error {
		return routes.HandleGetProfile(c, db)
	})

	app.Get("/profile", func(c *fiber.Ctx) error {
		return routes.HandleGetProfileWithUserID(c, db, c.Locals("userID").(string))
	})

	app.Get("/upvote/:postID", func(c *fiber.Ctx) error {
		return routes.HandleUpvotePost(c, db)
	})

	app.Get("follow/:id", func(c *fiber.Ctx) error {
		return routes.HandleFollowUser(c, db)
	})

	app.Post("/create-post", func(c *fiber.Ctx) error {
		return routes.HandleCreatePost(c, db)
	})

	app.Post("/create-comment", func(c *fiber.Ctx) error {
		return routes.HandleCommentPost(c, db)
	})

	app.Delete("/posts/:id", func(c *fiber.Ctx) error {
		return routes.HandleDeletePost(c, db)
	})

	app.Put("/posts/:id", func(c *fiber.Ctx) error {
		return routes.HandleEditPost(c, db)
	})

	app.Listen(PORT)
}
