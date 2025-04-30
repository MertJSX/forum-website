package routes

import (
	"database/sql"
	"fmt"

	"github.com/MertJSX/forum-website/server/types"
	"github.com/gofiber/fiber/v2"
)

func HandleCreateForum(c *fiber.Ctx, db *sql.DB) error {
	request := new(types.ForumRequest)

	if err := c.BodyParser(request); err != nil {
		return c.Status(400).JSON(types.ErrorResponse{
			IsError:  true,
			ErrorMsg: "Bad request",
		})
	}

	fmt.Println(request)

	if request.Token == "" {
		return c.Status(400).JSON(types.ErrorResponse{
			IsError:  true,
			ErrorMsg: "Missing required fields",
		})
	}

	// database.CreateNewUser();

	return c.SendString("User has created!")

}
