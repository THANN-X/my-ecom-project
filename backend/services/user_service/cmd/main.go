package main

import (
	"fmt"
	"user_service/internal/adapter/handler/httphandler"
	"user_service/internal/adapter/repository"
	"user_service/internal/core/domain"
	"user_service/internal/core/service"

	"github.com/gofiber/fiber/v2"
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

	app := fiber.New()

	dsn := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		DryRun: false,
	})
	if err != nil {
		panic("failed to connect to database")
	}

	print(db)
	fmt.Println("Database connected!")

	db.AutoMigrate(domain.User{})
	fmt.Println("Database migration completed!")

	userRepository := repository.NewUserRepositoryDB(db)
	userService := service.NewUserService(userRepository)
	userHandler := httphandler.NewHttpUserHandler(userService)

	// ctx := context.Background()

	// ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	// defer cancel() // สำคัญ: ต้อง cancel เสมอเมื่อจบการทำงาน

	// newUser := &domain.User{
	// 	FirstName: "Thann",
	// 	LastName:  "Khom",
	// 	Email:     "thann2@example.com",
	// 	Password:  "securepassword",
	// 	Phone:     "123-456-7890",
	// 	Address:   "123 Main St, City, Country",
	// 	Role:      "",
	// }
	// users, err := userRepository.AllUsers(ctx)

	//users, err := userRepository.FindByEmail(ctx, "thann@example.com")

	// user, err := userRepository.FindById(ctx, 1)

	// update := map[string]interface{}{
	// 	"first_name": "UpdatedName",
	// 	"last_name":  "UpdatedLastName",
	// 	"role":       "admin",
	// }

	// err = userRepository.Save(ctx, newUser)
	// err = userRepository.Update(ctx, 1, update)
	// err = userRepository.Delete(ctx, 2)

	// if err != nil {
	// 	log.Fatal(err)
	// }

	// for _, u := range users {
	// 	fmt.Printf("User: %+v\n", u.FirstName)
	// }
	//fmt.Println("User saved successfully:", user)
	// fmt.Println("Updated user:", user)
	// fmt.Println("User deleted successfully")
	app.Post("/register", userHandler.RegisterUser)

	app.Listen(":3000")
}
