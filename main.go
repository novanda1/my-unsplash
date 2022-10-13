package main

import (
	"context"
	"fmt"
	"time"

	"github.com/novanda1/my-unsplash/api"
	"github.com/novanda1/my-unsplash/conf"
	"github.com/novanda1/my-unsplash/storage"

	"github.com/sirupsen/logrus"
)

func main() {
	config, err := conf.LoadGlobal("")
	if err != nil {
		logrus.WithError(err).Fatal("unable to load config")
	}

	db, err := storage.Dial(config)
	if err != nil {
		logrus.Fatalf("error opening database: %+v", err)
	}

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	defer db.Client.Disconnect(ctx)

	api := api.NewApi(config, db)

	l := fmt.Sprintf("%v:%v", config.API.Host, config.API.Port)
	api.ListenAndServe(l)
}
