package main

import (
	"context"
	"fmt"
	"log"
	"time"
	"user_service/internal/adapter/repository"
	"user_service/internal/core/domain"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "myuser"
	password = "mypassword"
	dbname   = "user"
)

func main() {
	// dsn := "host=localhost user=myuser password=mypassword dbname=user port=5432 sslmode=disable"

	dsn := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect to database")
	}

	print(db)
	fmt.Print("Database connected!")

	db.AutoMigrate(domain.User{})
	fmt.Println("Database migration completed!")

	userRepository := repository.NewUserRepositoryDB(db)
	_ = userRepository

	ctx := context.Background()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel() // สำคัญ: ต้อง cancel เสมอเมื่อจบการทำงาน

	users, err := userRepository.FindByEmail(ctx, "thann@example.com")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(users)
}
