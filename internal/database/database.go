package database

import (
	"fmt"
	"log"

	"bliss.com/tfcatalogue/internal/config"
	_ "github.com/tursodatabase/libsql-client-go/libsql"
	sqlite "github.com/ytsruh/gorm-libsql"
	"gorm.io/gorm"
)

// Global variable to hold the database connection
var Database *gorm.DB

// Caching the configuration values

// Connect initializes the database connection
func Connect() error {
	var err error
	var (
		cfg        = config.Config()
		dbUsername = cfg.DbUser
		dbUrl      = fmt.Sprintf("%s?authToken=%s", cfg.TursoDbUrl, cfg.TursoDBAuthToken)
	)
	log.Printf("Connecting to database at %s", dbUsername)
	log.Printf("Connecting to database URL at %s", dbUrl)

	Database, err = gorm.Open(sqlite.New(sqlite.Config{
		DSN:        dbUrl,
		DriverName: "libsql",
	}), &gorm.Config{})

	if err != nil {
		log.Printf("Error connecting to the database: %v", config.Config())
		return err
	}

	log.Println("Database connection established successfully")
	return nil
}
