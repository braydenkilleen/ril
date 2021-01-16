package main

import (
	"fmt"
	"log"

	"github.com/braydenkilleen/ril/database"
	"github.com/braydenkilleen/ril/model"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func initDatabase() {
	var err error
	database.DBConn, err = gorm.Open("sqlite3", "ril_test.db")
	if err != nil {
		panic("failed to connect database")
	}
	fmt.Println("Connected to database.")
	database.DBConn.AutoMigrate(&model.Bookmark{})
	fmt.Println("Database Migrated")
}

func setupRoutes(app *fiber.App) {
	app.Get("/api/v1/bookmarks", model.GetBookmarks)
	app.Get("/api/v1/bookmarks/:id", model.GetBookmark)
	app.Post("/api/v1/bookmarks", model.NewBookmark)
	app.Delete("/api/v1/bookmarks/:id", model.DeleteBookmark)

	// 404 Handler
	app.Use(func(c *fiber.Ctx) error {
		return c.SendStatus(404)
	})
}

func main() {
	app := fiber.New()
	app.Use(cors.New())

	initDatabase()
	defer database.DBConn.Close()

	setupRoutes(app)

	log.Fatal(app.Listen(":3000"))
}
