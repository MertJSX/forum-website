package routes

import (
	"database/sql"

	"github.com/MertJSX/forum-website/server/database"
	"github.com/MertJSX/forum-website/server/types"
	"github.com/gofiber/fiber/v2"
)

func HandleGetUserPosts(c *fiber.Ctx, db *sql.DB) error {
	idParam := c.Params("id")
	var id string

	if idParam != "" {
		id = idParam
	} else {
		id = c.Locals("userID").(string)
	}
	var posts []types.Post
	var err error

	posts, err = database.GetPostsByUserID(db, id)

	if posts == nil {
		return c.Status(404).JSON(types.ErrorResponse{
			IsError:  true,
			ErrorMsg: "User don't have any posts",
		})
	}

	if err != nil {
		return c.Status(500).JSON(types.ErrorResponse{
			IsError:  true,
			ErrorMsg: "Internal server error",
		})
	}

	return c.JSON(posts)

}
