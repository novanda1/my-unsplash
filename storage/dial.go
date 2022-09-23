package storage

import (
	"github.com/novanda1/my-unsplash/conf"
	"github.com/pkg/errors"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Connection struct {
	*gorm.DB
}

func Dial(config *conf.GlobalConfiguration) (*Connection, error) {
	db, err := gorm.Open(postgres.Open(config.DB.URL), &gorm.Config{})
	if err != nil {
		return nil, errors.Wrap(err, "opening database connection")
	}

	return &Connection{db}, nil
}
