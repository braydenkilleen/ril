package main

import (
	"log"

	"github.com/braydenkilleen/ril/database"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"honnef.co/go/tools/config"
)

// func initDatabase() {
// 	var err error
// 	database.DB, err = gorm.Open("sqlite3", "bookmarks.db")
// 	if err != nil {
// 		panic("failed to connect database")
// 	}
// 	fmt.Println("Connected to database.")
// 	database.DB.AutoMigrate(&model.Bookmark{})
// 	fmt.Println("Database Migrated")
// }

// func setupRoutes(app *fiber.App) {
// 	//

// 	app.Get("/api/v1/bookmarks", model.GetBookmarks)
// 	app.Get("/api/v1/bookmarks/:id", model.GetBookmark)
// 	app.Post("/api/v1/bookmarks", model.NewBookmark)
// 	app.Delete("/api/v1/bookmarks/:id", model.DeleteBookmark)

// 	// 404 Handler
// 	app.Use(func(c *fiber.Ctx) error {
// 		return c.SendStatus(404)
// 	})
// }

// Setup ...
func Setup() *fiber.App {
	if err := config.Load(); err != nil {
		log.Fatalf("Couldn't load configuration file: %v", err)
	}

	database.Init()
	return server.Start()
}

func main() {
	app := fiber.New()
	app.Use(cors.New())

	database.Connect()
	// initDatabase()
	// defer database.DBConn.Close()

	// setupRoutes(app)
	// log.Fatal(app.Listen(":3000"))
}
