package user

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func GetUser(c *fiber.Ctx) error {
	idStr := c.Params("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid user ID"})
	}
	user := User{ID: id, Name: "田中", Age: 30}
	return c.JSON(user)
}

func GetUsers(c *fiber.Ctx) error {
	users := []User{
		{ID: 1, Name: "田中", Age: 30},
		{ID: 2, Name: "佐藤", Age: 25},
	}
	return c.JSON(users)
}

func CreateUser(c *fiber.Ctx) error {
	user := new(User)
	if err := c.BodyParser(user); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Cannot parse JSON"})
	}
	return c.JSON(user)
}
