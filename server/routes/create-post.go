package routes

import (
	"database/sql"
	"fmt"
	"strconv"

	"github.com/MertJSX/forum-website/server/database"
	"github.com/MertJSX/forum-website/server/types"
	"github.com/gofiber/fiber/v2"
)

func HandleCreatePost(c *fiber.Ctx, db *sql.DB) error {
	newPost := new(types.Post)

	if err := c.BodyParser(newPost); err != nil {
		return c.Status(400).JSON(types.ErrorResponse{
			IsError:  true,
			ErrorMsg: "Bad request",
		})
	}

	var err error
	newPost.UserId, err = strconv.Atoi(c.Locals("userID").(string))
	if err != nil {
		return c.Status(400).JSON(types.ErrorResponse{
			IsError:  true,
			ErrorMsg: "Bad request",
		})
	}

	if err := database.CreateNewPost(db, newPost); err != nil {
		fmt.Println("Error creating post:", err)
		return c.Status(500).JSON(types.ErrorResponse{
			IsError:  true,
			ErrorMsg: "Internal server error",
		})
	}

	return c.SendString("Post has created!")
}
