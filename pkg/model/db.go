package model

import "github.com/jinzhu/gorm"

// Migrate migrates all models.
func Migrate(db *gorm.DB) {
	db.AutoMigrate(&User{})
}
