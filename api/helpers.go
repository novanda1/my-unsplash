package api

import "github.com/gofiber/fiber/v2"

func ResponseError(c *fiber.Ctx, status int, msg string, data interface{}) error {
	var resp Response
	resp.Message = msg
	resp.Status = "error"
	resp.Data = data

	return c.Status(status).JSON(resp)
}

func ResponseSuccess(c *fiber.Ctx, status int, msg string, data interface{}) error {
	var resp Response
	resp.Message = msg
	resp.Status = "success"
	resp.Data = data

	return c.Status(status).JSON(resp)
}

func BadRequestError(c *fiber.Ctx, msg string, data interface{}) error {
	return ResponseError(c, fiber.StatusBadRequest, msg, data)
}

func InternalServerError(c *fiber.Ctx, msg string, data interface{}) error {
	return ResponseError(c, fiber.StatusInternalServerError, msg, data)
}

func SendOK(c *fiber.Ctx, msg string, data interface{}) error {
	return ResponseSuccess(c, fiber.StatusOK, msg, data)
}

func SendCreated(c *fiber.Ctx, msg string, data interface{}) error {
	return ResponseSuccess(c, fiber.StatusCreated, msg, data)
}
