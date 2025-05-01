package routes

import (
	"database/sql"

	"github.com/MertJSX/forum-website/server/database"
	"github.com/MertJSX/forum-website/server/types"
	"github.com/gofiber/fiber/v2"
)

func HandleGetProfile(c *fiber.Ctx, db *sql.DB) error {

	user, err := database.SearchForUsers(db, types.User{Name: c.Locals("username").(string)}, database.ByUsername)

	if err != nil {
		return c.Status(500).JSON(types.ErrorResponse{
			IsError:  true,
			ErrorMsg: "Internal server error",
		})
	}

	return c.JSON(user[0])

}
