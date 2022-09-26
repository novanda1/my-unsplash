package api

import (
	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/novanda1/my-unsplash/models"
)

type ErrorResponse struct {
	FailedField string
	Tag         string
	Value       string
}

var validate = validator.New()

func ValidateInsertImageParams(insertImageParams models.InsertImageDTO) []*ErrorResponse {
	var errors []*ErrorResponse
	err := validate.Struct(insertImageParams)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var element ErrorResponse
			element.FailedField = err.StructNamespace()
			element.Tag = err.Tag()
			element.Value = err.Param()
			errors = append(errors, &element)
		}
	}
	return errors
}

func (a *API) AddIMage(c *fiber.Ctx) error {
	c.Accepts("application/json")

	var resp Response

	params := new(models.InsertImageDTO)
	if err := c.BodyParser(params); err != nil {
		return err
	}

	errors := ValidateInsertImageParams(*params)
	if errors != nil {
		resp.Status = "error"
		resp.Data = errors
		resp.Message = "body error"

		return c.Status(fiber.StatusBadRequest).JSON(resp)
	}

	result, err := models.SaveImage(a.db, params)
	if err != nil {
		resp.Status = "error"
		resp.Data = result
		resp.Message = fmt.Sprintf("save image error: %s", err.Error())

		return c.JSON(resp)
	}

	resp.Status = "success"
	resp.Data = result
	resp.Message = "Succesfully added new image!"

	return c.JSON(resp)
}

func (a *API) GetImagesWithPagination(c *fiber.Ctx) error {
	var resp Response

	params := &models.GetImageDTO{Limit: 25, Cursor: ""}
	if err := c.QueryParser(params); err != nil {
		return err
	}

	result, err := models.GetImages(a.db, params)
	if err != nil {
		resp.Status = "error"
		resp.Data = result
		resp.Message = fmt.Sprintf("get image error: %s", err.Error())

		return c.JSON(resp)
	}

	resp.Status = "success"
	resp.Data = result

	return c.JSON(resp)
}
