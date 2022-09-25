package storage

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/novanda1/my-unsplash/conf"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

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
	defer client.Disconnect(ctx)

	databases, err := client.ListDatabaseNames(ctx, bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("database are %v", databases)

	return &Connection{client}, nil
}
