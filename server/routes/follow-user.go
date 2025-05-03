package routes

import (
	"database/sql"

	"github.com/MertJSX/forum-website/server/database"
	"github.com/gofiber/fiber/v2"
)

func HandleFollowUser(c *fiber.Ctx, db *sql.DB) error {
	userID := c.Locals("userID").(string)
	followUserID := c.Params("id")

	if userID == followUserID {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "You cannot follow yourself",
		})
	}

	followers, err := database.FollowUser(db, userID, followUserID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to follow user",
		})
	}

	isFollowing, _ := database.IsUserFollowing(db, userID, followUserID)

	return c.JSON(fiber.Map{
		"message":     "Successfully followed user",
		"followers":   followers,
		"isFollowing": isFollowing,
	})
}
