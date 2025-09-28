package auth

import (
	"github.com/ADMex1/goweb/database"
	"github.com/ADMex1/goweb/model"
	jwt "github.com/ADMex1/goweb/utils"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

var DB *gorm.DB

func init() {
	var err error
	DB, err = database.Connect()
	if err != nil {
		panic("Failed to connect to database: " + err.Error())
	}
}

func UserRegister(c *fiber.Ctx) error {
	var input struct {
		Name     string `json:"name"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	if err := c.BodyParser(&input); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"{-}Error": "Invalid Input",
		})
	}
	hash, _ := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)

	user := model.User{
		Name:     input.Name,
		Email:    input.Email,
		Password: string(hash),
		Role:     "user",
	}
	if err := DB.Create(&user).Error; err != nil {
		return c.Status(400).JSON(fiber.Map{
			"{-}Error": "Unable to Create User",
		})
	}
	token, err := jwt.JWT(int(user.InternalID), user.Name, user.Role, user.Email, "web")
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to generate token"})
	}
	return c.Status(201).JSON(fiber.Map{
		"user": fiber.Map{
			"id":    user.InternalID,
			"name":  user.Name,
			"email": user.Email,
			"role":  user.Role,
		},
		"token": token,
	})

}
func UserLogin(c *fiber.Ctx) error {
	var input struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	if err := c.BodyParser(&input); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"{-}Error": "Invalid Input",
		})
	}
	var user model.User
	if err := DB.Where("email= ?", input.Email).First(&user).Error; err != nil {
		return c.Status(401).JSON(fiber.Map{
			"{-}Error": "Invalid Email",
		})
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)); err != nil {
		return c.Status(401).JSON(fiber.Map{
			"{-}Error": "Invalid password",
		})
	}
	token, err := jwt.JWT(int(user.InternalID), user.Name, user.Role, user.Email, "web")
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to generate token"})
	}
	return c.Status(201).JSON(fiber.Map{
		"user": fiber.Map{
			"id":   user.InternalID,
			"name": user.Name,
		},
		"token": token,
	})
}
