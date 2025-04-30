package middleware

import (
	"fmt"

	"github.com/MertJSX/forum-website/server/types"
	"github.com/MertJSX/forum-website/server/utils"
	"github.com/gofiber/fiber/v2"
)

func CheckAuth(c *fiber.Ctx) error {

	var requestWithToken types.RequestWithToken

	fmt.Println("WALLS!")

	if err := c.BodyParser(requestWithToken); err != nil {
		return c.Status(400).JSON(types.ErrorResponse{
			IsError:  true,
			ErrorMsg: "Bad request",
		})
	}

	var username string
	var err error

	if username, err = utils.VerifyToken(requestWithToken.Token, "test"); err != nil {
		return c.Status(400).JSON(types.ErrorResponse{
			IsError:  true,
			ErrorMsg: "Invalid token",
		})
	}

	fmt.Println(username)

	c.Locals("username", username)

	return c.Next()
}
