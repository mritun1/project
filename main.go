package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/mritun1/project/database"
	"github.com/mritun1/project/users"
)

func setupRoutes(app *fiber.App) {
	app.Get("/api/v1/get_users", users.GetUsers)
	app.Post("/api/v1/insert_users", users.InsertUsers)
	app.Post("/api/v1/update_users", users.UpdateUsers)
	app.Delete("/api/v1/delete_users/:id", users.DeleteUsers)

}

func initDatabase() {
	var err error
	database.DBConn, err = gorm.Open("sqlite3", "users.db")
	if err != nil {
		panic("Failed to connect to database")
	}
	fmt.Println("Database created successfully")

	database.DBConn.AutoMigrate(&users.Users{})
	fmt.Println("Database Migrated")
}

func main() {
	app := fiber.New()

	initDatabase()
	defer database.DBConn.Close()

	setupRoutes(app)

	app.Listen(":4000")
}
