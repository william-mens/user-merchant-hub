package handler

import (
	"errors"
	"log"

	"bliss.com/tfcatalogue/api/services"
	"bliss.com/tfcatalogue/internal/helpers"
	"github.com/go-sql-driver/mysql"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func GetMerchants(c *fiber.Ctx) error {

	merch := &helpers.Merchant{
		MerchantId:  c.Query("merchantId"),
		CompanyName: c.Query("companyName"),
		Limit:       c.Query("limit"),
		Tags:        c.Query("tags"),
		Alias:       c.Query("alias"),
	}

	errs := helpers.ValidateData(merch)
	if len(errs) > 0 {
		return helpers.BadRequestResponse(
			c,
			errs,
		)
	}

	request := c.Queries()

	merchants, err := services.GetAll(request)
	if err != nil {

		log.Printf("An error occurred retrieving merchants: %v", err)
		return helpers.InternalServerErrorResponse(
			c,
			nil,
		)
	}

	if len(merchants.Result) == 0 {
		return helpers.NotFoundResponse(c, nil)
	}

	return helpers.SuccessResponse(c, "Success", merchants)

}

func SetupMerchant(c *fiber.Ctx) error {

	merch := new(helpers.SetupMerchant)
	if err := c.BodyParser(merch); err != nil {
		return helpers.BadRequestResponse(c, []string{"Invalid request payload"})
	}

	errs := helpers.ValidateData(merch)

	if errs != nil {
		return helpers.BadRequestResponse(
			c,
			errs,
		)
	}

	merchants, err := services.Store(merch)
	if err != nil {
		if mysqlErr, ok := err.(*mysql.MySQLError); ok && mysqlErr.Number == 1062 {

			log.Printf("debuging exception:%v", mysqlErr.SQLState)
			return helpers.BadRequestResponse(c, []string{"Record already exists, kindly check your parameters"})
		}

		log.Printf("An error occurred creating merchants: %v", err)
		return helpers.InternalServerErrorResponse(
			c,
			nil,
		)
	}

	return helpers.SuccessResponse(c, "Success", merchants)

}

func UpdateMerchant(c *fiber.Ctx) error {

	merch := new(helpers.SetupMerchant)
	if err := c.BodyParser(merch); err != nil {
		return helpers.BadRequestResponse(c, []string{"Invalid request payload"})
	}

	merchant, err := services.Update(merch)
	if err != nil {
		if mysqlErr, ok := err.(*mysql.MySQLError); ok && mysqlErr.Number == 1062 {
			log.Printf("debuging exception:%v", mysqlErr.SQLState)
			return helpers.BadRequestResponse(c, []string{"Record already exists, kindly check your parameters"})
		}
		log.Printf("An error occurred updating merchants: %v", err)

		return helpers.InternalServerErrorResponse(
			c,
			nil,
		)
	}
	return helpers.SuccessResponse(c, "Success", merchant)

}

func DeleteMerchant(c *fiber.Ctx) error {
	merchantId := c.Params("merchantId")
	if len(merchantId) == 0 {
		return helpers.BadRequestResponse(c, []string{"merchantId is required"})
	}

	merchant, err := services.Destroy(merchantId)
	log.Printf("An error occurred deleting merchant: %v", err)

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return helpers.BadRequestResponse(c, []string{"merchantId not found"})
		}

		return helpers.InternalServerErrorResponse(
			c,
			nil,
		)
	}

	return helpers.SuccessResponse(c, "Success", merchant)

}
