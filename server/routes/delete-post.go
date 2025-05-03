package routes

import (
	"database/sql"

	"github.com/MertJSX/forum-website/server/database"
	"github.com/gofiber/fiber/v2"
)

func HandleDeletePost(c *fiber.Ctx, db *sql.DB) error {
	postID := c.Params("id")
	if postID == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Post ID is required",
		})
	}

	userID := c.Locals("userID").(string)

	// Check if the user is the author of the post
	isAuthor, err := database.IsPostAuthor(db, userID, postID)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to verify post author",
		})
	}

	if !isAuthor {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"error": "You are not authorized to delete this post",
		})
	}

	err = database.DeletePost(db, postID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to delete post",
		})
	}

	return c.JSON(fiber.Map{
		"message": "Post deleted successfully",
	})
}
