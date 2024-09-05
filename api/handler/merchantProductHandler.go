package handler

import (
	"log"

	"bliss.com/tfcatalogue/api/services"
	"bliss.com/tfcatalogue/internal/helpers"
	"github.com/gofiber/fiber/v2"
)

func GetMerchantProducts(c *fiber.Ctx) error {

	merchProductRequest := &helpers.GetMerchantProducts{
		MerchantId:        c.Query("merchantId"),
		MerchantProductId: c.Query("merchantProductId"),
	}

	errs := helpers.ValidateData(merchProductRequest)
	if len(errs) > 0 {
		return helpers.BadRequestResponse(
			c,
			errs,
		)
	}
	requests := c.Queries()

	merchants, err := services.GetMerhantProducts(requests)
	if err != nil {

		log.Printf("An error occurred retrieving merchant products: %v", err)
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
