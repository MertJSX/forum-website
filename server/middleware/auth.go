package middleware

import (
	"fmt"

	"github.com/MertJSX/forum-website/server/types"
	"github.com/MertJSX/forum-website/server/utils"
	"github.com/gofiber/fiber/v2"
)

func CheckAuth(c *fiber.Ctx) error {
	var token string = c.Get("Authorization")
	if token == "" {
		return c.Status(401).JSON(types.ErrorResponse{
			IsError:  true,
			ErrorMsg: "Token is missing",
		})
	}

	var userID string
	var username string
	var err error

	if userID, username, err = utils.VerifyToken(token, "test"); err != nil {
		fmt.Println("Error: ", err)
		return c.Status(401).JSON(types.ErrorResponse{
			IsError:  true,
			ErrorMsg: "Invalid token",
		})
	}

	c.Locals("userID", userID)
	c.Locals("username", username)

	return c.Next()
}
