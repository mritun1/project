package users

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jinzhu/gorm"
	"github.com/mritun1/project/database"
)

type Users struct {
	gorm.Model
	Name        string `json:"name"`
	Dob         int    `json:"dob"`
	Address     string `json:"address"`
	Description string `json:"description"`
}

//GET ALL USERS FROM DATABASE
func GetUsers(c *fiber.Ctx) error {
	db := database.DBConn
	var users []Users
	db.Find(&users)
	return c.Status(200).JSON(users)
}

//INSERT USER DATA INTO THE DATABASE
func InsertUsers(c *fiber.Ctx) error {
	db := database.DBConn

	type UpdateUsr struct {
		ID          string `json:"id"`
		Name        string `json:"name"`
		Dob         int    `json:"dob"`
		Address     string `json:"address"`
		Description string `json:"description"`
	}
	var updateData UpdateUsr
	if err := c.BodyParser(&updateData); err != nil {
		return c.Status(500).JSON(err.Error())
	}

	var user Users
	user.Name = updateData.Name
	user.Dob = updateData.Dob
	user.Address = updateData.Address
	user.Description = updateData.Description
	db.Create(&user)
	return c.Status(200).JSON(user)
}

//UPDATE USER DATA INTO THE DATABASE
func UpdateUsers(c *fiber.Ctx) error {
	db := database.DBConn

	type UpdateUsr struct {
		ID          string `json:"id"`
		Name        string `json:"name"`
		Dob         string `json:"dob"`
		Address     string `json:"address"`
		Description string `json:"description"`
	}
	var updateData UpdateUsr
	if err := c.BodyParser(&updateData); err != nil {
		return c.Status(500).JSON(err.Error())
	}

	var user Users
	db.Model(&user).Where("ID = ?", updateData.ID).Update("name", updateData.Name)
	db.Model(&user).Where("ID = ?", updateData.ID).Update("dob", updateData.Dob)
	db.Model(&user).Where("ID = ?", updateData.ID).Update("address", updateData.Address)
	db.Model(&user).Where("ID = ?", updateData.ID).Update("description", updateData.Description)

	return c.Status(200).JSON("Update success")
}

//DELETE USER DATA FROM DATABASE
func DeleteUsers(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.DBConn
	var user Users
	db.First(&user, id)
	if user.Name == "" {
		return c.Status(500).JSON("No user Found")

	}
	db.Delete(&user)
	return c.Status(200).JSON(user)
}
