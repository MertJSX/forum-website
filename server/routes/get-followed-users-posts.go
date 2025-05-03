package routes

import (
	"database/sql"

	"github.com/MertJSX/forum-website/server/database"
	"github.com/gofiber/fiber/v2"
)

func HandleGetFollowedUsersPosts(c *fiber.Ctx, db *sql.DB) error {
	userID := c.Locals("userID").(string)

	posts, err := database.GetPostsByFollowedUsers(db, userID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to fetch followed users' posts",
		})
	}

	return c.JSON(posts)

}
