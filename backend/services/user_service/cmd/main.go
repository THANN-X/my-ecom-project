package main

import (
	"fmt"
	"os"
	"user_service/internal/adapter/handler/httphandler"
	"user_service/internal/adapter/repository"
	"user_service/internal/core/domain"
	"user_service/internal/core/service"

	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	// Load database configuration from environment variables
	dbHost := os.Getenv("DB_HOST")
	if dbHost == "" {
		dbHost = "localhost"
	}

	dbPort := os.Getenv("DB_PORT")
	if dbPort == "" {
		dbPort = "5432"
	}

	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	fmt.Println(dbHost)

	// dsn := fmt.Sprintf("host=%s port=%d user=%s "+
	// 	"password=%s dbname=%s sslmode=disable",
	// 	dbHost, port, user, password, dbname)

	// Connect to the database
	dsn := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		dbHost, dbPort, dbUser, dbPassword, dbName)

	fmt.Println(dsn)

	// Open database connection
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

	// Initialize repository, service, and handler
	userRepository := repository.NewUserRepositoryDB(db)
	userService := service.NewUserService(userRepository)
	userHandler := httphandler.NewUserHandler(userService)

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

	// Initialize Fiber app and routes
	app := fiber.New()

	// Define routes
	app.Post("/users/register", userHandler.RegisterUser)
	app.Post("/users/update/:id", userHandler.UpdateUserProfile)
	app.Post("/users/chgpass/:id", userHandler.ChangePassword)
	app.Post("/users/login", userHandler.LoginUser)
	app.Get("/users/:id", userHandler.GetUserProfile)

	// Start the server
	app.Listen(":3001")
}
