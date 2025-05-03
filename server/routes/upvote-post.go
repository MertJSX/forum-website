package routes

import (
	"database/sql"

	"github.com/MertJSX/forum-website/server/database"
	"github.com/gofiber/fiber/v2"
)

func HandleUpvotePost(c *fiber.Ctx, db *sql.DB) error {
	// Get the post ID from the URL parameters
	postID := c.Params("postID")
	if postID == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Post ID is required",
		})
	}

	// Get the user ID from the request context (assuming you have middleware to set this)
	userID := c.Locals("userID").(string)

	// Upvote the post in the database
	currentUpvotes, err := database.UpvotePost(db, userID, postID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to upvote post",
		})
	}

	return c.JSON(fiber.Map{
		"message": "Post upvoted successfully",
		"upvotes": currentUpvotes,
	})
}
