package routes

import (
	"database/sql"

	"github.com/MertJSX/forum-website/server/database"
	"github.com/gofiber/fiber/v2"
)

func HandleEditPost(c *fiber.Ctx, db *sql.DB) error {
	postID := c.Params("id")
	if postID == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Post ID is required",
		})
	}

	userID := c.Locals("userID").(string)
	if userID == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Unauthorized",
		})
	}

	isAuthor, err := database.IsPostAuthor(db, postID, userID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to verify post author",
		})
	}
	if !isAuthor {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"error": "You are not the author of this post",
		})
	}

	var postData struct {
		Title   string `json:"title"`
		Content string `json:"content"`
	}

	if err := c.BodyParser(&postData); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	err = database.UpdatePost(db, postID, postData.Title, postData.Content)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to update post",
		})
	}

	return c.JSON(fiber.Map{
		"message": "Post updated successfully",
	})
}
