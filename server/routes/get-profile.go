package routes

import (
	"database/sql"
	"fmt"
	"strconv"

	"github.com/MertJSX/forum-website/server/database"
	"github.com/MertJSX/forum-website/server/types"
	"github.com/gofiber/fiber/v2"
)

func HandleGetProfile(c *fiber.Ctx, db *sql.DB) error {

	idParam := c.Params("id")
	var id int

	if idParam != "" {
		parsedID, err := strconv.Atoi(idParam)
		if err != nil {
			return c.Status(400).JSON(types.ErrorResponse{
				IsError:  true,
				ErrorMsg: "Invalid ID format",
			})
		}
		id = parsedID
	}
	var user []types.User
	var err error

	user, err = database.SearchForUsers(db, types.User{ID: &id}, database.ByID)

	if user == nil {
		return c.Status(404).JSON(types.ErrorResponse{
			IsError:  true,
			ErrorMsg: "User not found",
		})
	}

	if err != nil {
		return c.Status(500).JSON(types.ErrorResponse{
			IsError:  true,
			ErrorMsg: "Internal server error",
		})
	}

	isFollowing, err := database.IsUserFollowing(db, c.Locals("userID").(string), fmt.Sprintf("%d", *user[0].ID))

	fmt.Println(err)

	return c.JSON(fiber.Map{
		"isMe":        c.Locals("userID") == fmt.Sprintf("%d", *user[0].ID),
		"isFollowing": isFollowing,
		"user":        user[0],
	})

}

func HandleGetProfileWithUserID(c *fiber.Ctx, db *sql.DB, userID string) error {
	var user []types.User
	var err error
	parsedID, err := strconv.Atoi(userID)
	if err != nil {
		return c.Status(400).JSON(types.ErrorResponse{
			IsError:  true,
			ErrorMsg: "Invalid ID format",
		})
	}

	user, err = database.SearchForUsers(db, types.User{ID: &parsedID}, database.ByID)

	if err != nil {
		return c.Status(500).JSON(types.ErrorResponse{
			IsError:  true,
			ErrorMsg: "Internal server error",
		})
	}

	isFollowing, _ := database.IsUserFollowing(db, c.Locals("userID").(string), fmt.Sprintf("%d", *user[0].ID))

	return c.JSON(fiber.Map{
		"isMe":        c.Locals("userID") == fmt.Sprintf("%d", *user[0].ID),
		"isFollowing": isFollowing,
		"user":        user[0],
	})

}
