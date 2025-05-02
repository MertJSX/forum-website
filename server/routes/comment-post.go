package routes

import (
	"database/sql"

	"github.com/MertJSX/forum-website/server/database"
	"github.com/MertJSX/forum-website/server/types"
	"github.com/gofiber/fiber/v2"
)

func HandleCommentPost(c *fiber.Ctx, db *sql.DB) error {
	var comment types.Comment

	comment.UserId = c.Locals("userID").(*int)

	if err := c.BodyParser(&comment); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid request body")
	}

	err := database.CreateNewComment(db, &comment)

	if err != nil {
		return c.Status(520).JSON(types.ErrorResponse{
			IsError:  true,
			ErrorMsg: "Failed to create comment",
		})
	}

	return c.Status(fiber.StatusCreated).SendString("Comment posted successfully")
}
