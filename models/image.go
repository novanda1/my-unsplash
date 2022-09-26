package models

import (
	"context"
	"errors"
	"log"
	"time"

	"github.com/novanda1/my-unsplash/storage"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Image struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Label     string             `bson:"label,omitempty" json:"label"`
	Url       string             `bson:"url,omitempty" json:"url"`
	CreatedAt int64              `json:"createdAt"`
}

type InsertImageDTO struct {
	Label string `json:"label" validate:"required,min=3,max=25"`
	Url   string `json:"url" validate:"required,min=10"`
}

type GetImageDTO struct {
	Cursor string `query:"cursor"`
	Limit  int64  `query:"limit"`
}

func SaveImage(storage *storage.Connection, p *InsertImageDTO) (*mongo.InsertOneResult, error) {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	image := Image{
		Label:     p.Label,
		Url:       p.Url,
		CreatedAt: time.Now().Unix(),
	}

	result, err := storage.ImageCollection().InsertOne(ctx, image)
	if err != nil {
		return nil, err
	}

	return result, nil
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

func GetImages(storage *storage.Connection, p *GetImageDTO) ([]Image, error) {
	images := make([]Image, 0)
	ctx := context.Background()

	var lastIdPrimitive primitive.ObjectID
	lastId, err := primitive.ObjectIDFromHex(p.Cursor)
	if err == nil {
		lastIdPrimitive = lastId
	}

	opts := options.FindOptions{Limit: &p.Limit}
	filter := bson.M{
		"_id": bson.D{{Key: "$gt", Value: lastIdPrimitive}},
	}
	cursor, err := storage.ImageCollection().Find(ctx, filter, &opts)
	if err != nil {
		return nil, err
	}

	for cursor.Next(ctx) {
		var image Image
		if err := cursor.Decode(&image); err != nil {
			log.Println(err)
		}

		images = append(images, image)
	}

	return images, nil
}
