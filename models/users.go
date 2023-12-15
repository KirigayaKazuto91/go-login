package models

import (
	"log"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `gorm:"unique"`
	Password string
}

func MigrateUsers(db *gorm.DB) error{
	err := db.AutoMigrate(&User{})
	if err != nil {
		log.Fatal("Migration Failed", err)
	}
	return err
}