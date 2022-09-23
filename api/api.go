package api

import (
	"net/http"
	"time"

	"github.com/novanda1/my-unsplash/conf"
	"github.com/novanda1/my-unsplash/storage"
	"github.com/sirupsen/logrus"
)

type API struct {
	handler http.Handler
	config  *conf.GlobalConfiguration
	version string
}

func NewApi(config *conf.GlobalConfiguration, db *storage.Connection) *API {
	api := &API{config: config, version: "1"}
	api.handler = nil

	return api
}

func (a *API) ListenAndServe(hostAndPort string) {
	log := logrus.WithField("component", "api")
	server := &http.Server{
		Addr:              hostAndPort,
		Handler:           a.handler,
		ReadHeaderTimeout: 2 * time.Second, // to mitigate a Slowloris attack
	}

	done := make(chan struct{})
	defer close(done)

	if err := server.ListenAndServe(); err != http.ErrServerClosed {
		log.WithError(err).Fatal("http server listen failed")
	}
}
