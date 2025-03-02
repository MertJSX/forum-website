package routes

import (
	"fmt"

	"github.com/MertJSX/forum-website/server/types"
	"github.com/gofiber/fiber/v2"
)

func HandleRegisterUser(c *fiber.Ctx) error {
	reqBody := new(types.User)

	if err := c.BodyParser(reqBody); err != nil {
		return err
	}

	fmt.Println(reqBody.Name)
	return c.JSON(reqBody)
}
