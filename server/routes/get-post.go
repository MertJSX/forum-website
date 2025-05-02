package routes

import (
	"database/sql"

	"github.com/MertJSX/forum-website/server/database"
	"github.com/MertJSX/forum-website/server/types"
	"github.com/gofiber/fiber/v2"
)

func HandleGetPost(c *fiber.Ctx, db *sql.DB) error {
	postID := c.Params("id")
	post, err := database.GetPostByID(db, postID)

	if err != nil {
		return c.Status(500).JSON(types.ErrorResponse{
			IsError:  true,
			ErrorMsg: "Internal server error",
		})
	}

	if post == nil {
		return c.Status(404).JSON(types.ErrorResponse{
			IsError:  true,
			ErrorMsg: "Post not found",
		})
	}

	return c.JSON(post)
}
