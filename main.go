package main

import (
	"fmt"
	"log"
	"os"

	"bliss.com/tfcatalogue/api/handler"
	"bliss.com/tfcatalogue/internal/database"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {

	app := fiber.New()

	err := godotenv.Load()
	environment := os.Getenv("ENV")
	if environment == "development" {
		if err != nil {
			log.Fatal("Error loading .env file", err)
		}
	}

	fmt.Println(os.Getenv("DB_HOST"))
	database.Connect()

	app.Get("/merchants", handler.GetMerchants)
	app.Post("/merchants", handler.SetupMerchant)
	app.Post("/auth/register", handler.BeginRegistration)
	app.Put("/merchants", handler.UpdateMerchant)
	app.Delete("/merchants/:merchantId", handler.DeleteMerchant)
	app.Get("/merchant/products", handler.GetMerchantProducts)

	log.Fatal(app.Listen(":80"))
}
