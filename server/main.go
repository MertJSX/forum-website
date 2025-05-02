package main

import (
	"database/sql"
	"log"

	"github.com/MertJSX/forum-website/server/database"
	"github.com/MertJSX/forum-website/server/middleware"
	"github.com/MertJSX/forum-website/server/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	_ "github.com/mattn/go-sqlite3"
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

	var PORT string = ":3000"

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	app.Post("/register-user", func(c *fiber.Ctx) error {
		return routes.HandleRegisterUser(c, db)
	})

	app.Post("/login-user", func(c *fiber.Ctx) error {
		return routes.HandleLoginUser(c, db)
	})

	app.Get("/posts", func(c *fiber.Ctx) error {
		return routes.HandleGetPosts(c, db)
	})

	app.Get("/posts/:id", func(c *fiber.Ctx) error {
		return routes.HandleGetPost(c, db)
	})

	app.Use("/", func(c *fiber.Ctx) error {
		return middleware.CheckAuth(c)
	})

	app.Get("/profile", func(c *fiber.Ctx) error {
		return routes.HandleGetProfile(c, db)
	})

	app.Post("/create-post", func(c *fiber.Ctx) error {
		return routes.HandleCreatePost(c, db)
	})

	app.Post("/create-comment", func(c *fiber.Ctx) error {
		return routes.HandleCommentPost(c, db)
	})

	app.Listen(PORT)
}
