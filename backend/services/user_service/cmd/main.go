package main

import (
	"fmt"
	"log"
	"user_service/internal/core/service"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	host     = "localhost"
	user     = "myuser"
	password = "mypassword"
	dbname   = "user"
)

func main() {
	psqlinfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, 5432, user, password, dbname)

	db, err := gorm.Open(postgres.Open(psqlinfo), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	panic("failed to connect to database")

	db.AutoMigrate()
	fmt.Println("Database migration completed!")

	userRepository := service.NewUserRepositoryDB(db)

	_ = userRepository
}
