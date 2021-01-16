package database

import (
	"github.com/jinzhu/gorm"
)

// DB is the database connection
var (
	DBConn *gorm.DB
)

// // Connect established db connection & applies migrations
// func Connect() {
// 	var err error

// 	DB, err =

// 	db, err := gorm.Open(sqlite.Open(config.DB), &gorm.Config{
// 		NowFunc: func() time.Time { return time.Now().Local() },
// 		Logger:  logger.Default.LogMode(logger.Info),
// 	})
// 	if err != nil {
// 		fmt.Println("[DATABASE] - CONNECTION ERROR")
// 		panic(err)

// 	}

// 	DB = db
// 	fmt.Println("[DATABASE] - CONNECTED")
// }

// // Migrate peforms migrations on all tables
// func Migrate(tables ...interface{}) error {
// 	return DB.AutoMigrate(tables...)
// }
