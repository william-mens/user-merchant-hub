package helpers

import (
	"github.com/gofiber/fiber/v2"
)

type ResponseWrapper struct {
	ResponseCode    int         `json:"responseCode"`
	ResponseMessage string      `json:"responseMessage"`
	Data            interface{} `json:"data"`
}

func sendResponse(c *fiber.Ctx, statusCode int, message string, data interface{}) error {

	if data == nil {
		data = map[string]interface{}{}
	}

	responseStruct := ResponseWrapper{
		ResponseCode:    statusCode,
		ResponseMessage: message,
		Data:            data,
	}
	res := c.Status(statusCode).JSON(responseStruct)
	return res
}

func SuccessResponse(c *fiber.Ctx, message string, data interface{}) error {
	err := sendResponse(c, fiber.StatusOK, message, data)
	return err

}

func InternalServerErrorResponse(c *fiber.Ctx, data interface{}) error {
	err := sendResponse(c, fiber.StatusInternalServerError, "An error occurred,Please try again later", data)
	return err
}

func BadRequestResponse(c *fiber.Ctx, data interface{}) error {
	err := sendResponse(c, fiber.StatusBadRequest, "validation failed kindly check your parameters", data)
	return err
}

func NotFoundResponse(c *fiber.Ctx, data interface{}) error {

	err := sendResponse(c, fiber.StatusNotFound, "Record not found", data)
	return err
}

func CustomResponse(c *fiber.Ctx, statusCode int, message string, data interface{}) error {
	err := sendResponse(c, statusCode, message, data)
	return err
}
