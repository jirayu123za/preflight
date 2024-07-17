package main

import (
	"fmt"
	"log"
	"os"
	"preflight/internal/adapters"
	"preflight/internal/core/services"
	"preflight/internal/models"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	dsn := os.Getenv("DATABASE_DSN")
	if err != nil {
		log.Fatal("Error loading .env file")

	}

	db, err := ConnectDB(dsn)
	if err != nil {
		log.Fatal(err)
	}

	app := fiber.New()

	studentRepo := adapters.NewGormStudentRepository(db)
	studentService := services.NewStudentService(studentRepo)
	studentHandler := adapters.NewHttpStudentHandler(studentService)

	app.Put("/addStudent", studentHandler.Add)

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Sever is running")
	})

	port := os.Getenv("PORT")
	app.Listen(":" + port)
	fmt.Println("Server is running on port", port)
}

func ConnectDB(dsn string) (*gorm.DB, error) {
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold: time.Second, // Slow SQL threshold
			LogLevel:      logger.Info, // Log level
			Colorful:      true,        // Enable color
		},
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		panic("Failed to connect to database")
	}
	fmt.Println("Connected to database successfully")

	err = db.AutoMigrate(&models.Student{})
	if err != nil {
		return nil, err
	}
	fmt.Println("Database migration completed!")

	return db, err
}
