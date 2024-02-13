package model

import "gorm.io/gorm"

var db *gorm.DB

func init() {
	var err error
	db, err = gorm.Open("sqlite3", "db/test.db")
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&Task{})
}
