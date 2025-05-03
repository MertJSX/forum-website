package routes

import (
	"database/sql"
	"strconv"

	"github.com/MertJSX/forum-website/server/database"
	"github.com/gofiber/fiber/v2"
)

func HandleGetPostComments(c *fiber.Ctx, db *sql.DB) error {
	postID := c.Params("id")
	if postID == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Post ID is required",
		})
	}

	// Convert postID to integer
	postIDInt, err := strconv.Atoi(postID)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid Post ID",
		})
	}

	// Fetch comments from the database
	comments, err := database.GetCommentsForPost(db, postIDInt)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to fetch comments",
		})
	}

	return c.JSON(comments)
}
