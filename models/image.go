package models

import (
	"context"
	"errors"
	"time"

	"github.com/novanda1/my-unsplash/storage"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Image struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Label     string             `bson:"label,omitempty" json:"label"`
	Url       string             `bson:"url,omitempty" json:"url"`
	Width     int64              `bson:"w,omitempty" json:"w"`
	Height    int64              `bson:"h,omitempty" json:"h"`
	Hash      string             `bson:"hash,omitempty" json:"hash"`
	CreatedAt int64              `json:"createdAt"`
}

type InsertImageDTO struct {
	Label  string `json:"label" validate:"required,min=3,max=25"`
	Url    string `json:"url" validate:"required,min=10"`
	Width  int64  `json:"w" validate:"required"`
	Height int64  `json:"h" validate:"required"`
	Hash   string `json:"hash" validate:"required"`
}

type GetImageDTO struct {
	Cursor string `query:"cursor"`
	Limit  int64  `query:"limit"`
}

type SearchImageDTO struct {
	Cursor string `query:"cursor"`
	Limit  int64  `query:"limit"`
	Search string `query:"search"`
}

func paginateImages(storage *storage.Connection, p *SearchImageDTO) ([]Image, error) {
	ctx := context.Background()
	images := make([]Image, 0)

	err := CreateIndex(storage)
	if err != nil {
		return nil, err
	}

	filter := bson.D{}
	opts := options.FindOptions{Limit: &p.Limit, Sort: bson.D{{Key: "_id", Value: -1}}}

	if p.Cursor != "" {
		lastId, err := primitive.ObjectIDFromHex(p.Cursor)
		if err != nil {
			return nil, err
		}

		filter = append(filter, bson.E{Key: "_id", Value: bson.D{{Key: "$lt", Value: lastId}}})
	}

	if p.Search != "" {
		filter = append(filter, bson.E{Key: "$text", Value: bson.D{{Key: "$search", Value: p.Search}}})
	}

	cursor, err := storage.ImageCollection().Find(context.TODO(), filter, &opts)
	if err != nil {
		return nil, err
	}

	for cursor.Next(ctx) {
		var image Image
		if err := cursor.Decode(&image); err != nil {
			logrus.Println(err)
		}

		images = append(images, image)
	}

	return images, nil
}

func SaveImage(storage *storage.Connection, p *InsertImageDTO) (*Image, error) {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	image := Image{
		Label:     p.Label,
		Url:       p.Url,
		Width:     p.Width,
		Height:    p.Height,
		Hash:      p.Hash,
		CreatedAt: time.Now().Unix(),
	}

	result, err := storage.ImageCollection().InsertOne(ctx, image)
	if err != nil {
		return nil, err
	}

	oid, ok := result.InsertedID.(primitive.ObjectID)
	if ok != true {
		return nil, errors.New("not valid id")
	}

	image.ID = oid

	return &image, nil
}

func DeleteImage(storage *storage.Connection, id string) (bool, error) {
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return false, err
	}

	ctx := context.Background()

	result, err := storage.ImageCollection().DeleteOne(ctx, bson.M{"_id": objectId})
	if err != nil {
		return false, err
	}

	if result.DeletedCount == 0 {
		return false, errors.New("No image found")
	}

	return true, nil
}

func GetImages(storage *storage.Connection, p *GetImageDTO) ([]Image, error) {
	var params SearchImageDTO
	params.Cursor = p.Cursor
	params.Limit = p.Limit

	images, err := paginateImages(storage, &params)
	if err != nil {
		return nil, err
	}

	return images, nil
}

func GetImage(storage *storage.Connection, id string) (*Image, error) {
	image := new(Image)

	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	ctx := context.Background()
	result := storage.ImageCollection().FindOne(ctx, bson.M{"_id": objectId})

	if result.Err() != nil {
		return nil, result.Err()
	}

	result.Decode(image)

	return image, nil
}

func CreateIndex(storage *storage.Connection) error {
	model := mongo.IndexModel{Keys: bson.D{{"label", "text"}}}
	_, err := storage.ImageCollection().Indexes().CreateOne(context.TODO(), model)
	if err != nil {
		return err
	}

	return nil
}

func Search(storage *storage.Connection, p *SearchImageDTO) ([]Image, error) {
	images, err := paginateImages(storage, p)
	if err != nil {
		return nil, err
	}

	return images, nil
}
