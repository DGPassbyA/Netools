package database

import (
	"log"
	"main/config"
	"main/model"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var SDB *gorm.DB

func SqliteInit() {
	db, err := gorm.Open(sqlite.Open(config.SqliteDBPath), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Init Sqlite successfully.")
	SDB = db
}

//run at first time with sqlite
func SqliteInitSchema() {
	SDB.AutoMigrate(&model.Bookmark{})
}
