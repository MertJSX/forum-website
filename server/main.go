package main

import (
	"database/sql"
	"log"

	"github.com/MertJSX/forum-website/server/database"
	"github.com/MertJSX/forum-website/server/routes"
	"github.com/gofiber/fiber/v2"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	app := fiber.New()
	db, err := sql.Open("sqlite3", "./database/database.db")
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	database.CreateUserTable(db)

	var PORT string = ":3000"

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	app.Post("/register-user", func(c *fiber.Ctx) error {
		return routes.HandleRegisterUser(c)
	})

	app.Listen(PORT)
}
