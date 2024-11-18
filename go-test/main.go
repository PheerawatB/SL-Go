package main

import (
	"fmt"
	"gomod-test/models"

	// "database/sql"
	// "log"
	"os"

	"github.com/gofiber/fiber/v2"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	jwtware "github.com/gofiber/jwt/v2"
	// "github.com/golang-migrate/migrate/v4"
	// "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
)

const (
	host     = "localhost"  // or the Docker service name if running in another container
	port     = 5432         // default PostgreSQL port
	username = "postgres"   // as defined in docker-compose.yml
	password = "mypassword" // as defined in docker-compose.yml
	dbname   = "mydatabase" // as defined in docker-compose.yml
)

// @title API
// @description This is a sample server for a book API.
// @version 1.0
// @host localhost:8080
// @BasePath /
// @schemes http
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	// Configure your PostgreSQL database details here
	dsn := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, username, password, dbname)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect to database")
	}
	db.AutoMigrate(&models.Mobile{})
	fmt.Println("Database migration completed!")
	// dsn := "postgres://myuser:mypassword@localhost:5432/mydatabase?sslmode=disable"

	// db, err := sql.Open("postgres", dsn)
	// if err != nil {
	// 	log.Fatalf("Failed to connect to database: %v", err)
	// }
	// defer db.Close()

	// driver, err := postgres.WithInstance(db, &postgres.Config{})
	// if err != nil {
	// 	log.Fatalf("Failed to create Postgres driver: %v", err)
	// }

	// m, err := migrate.NewWithDatabaseInstance(
	// 	"file://migrations",
	// 	"postgres",
	// 	driver,
	// )
	// if err != nil {
	// 	log.Fatalf("Failed to create migration instance: %v", err)
	// }

	// if err := m.Up(); err != nil && err != migrate.ErrNoChange {
	// 	log.Fatalf("Migration failed: %v", err)
	// }

	// log.Println("Migration completed successfully!")
	app := fiber.New()

	// app.Get("/swagger/*", swagger.HandlerDefault) //default

	// JWT Secret Key
	secretKey := os.Getenv("SECRET_KEY")

	app.Post("/login", login(secretKey))

	// JWT Middleware
	app.Use(jwtware.New(jwtware.Config{
		SigningKey: []byte(secretKey),
	}))

	// Middleware to extract user data from JWT
	app.Use(extractUserFromJWT)

	// Group routes under /book
	bookGroup := app.Group("/book")

	// Apply the isAdmin middleware only to the /book routes
	bookGroup.Use(isAdmin)

	// Now, only authenticated admins can access these routes
	bookGroup.Get("/", getBooks)
	bookGroup.Get("/:id", getBook)
	bookGroup.Post("/", createBook)
	bookGroup.Put("/:id", updateBook)
	bookGroup.Delete("/:id", deleteBook)

	app.Get("/api/config", getConfig)
	app.Listen(":8080")
}

func getConfig(c *fiber.Ctx) error {
	// Example: Return a configuration value from environment variable
	secretKey := os.Getenv("SECRET_KEY")
	if secretKey == "" {
		secretKey = "defaultSecret" // Default value if not specified
	}

	return c.JSON(fiber.Map{
		"secret_key": secretKey,
	})
}
