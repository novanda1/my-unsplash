package api

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/novanda1/my-unsplash/conf"
	"github.com/novanda1/my-unsplash/storage"
	"github.com/sirupsen/logrus"
)

type API struct {
	db      *storage.Connection
	handler *fiber.App
	config  *conf.GlobalConfiguration
	version string
}

type Response struct {
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func NewApi(config *conf.GlobalConfiguration, db *storage.Connection) *API {
	api := &API{config: config, version: "1"}
	api.db = db

	app := fiber.New()
	api_route := app.Group("/api")
	v1_route := api_route.Group("/v1")
	v1_image_route := v1_route.Group("/images")

	v1_image_route.Post("/", api.AddImageHandler)
	v1_image_route.Get("/", api.GetImagesHandler)
	v1_image_route.Get("/search", api.SearchImageHandler)
	v1_image_route.Get("/:id", api.GetImageHandler)
	v1_image_route.Delete("/:id", api.DeleteImageHandler)

	app.Use(logger.New())
	app.Use(cors.New(cors.Config{
		AllowOrigins: config.API.ExternalURL,
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	api.handler = app

	return api
}

func (a *API) ListenAndServe(hostAndPort string) {
	log := logrus.WithField("component", "api")

	done := make(chan struct{})
	defer close(done)

	if err := a.handler.Listen(hostAndPort); err != http.ErrServerClosed {
		log.WithError(err).Fatal("http server listen failed")
	}
}
