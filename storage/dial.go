package storage

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/novanda1/my-unsplash/conf"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"

	"github.com/pkg/errors"
)

type Connection struct {
	*mongo.Client
}

func Dial(config *conf.GlobalConfiguration) (*Connection, error) {
	client, err := mongo.NewClient(options.Client().ApplyURI(config.DB.URL))

	if err != nil {
		return nil, errors.Wrap(err, "opening database connection")
	}

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	if err := client.Ping(context.TODO(), readpref.Primary()); err != nil {
		panic(err)
	}
	fmt.Println("Mongo DB: Successfully connected and pinged.")

	return &Connection{client}, nil
}

func (c *Connection) UnsplashDatabase() *mongo.Database {
	return c.Client.Database("unsplash")
}

func (c *Connection) ImageCollection() *mongo.Collection {
	return c.UnsplashDatabase().Collection("images")
}
