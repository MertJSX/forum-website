package routes

import (
	"database/sql"

	"github.com/MertJSX/forum-website/server/database"
	"github.com/MertJSX/forum-website/server/types"
	"github.com/gofiber/fiber/v2"
)

func HandleCreatePost(c *fiber.Ctx, db *sql.DB) error {
	request := new(types.Post)

	if err := c.BodyParser(request); err != nil {
		return c.Status(400).JSON(types.ErrorResponse{
			IsError:  true,
			ErrorMsg: "Bad request",
		})
	}

	if err := database.CreateNewPost(db, request); err != nil {
		return c.Status(500).JSON(types.ErrorResponse{
			IsError:  true,
			ErrorMsg: "Internal server error",
		})
	}

	return c.SendString("Post has created!")
}
