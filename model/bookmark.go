package model

import (
	"github.com/braydenkilleen/ril/database"
	"github.com/gofiber/fiber/v2"
	"github.com/jinzhu/gorm"
)

// Bookmark ...
type Bookmark struct {
	gorm.Model
	Title string `gorm:"not null" json:"title"`
	URL   string `gorm:"unique_index;not null" json:"url"`
}

// GetBookmarks ...
func GetBookmarks(c *fiber.Ctx) error {
	db := database.DBConn
	var bookmarks []Bookmark
	db.Find(&bookmarks)
	return c.JSON(bookmarks)
}

// GetBookmark ...
func GetBookmark(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.DBConn
	var bookmark Bookmark
	db.Find(&bookmark, id)
	return c.JSON(bookmark)
}

// NewBookmark ...
func NewBookmark(c *fiber.Ctx) error {
	db := database.DBConn
	bookmark := new(Bookmark)
	if err := c.BodyParser(bookmark); err != nil {
		return c.Status(503).JSON(err)
	}
	db.Create(&bookmark)
	return c.JSON(bookmark)
}

// DeleteBookmark ...
func DeleteBookmark(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.DBConn

	var bookmark Bookmark
	db.First(&bookmark, id)
	if bookmark.Title == "" {
		return c.Status(500).SendString("No Bookmark Found with ID")
	}
	db.Delete(&bookmark)
	return c.SendString("Book markSuccessfully deleted")
}
