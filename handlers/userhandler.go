package handlers

import "github.com/gofiber/fiber/v2"

func Login(ctx *fiber.Ctx) error {

	type User struct{
		Username string `json:"username"`
		Password string `json:"password"`

	}
	var user User
	err:=ctx.BodyParser(&user)
	if err!=nil {
		return fiber.NewError(fiber.StatusBadRequest,"Error Parsing the body")
	}
	return nil

}