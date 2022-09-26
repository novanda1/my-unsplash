package models

import (
	"context"

	"github.com/novanda1/my-unsplash/storage"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Image struct {
	ID    primitive.ObjectID `bson:"_id,omitempty"`
	Label string             `bson:"label,omitempty"`
	Url   string             `bson:"url,omitempty"`
}

type InsertImageDTO struct {
	label string
	url   string
}

func (Image) SaveImage(ctx context.Context, storage *storage.Connection, p InsertImageDTO) *mongo.InsertOneResult {
	image := Image{
		Label: p.label,
		Url:   p.url,
	}

	result, err := storage.ImageCollection().InsertOne(ctx, image)
	if err != nil {
		logrus.Panicf("Failed to insert image %s", err.Error())
	}

	return result
}

func (Image) DeleteImage(ctx context.Context, storage *storage.Connection, id string) {
	idPrimitive, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		logrus.Fatal("primitive.ObjectIDFromHex ERROR:", err)
	}

	result, err := storage.ImageCollection().DeleteOne(ctx, bson.M{"_id": idPrimitive})
	if err != nil {
		logrus.Fatal("error deleting object:", err)
	}

	if result.DeletedCount == 0 {
		logrus.Error("object not found")
	}
}
