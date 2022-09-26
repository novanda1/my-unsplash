package models

import (
	"context"
	"errors"

	"github.com/novanda1/my-unsplash/storage"
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
	Label string `json:"label"`
	Url   string `json:"url"`
}

func SaveImage(ctx context.Context, storage *storage.Connection, p *InsertImageDTO) (*mongo.InsertOneResult, error) {
	image := Image{
		Label: p.Label,
		Url:   p.Url,
	}

	result, err := storage.ImageCollection().InsertOne(ctx, image)
	if err != nil {
		return nil, err
	}

	return result, bson.ErrDecodeToNil
}

func DeleteImage(ctx context.Context, storage *storage.Connection, id string) error {
	idPrimitive, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	result, err := storage.ImageCollection().DeleteOne(ctx, bson.M{"_id": idPrimitive})
	if err != nil {
		return err
	}

	if result.DeletedCount == 0 {
		return errors.New("object not found")
	}

	return nil
}
