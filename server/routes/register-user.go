package routes

import (
	"database/sql"

	"github.com/MertJSX/forum-website/server/database"
	"github.com/MertJSX/forum-website/server/types"
	"github.com/gofiber/fiber/v2"
)

func HandleRegisterUser(c *fiber.Ctx, db *sql.DB) error {
	user := new(types.User)

	if err := c.BodyParser(user); err != nil {
		return c.Status(400).JSON(types.ErrorResponse{
			IsError:  true,
			ErrorMsg: "Bad request",
		})
	}

	if user.Name == "" || user.Email == "" || user.Password == "" {
		return c.Status(400).JSON(types.ErrorResponse{
			IsError:  true,
			ErrorMsg: "Missing required fields",
		})
	}

	foundUser, err := database.SearchForUsers(db, *user, database.ByEmail)

	if err != nil {
		return c.Status(500).JSON(types.ErrorResponse{
			IsError:  true,
			ErrorMsg: "Internal server error",
		})
	}

	if foundUser != nil {
		return c.Status(400).JSON(types.ErrorResponse{
			IsError:  true,
			ErrorMsg: "Email already exists",
		})
	}

	foundUser, err = database.SearchForUsers(db, *user, database.ByUsername)

	if err != nil {
		return c.Status(500).JSON(types.ErrorResponse{
			IsError:  true,
			ErrorMsg: "Internal server error",
		})
	}

	if foundUser != nil {
		return c.Status(400).JSON(types.ErrorResponse{
			IsError:  true,
			ErrorMsg: "Username already exists",
		})
	}

	if len(user.Password) < 7 {
		return c.Status(400).JSON(types.ErrorResponse{
			IsError:  true,
			ErrorMsg: "Password too short",
		})
	}

	database.CreateNewUser(db, *user)

	return c.SendString("User has created!")

}
