package api

import (
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

func (a *API) AddImageHandler(c *fiber.Ctx) error {
	c.Accepts("application/json")

	params := new(models.InsertImageDTO)
	if err := c.BodyParser(params); err != nil {
		return BadRequestError(c, err.Error(), nil)
	}

	errors := ValidateInsertImageParams(*params)
	if errors != nil {
		return BadRequestError(c, "body error", errors)
	}

	result, err := models.SaveImage(a.db, params)
	if err != nil {
		return InternalServerError(c, err.Error(), result)
	}

	return SendCreated(c, "Succesfully added new image!", result)
}

func (a *API) GetImagesHandler(c *fiber.Ctx) error {
	params := &models.GetImageDTO{Limit: 25, Cursor: ""}
	if err := c.QueryParser(params); err != nil {
		return BadRequestError(c, err.Error(), nil)
	}

	result, err := models.GetImages(a.db, params)
	if err != nil {
		return InternalServerError(c, err.Error(), nil)
	}

	return SendOK(c, "", result)
}

func (a *API) GetImageHandler(c *fiber.Ctx) error {
	id := c.Params("id")

	image, err := models.GetImage(a.db, id)
	if err != nil {
		return InternalServerError(c, err.Error(), nil)
	}

	return SendOK(c, "", image)
}

func (a *API) DeleteImageHandler(c *fiber.Ctx) error {
	id := c.Params("id")

	image, err := models.DeleteImage(a.db, id)
	if err != nil {
		return BadRequestError(c, err.Error(), nil)
	}

	return SendOK(c, "Successfully deleted image!", image)
}

func (a *API) SearchImageHandler(c *fiber.Ctx) error {
	params := &models.SearchImageDTO{Limit: 25, Cursor: ""}
	if err := c.QueryParser(params); err != nil {
		return BadRequestError(c, err.Error(), nil)
	}

	result, err := models.Search(a.db, params)
	if err != nil {
		return InternalServerError(c, err.Error(), nil)
	}

	return SendOK(c, "", result)
}
