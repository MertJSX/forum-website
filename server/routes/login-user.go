package routes

import (
	"database/sql"
	"fmt"

	"github.com/MertJSX/forum-website/server/database"
	"github.com/MertJSX/forum-website/server/types"
	"github.com/MertJSX/forum-website/server/utils"
	"github.com/gofiber/fiber/v2"
)

func HandleLoginUser(c *fiber.Ctx, db *sql.DB, jwtsecret string) error {
	loginReqBody := new(types.LoginRequest)

	if err := c.BodyParser(loginReqBody); err != nil {
		return c.Status(400).JSON(types.ErrorResponse{
			IsError:  true,
			ErrorMsg: "Bad request",
		})
	}

	user := new(types.User)
	var foundUser []types.User
	var err error
	var possibleErrorMsg string

	if loginReqBody.Email != "" {
		possibleErrorMsg = "Email or password is wrong"
		user.Email = loginReqBody.Email
		user.Password = loginReqBody.Password
		foundUser, err = database.SearchForUsers(db, *user, database.ByEmailAndPassword)
	} else {
		possibleErrorMsg = "Username or password is wrong"
		user.Name = loginReqBody.Name
		user.Password = loginReqBody.Password
		foundUser, err = database.SearchForUsers(db, *user, database.ByUsernameAndPassword)
	}

	if err == nil && foundUser != nil {
		user.Name = foundUser[0].Name
		user.ID = foundUser[0].ID
		var userID string = fmt.Sprintf("%d", *foundUser[0].ID)

		token, err := utils.CreateToken(userID, foundUser[0].Name, jwtsecret)

		if err != nil {
			return c.Status(500).JSON(types.ErrorResponse{
				IsError:  true,
				ErrorMsg: "Unknown error!",
			})
		}

		return c.Status(200).JSON(types.LoginResponse{
			Msg:   "Successfully logged in!",
			Token: token,
		})
	}

	if foundUser == nil {
		return c.Status(200).JSON(types.ErrorResponse{IsError: true, ErrorMsg: possibleErrorMsg})
	}

	return c.Status(200).SendString("Nothing")
}
