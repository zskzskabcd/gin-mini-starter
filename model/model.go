package model

import (
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

// Setup initializes the database instance
func Setup() {
	db, err := gorm.Open(sqlite.Open("data.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	Migration(db)
	DB = db
}

func migrate(db *gorm.DB, models ...interface{}) {
	err := db.AutoMigrate(models...)
	if err != nil {
		panic("failed to migrate database")
		return
	}
}

// Migration migrate the schema
func Migration(db *gorm.DB) {
	// Migrate the schema
	//migrate(db, &ActivationCode{}, &Token{})
}
