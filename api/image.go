package api

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/novanda1/my-unsplash/models"
)

func (a *API) AddIMage(c *fiber.Ctx) error {
	c.Accepts("application/json")
	ctx := c.Context()

	var resp Response
	params := new(models.InsertImageDTO)

	if err := c.BodyParser(params); err != nil {
		return err
	}

	result, err := models.SaveImage(ctx, a.db, params)
	if err != nil {
		resp.Status = "error"
		resp.Data = result
		resp.Message = fmt.Sprintf("save image error: %s", err.Error())

		return c.JSON(resp)
	}

	resp.Status = "success"
	resp.Data = result
	resp.Message = ""

	return c.JSON(resp)
}
