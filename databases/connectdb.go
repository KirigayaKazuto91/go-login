package databases

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDB()(*gorm.DB, error){
	dsn := "host=localhost user=postgres password=root dbname=projectone port=5432 sslmode=disable"

	conn, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil{
		log.Fatal("Database Connection Failed:", err)
	}
	fmt.Println("Connected to database!")
	return conn, nil
}