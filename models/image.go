package models

import (
	"context"
	"errors"
	"fmt"
	"log"
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

type SearchImageDTO struct {
	Cursor string `query:"cursor"`
	Limit  int64  `query:"limit"`
	Search string `query:"search"`
}

func SaveImage(storage *storage.Connection, p *InsertImageDTO) (*Image, error) {
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

func Seed(storage *storage.Connection) error {
	data := []interface{}{
		Image{
			Label:     "Samsung Memory\nMemory storage made for everyone ↗",
			Url:       "https://images.unsplash.com/photo-1657214058658-51ce3eb5bebb?ixlib=rb-1.2.1&ixid=MnwxMjA3fDF8MHxlZGl0b3JpYWwtZmVlZHwxfHx8ZW58MHx8fHw%3D&w=1000&q=80",
			CreatedAt: time.Now().Unix(),
		},
		Image{
			Label:     "Jeferson Argueta\nAvailable for hire",
			Url:       "https://images.unsplash.com/photo-1664135917329-601bb51aa850?ixlib=rb-1.2.1&ixid=MnwxMjA3fDB8MHxlZGl0b3JpYWwtZmVlZHw0fHx8ZW58MHx8fHw%3D&w=1000&q=80",
			CreatedAt: time.Now().Unix(),
		},
		Image{
			Label:     "Ahmed\nAvailable for hire",
			Url:       "https://images.unsplash.com/photo-1664115532297-a2b3b8c2af38?ixlib=rb-1.2.1&ixid=MnwxMjA3fDB8MHxlZGl0b3JpYWwtZmVlZHw5fHx8ZW58MHx8fHw%3D&w=1000&q=80",
			CreatedAt: time.Now().Unix(),
		},
		Image{
			Label:     "Christine Kozak",
			Url:       "https://images.unsplash.com/photo-1664096219883-7857e422494d?ixlib=rb-1.2.1&ixid=MnwxMjA3fDB8MHxlZGl0b3JpYWwtZmVlZHwxMnx8fGVufDB8fHx8&w=1000&q=80",
			CreatedAt: time.Now().Unix(),
		},
		Image{
			Label:     "Marek Piwnicki\nAvailable for hire",
			Url:       "https://images.unsplash.com/photo-1664125010896-7d9b606c0369?ixlib=rb-1.2.1&ixid=MnwxMjA3fDB8MHxlZGl0b3JpYWwtZmVlZHwxNXx8fGVufDB8fHx8&w=1000&q=80",
			CreatedAt: time.Now().Unix(),
		},
		Image{
			Label:     "Pawel Czerwinski",
			Url:       "https://images.unsplash.com/photo-1664111544499-aa9a6c7de16f?ixlib=rb-1.2.1&ixid=MnwxMjA3fDB8MHxlZGl0b3JpYWwtZmVlZHwxN3x8fGVufDB8fHx8&w=1000&q=80",
			CreatedAt: time.Now().Unix(),
		},
		Image{
			Label:     "Sung Jin Cho",
			Url:       "https://images.unsplash.com/photo-1664104995040-49d9a268ab7c?ixlib=rb-1.2.1&ixid=MnwxMjA3fDB8MHxlZGl0b3JpYWwtZmVlZHwyMHx8fGVufDB8fHx8&w=1000&q=80",
			CreatedAt: time.Now().Unix(),
		},
		Image{
			Label:     "Samsung Memory\nMemory storage made for everyone ↗",
			Url:       "https://images.unsplash.com/photo-1659535880591-78ed91b35158?ixlib=rb-1.2.1&ixid=MnwxMjA3fDF8MHxlZGl0b3JpYWwtZmVlZHwyMXx8fGVufDB8fHx8&w=1000&q=80",
			CreatedAt: time.Now().Unix(),
		},
		Image{
			Label:     "Joel Lee",
			Url:       "https://images.unsplash.com/photo-1664111601108-993c57ee124b?ixlib=rb-1.2.1&ixid=MnwxMjA3fDB8MHxlZGl0b3JpYWwtZmVlZHwyNHx8fGVufDB8fHx8&w=1000&q=80",
			CreatedAt: time.Now().Unix(),
		},
		Image{
			Label:     "Steffen Lemmerzahl",
			Url:       "https://images.unsplash.com/photo-1664112742677-9478378e31e5?ixlib=rb-1.2.1&ixid=MnwxMjA3fDB8MHxlZGl0b3JpYWwtZmVlZHwyN3x8fGVufDB8fHx8&w=1000&q=80",
			CreatedAt: time.Now().Unix(),
		},
		Image{
			Label:     "Windows\nCreate great things with Windows 11 & Microsoft 365 ↗",
			Url:       "https://images.unsplash.com/photo-1662581872277-0fd0bf3ae8f6?ixlib=rb-1.2.1&ixid=MnwxMjA3fDF8MHxlZGl0b3JpYWwtZmVlZHwzMHx8fGVufDB8fHx8&w=1000&q=80",
			CreatedAt: time.Now().Unix(),
		},
		Image{
			Label:     "Susie Burleson\nAvailable for hire",
			Url:       "https://images.unsplash.com/photo-1664122802538-19fcff974fea?ixlib=rb-1.2.1&ixid=MnwxMjA3fDB8MHxlZGl0b3JpYWwtZmVlZHwzM3x8fGVufDB8fHx8&w=1000&q=80",
			CreatedAt: time.Now().Unix(),
		},
		Image{
			Label:     "Kellen Riggin\nAvailable for hire",
			Url:       "https://images.unsplash.com/photo-1664123238749-44f8830365e5?ixlib=rb-1.2.1&ixid=MnwxMjA3fDB8MHxlZGl0b3JpYWwtZmVlZHwzNnx8fGVufDB8fHx8&w=1000&q=80",
			CreatedAt: time.Now().Unix(),
		},
		Image{
			Label:     "Ali Drabo",
			Url:       "https://images.unsplash.com/photo-1663977574293-f65523f4df89?ixlib=rb-1.2.1&ixid=MnwxMjA3fDB8MHxlZGl0b3JpYWwtZmVlZHwzOXx8fGVufDB8fHx8&w=1000&q=80",
			CreatedAt: time.Now().Unix(),
		},
		Image{
			Label:     "Jez Timms\nAvailable for hire",
			Url:       "https://images.unsplash.com/photo-1664107014454-3d5758bc32e5?ixlib=rb-1.2.1&ixid=MnwxMjA3fDB8MHxlZGl0b3JpYWwtZmVlZHw0Mnx8fGVufDB8fHx8&w=1000&q=80",
			CreatedAt: time.Now().Unix(),
		},
		Image{
			Label:     "Jeferson Argueta\nAvailable for hire",
			Url:       "https://images.unsplash.com/photo-1664136162748-88b2d5edaea1?ixlib=rb-1.2.1&ixid=MnwxMjA3fDB8MHxlZGl0b3JpYWwtZmVlZHw0NHx8fGVufDB8fHx8&w=1000&q=80",
			CreatedAt: time.Now().Unix(),
		},
		Image{
			Label:     "Windows\nCreate great things with Windows 11 & Microsoft 365 ↗",
			Url:       "https://images.unsplash.com/photo-1662581872342-3f8e0145668f?ixlib=rb-1.2.1&ixid=MnwxMjA3fDF8MHxlZGl0b3JpYWwtZmVlZHw0OHx8fGVufDB8fHx8&w=1000&q=80",
			CreatedAt: time.Now().Unix(),
		},
		Image{
			Label:     "Marek Piwnicki\nAvailable for hire",
			Url:       "https://images.unsplash.com/photo-1664125010409-8a7f4d82ec07?ixlib=rb-1.2.1&ixid=MnwxMjA3fDB8MHxlZGl0b3JpYWwtZmVlZHw0OXx8fGVufDB8fHx8&w=1000&q=80",
			CreatedAt: time.Now().Unix(),
		},
		Image{
			Label:     "Pawel Czerwinski",
			Url:       "https://images.unsplash.com/photo-1664037109833-5230a4640662?ixlib=rb-1.2.1&ixid=MnwxMjA3fDB8MHxlZGl0b3JpYWwtZmVlZHw1MXx8fGVufDB8fHx8&w=1000&q=80",
			CreatedAt: time.Now().Unix(),
		},
		Image{
			Label:     "Boxed Water Is Better\nPlant-based. Build a better planet. ↗",
			Url:       "https://images.unsplash.com/photo-1659482633453-4e51c3e95112?ixlib=rb-1.2.1&ixid=MnwxMjA3fDF8MHxlZGl0b3JpYWwtZmVlZHw1M3x8fGVufDB8fHx8&w=1000&q=80",
			CreatedAt: time.Now().Unix(),
		},
		Image{
			Label:     "Jeferson Argueta\nAvailable for hire",
			Url:       "https://images.unsplash.com/photo-1663409189430-418c72ff3bad?ixlib=rb-1.2.1&ixid=MnwxMjA3fDB8MHxlZGl0b3JpYWwtZmVlZHw1Nnx8fGVufDB8fHx8&w=1000&q=80",
			CreatedAt: time.Now().Unix(),
		},
		Image{
			Label:     "Julia Blumberg\nAvailable for hire",
			Url:       "https://images.unsplash.com/photo-1664058267614-ccb00305ee75?ixlib=rb-1.2.1&ixid=MnwxMjA3fDB8MHxlZGl0b3JpYWwtZmVlZHw1OXx8fGVufDB8fHx8&w=1000&q=80",
			CreatedAt: time.Now().Unix(),
		},
		Image{
			Label:     "Collin Ross",
			Url:       "https://images.unsplash.com/photo-1663877254454-5e669a7b27e6?ixlib=rb-1.2.1&ixid=MnwxMjA3fDB8MHxlZGl0b3JpYWwtZmVlZHw2Mnx8fGVufDB8fHx8&w=1000&q=80",
			CreatedAt: time.Now().Unix(),
		},
		Image{
			Label:     "Martin Eriksson",
			Url:       "https://images.unsplash.com/photo-1664033333006-6f8d1e29bd14?ixlib=rb-1.2.1&ixid=MnwxMjA3fDB8MHxlZGl0b3JpYWwtZmVlZHw2NXx8fGVufDB8fHx8&w=1000&q=80",
			CreatedAt: time.Now().Unix(),
		},
		Image{
			Label:     "Pawel Czerwinski",
			Url:       "https://images.unsplash.com/photo-1664040271546-be29247975af?ixlib=rb-1.2.1&ixid=MnwxMjA3fDB8MHxlZGl0b3JpYWwtZmVlZHw2N3x8fGVufDB8fHx8&w=1000&q=80",
			CreatedAt: time.Now().Unix(),
		},
		Image{
			Label:     "Susan Wilkinson\nAvailable for hire",
			Url:       "https://images.unsplash.com/photo-1664042913846-abb18eba6226?ixlib=rb-1.2.1&ixid=MnwxMjA3fDB8MHxlZGl0b3JpYWwtZmVlZHw3MHx8fGVufDB8fHx8&w=1000&q=80",
			CreatedAt: time.Now().Unix(),
		},
		Image{
			Label:     "Megan O'Hanlon\nAvailable for hire",
			Url:       "https://images.unsplash.com/photo-1663788917276-7dfeddd2683c?ixlib=rb-1.2.1&ixid=MnwxMjA3fDB8MHxlZGl0b3JpYWwtZmVlZHw3NHx8fGVufDB8fHx8&w=1000&q=80",
			CreatedAt: time.Now().Unix(),
		},
		Image{
			Label:     "Samsung Memory\nMemory storage made for everyone ↗",
			Url:       "https://images.unsplash.com/photo-1657214059139-dc58d16118ed?ixlib=rb-1.2.1&ixid=MnwxMjA3fDF8MHxlZGl0b3JpYWwtZmVlZHw3N3x8fGVufDB8fHx8&w=1000&q=80",
			CreatedAt: time.Now().Unix(),
		},
		Image{
			Label:     "Jeferson Argueta\nAvailable for hire",
			Url:       "https://images.unsplash.com/photo-1663867405122-c0d4c29cece3?ixlib=rb-1.2.1&ixid=MnwxMjA3fDB8MHxlZGl0b3JpYWwtZmVlZHw4MHx8fGVufDB8fHx8&w=1000&q=80",
			CreatedAt: time.Now().Unix(),
		},
		Image{
			Label:     "Kateryna Hliznitsova\nAvailable for hire",
			Url:       "https://images.unsplash.com/photo-1664029593174-383050751e65?ixlib=rb-1.2.1&ixid=MnwxMjA3fDB8MHxlZGl0b3JpYWwtZmVlZHw4NXx8fGVufDB8fHx8&w=1000&q=80",
			CreatedAt: time.Now().Unix(),
		},
		Image{
			Label:     "David Karp\nAvailable for hire",
			Url:       "https://images.unsplash.com/photo-1664018625610-7538a8367f88?ixlib=rb-1.2.1&ixid=MnwxMjA3fDB8MHxlZGl0b3JpYWwtZmVlZHw4OHx8fGVufDB8fHx8&w=1000&q=80",
			CreatedAt: time.Now().Unix(),
		},
		Image{
			Label:     "Boxed Water Is Better\nPlant-based. Build a better planet. ↗",
			Url:       "https://images.unsplash.com/photo-1590074072786-a66914d668f1?ixlib=rb-1.2.1&ixid=MnwxMjA3fDF8MHxlZGl0b3JpYWwtZmVlZHw5Mnx8fGVufDB8fHx8&w=1000&q=80",
			CreatedAt: time.Now().Unix(),
		},
		Image{
			Label:     "Christine Kozak",
			Url:       "https://images.unsplash.com/photo-1664096555728-b4de88d0455f?ixlib=rb-1.2.1&ixid=MnwxMjA3fDB8MHxlZGl0b3JpYWwtZmVlZHw5M3x8fGVufDB8fHx8&w=1000&q=80",
			CreatedAt: time.Now().Unix(),
		},
		Image{
			Label:     "Kellen Riggin\nAvailable for hire",
			Url:       "https://images.unsplash.com/photo-1664054698371-c304f40d436a?ixlib=rb-1.2.1&ixid=MnwxMjA3fDB8MHxlZGl0b3JpYWwtZmVlZHw5Nnx8fGVufDB8fHx8&w=1000&q=80",
			CreatedAt: time.Now().Unix(),
		},
		Image{
			Label:     "martin bennie\nAvailable for hire",
			Url:       "https://images.unsplash.com/photo-1664030311590-3162a492ab1c?ixlib=rb-1.2.1&ixid=MnwxMjA3fDB8MHxlZGl0b3JpYWwtZmVlZHwxMDF8fHxlbnwwfHx8fA%3D%3D&w=1000&q=80",
			CreatedAt: time.Now().Unix(),
		},
		Image{
			Label:     "Jessica Christian\nAvailable for hire",
			Url:       "https://images.unsplash.com/photo-1664047257778-6c60bedfc9c4?ixlib=rb-1.2.1&ixid=MnwxMjA3fDB8MHxlZGl0b3JpYWwtZmVlZHwxMDR8fHxlbnwwfHx8fA%3D%3D&w=1000&q=80",
			CreatedAt: time.Now().Unix(),
		},
		Image{
			Label:     "alea Film\nAvailable for hire",
			Url:       "https://images.unsplash.com/photo-1664027609936-1f89ef716c7b?ixlib=rb-1.2.1&ixid=MnwxMjA3fDB8MHxlZGl0b3JpYWwtZmVlZHwxMDZ8fHxlbnwwfHx8fA%3D%3D&w=1000&q=80",
			CreatedAt: time.Now().Unix(),
		},
		Image{
			Label:     "Kellen Riggin\nAvailable for hire",
			Url:       "https://images.unsplash.com/photo-1664075550378-6b07dd9830a0?ixlib=rb-1.2.1&ixid=MnwxMjA3fDB8MHxlZGl0b3JpYWwtZmVlZHwxMDl8fHxlbnwwfHx8fA%3D%3D&w=1000&q=80",
			CreatedAt: time.Now().Unix(),
		},
		Image{
			Label:     "Ben Iwara\nAvailable for hire",
			Url:       "https://images.unsplash.com/photo-1664076709329-a0f5f0ac6b31?ixlib=rb-1.2.1&ixid=MnwxMjA3fDB8MHxlZGl0b3JpYWwtZmVlZHwxMTB8fHxlbnwwfHx8fA%3D%3D&w=1000&q=80",
			CreatedAt: time.Now().Unix(),
		},
		Image{
			Label:     "Dibakar Roy\nAvailable for hire",
			Url:       "https://images.unsplash.com/photo-1664077857279-603b1a425d12?ixlib=rb-1.2.1&ixid=MnwxMjA3fDB8MHxlZGl0b3JpYWwtZmVlZHwxMTR8fHxlbnwwfHx8fA%3D%3D&w=1000&q=80",
			CreatedAt: time.Now().Unix(),
		},
		Image{
			Label:     "Samsung Memory\nMemory storage made for everyone ↗",
			Url:       "https://images.unsplash.com/photo-1659536540446-441fcafdd73a?ixlib=rb-1.2.1&ixid=MnwxMjA3fDF8MHxlZGl0b3JpYWwtZmVlZHwxMTd8fHxlbnwwfHx8fA%3D%3D&w=1000&q=80",
			CreatedAt: time.Now().Unix(),
		},
		Image{
			Label:     "Viktor Mindt",
			Url:       "https://images.unsplash.com/photo-1664025412567-52b83a54a233?ixlib=rb-1.2.1&ixid=MnwxMjA3fDB8MHxlZGl0b3JpYWwtZmVlZHwxMTl8fHxlbnwwfHx8fA%3D%3D&w=1000&q=80",
			CreatedAt: time.Now().Unix(),
		},
		Image{
			Label:     "Vladislav Nahorny\nAvailable for hire",
			Url:       "https://images.unsplash.com/photo-1664052961779-607fab0ea1d2?ixlib=rb-1.2.1&ixid=MnwxMjA3fDB8MHxlZGl0b3JpYWwtZmVlZHwxMjB8fHxlbnwwfHx8fA%3D%3D&w=1000&q=80",
			CreatedAt: time.Now().Unix(),
		},
		Image{
			Label:     "Mailchimp\nGet the advanced tools you need to grow ↗",
			Url:       "https://images.unsplash.com/photo-1661956602139-ec64991b8b16?ixlib=rb-1.2.1&ixid=MnwxMjA3fDF8MHxlZGl0b3JpYWwtZmVlZHwxMjJ8fHxlbnwwfHx8fA%3D%3D&w=1000&q=80",
			CreatedAt: time.Now().Unix(),
		},
		Image{
			Label:     "Mark Bishop\nAvailable for hire",
			Url:       "https://images.unsplash.com/photo-1664056350275-f7121af5cff0?ixlib=rb-1.2.1&ixid=MnwxMjA3fDB8MHxlZGl0b3JpYWwtZmVlZHwxMjV8fHxlbnwwfHx8fA%3D%3D&w=1000&q=80",
			CreatedAt: time.Now().Unix(),
		},
		Image{
			Label:     "Lauren Pelesky\nAvailable for hire",
			Url:       "https://images.unsplash.com/photo-1663765583971-8804289c2f38?ixlib=rb-1.2.1&ixid=MnwxMjA3fDB8MHxlZGl0b3JpYWwtZmVlZHwxMjh8fHxlbnwwfHx8fA%3D%3D&w=1000&q=80",
			CreatedAt: time.Now().Unix(),
		},
		Image{
			Label:     "Samsung Memory\nMemory storage made for everyone ↗",
			Url:       "https://images.unsplash.com/photo-1659536540437-510ce63eb672?ixlib=rb-1.2.1&ixid=MnwxMjA3fDF8MHxlZGl0b3JpYWwtZmVlZHwxMzJ8fHxlbnwwfHx8fA%3D%3D&w=1000&q=80",
			CreatedAt: time.Now().Unix(),
		},
		Image{
			Label:     "Lauren Pelesky\nAvailable for hire",
			Url:       "https://images.unsplash.com/photo-1662990502014-9b73af9383d9?ixlib=rb-1.2.1&ixid=MnwxMjA3fDB8MHxlZGl0b3JpYWwtZmVlZHwxMzN8fHxlbnwwfHx8fA%3D%3D&w=1000&q=80",
			CreatedAt: time.Now().Unix(),
		},
		Image{
			Label:     "Andrea BA\nAvailable for hire",
			Url:       "https://images.unsplash.com/photo-1655749534279-20792c3a546d?ixlib=rb-1.2.1&ixid=MnwxMjA3fDB8MHxlZGl0b3JpYWwtZmVlZHwxMzZ8fHxlbnwwfHx8fA%3D%3D&w=1000&q=80",
			CreatedAt: time.Now().Unix(),
		},
		Image{
			Label:     "Jackson Graham",
			Url:       "https://images.unsplash.com/photo-1663979747290-57d502e9b63f?ixlib=rb-1.2.1&ixid=MnwxMjA3fDB8MHxlZGl0b3JpYWwtZmVlZHwxNDB8fHxlbnwwfHx8fA%3D%3D&w=1000&q=80",
			CreatedAt: time.Now().Unix(),
		},
		Image{
			Label:     "MARIOLA GROBELSKA\nAvailable for hire",
			Url:       "https://images.unsplash.com/photo-1663966752161-affc2600c845?ixlib=rb-1.2.1&ixid=MnwxMjA3fDB8MHxlZGl0b3JpYWwtZmVlZHwxNDR8fHxlbnwwfHx8fA%3D%3D&w=1000&q=80",
			CreatedAt: time.Now().Unix(),
		},
		Image{
			Label:     "Frank Ching",
			Url:       "https://images.unsplash.com/photo-1663996806932-357eddab9b50?ixlib=rb-1.2.1&ixid=MnwxMjA3fDB8MHxlZGl0b3JpYWwtZmVlZHwxNDZ8fHxlbnwwfHx8fA%3D%3D&w=1000&q=80",
			CreatedAt: time.Now().Unix(),
		},
		Image{
			Label:     "Marek Piwnicki\nAvailable for hire",
			Url:       "https://images.unsplash.com/photo-1663953009099-a83b8c8d065b?ixlib=rb-1.2.1&ixid=MnwxMjA3fDB8MHxlZGl0b3JpYWwtZmVlZHwxNTB8fHxlbnwwfHx8fA%3D%3D&w=1000&q=80",
			CreatedAt: time.Now().Unix(),
		},
		Image{
			Label:     "Mailchimp\nDesign on-brand assets with just a click ↗",
			Url:       "https://images.unsplash.com/photo-1661956602926-db6b25f75947?ixlib=rb-1.2.1&ixid=MnwxMjA3fDF8MHxlZGl0b3JpYWwtZmVlZHwxNTJ8fHxlbnwwfHx8fA%3D%3D&w=1000&q=80",
			CreatedAt: time.Now().Unix(),
		},
		Image{
			Label:     "Shayna Douglas\nAvailable for hire",
			Url:       "https://images.unsplash.com/photo-1663942465104-f6f4d7774472?ixlib=rb-1.2.1&ixid=MnwxMjA3fDB8MHxlZGl0b3JpYWwtZmVlZHwxNTV8fHxlbnwwfHx8fA%3D%3D&w=1000&q=80",
			CreatedAt: time.Now().Unix(),
		},
		Image{
			Label:     "Viktor Mindt",
			Url:       "https://images.unsplash.com/photo-1663949611868-c78e82498286?ixlib=rb-1.2.1&ixid=MnwxMjA3fDB8MHxlZGl0b3JpYWwtZmVlZHwxNTh8fHxlbnwwfHx8fA%3D%3D&w=1000&q=80",
			CreatedAt: time.Now().Unix(),
		},
		Image{
			Label:     "Matt Drenth\nAvailable for hire",
			Url:       "https://images.unsplash.com/photo-1663875967691-15b02702931c?ixlib=rb-1.2.1&ixid=MnwxMjA3fDB8MHxlZGl0b3JpYWwtZmVlZHwxNjB8fHxlbnwwfHx8fA%3D%3D&w=1000&q=80",
			CreatedAt: time.Now().Unix(),
		},
		Image{
			Label:     "Samsung Memory\nMemory storage made for everyone ↗",
			Url:       "https://images.unsplash.com/photo-1659535998184-15d6c9f5f873?ixlib=rb-1.2.1&ixid=MnwxMjA3fDF8MHxlZGl0b3JpYWwtZmVlZHwxNjJ8fHxlbnwwfHx8fA%3D%3D&w=1000&q=80",
			CreatedAt: time.Now().Unix(),
		},
		Image{
			Label:     "Alex Marinez",
			Url:       "https://images.unsplash.com/photo-1663940019982-c14294717dbd?ixlib=rb-1.2.1&ixid=MnwxMjA3fDB8MHxlZGl0b3JpYWwtZmVlZHwxNjZ8fHxlbnwwfHx8fA%3D%3D&w=1000&q=80",
			CreatedAt: time.Now().Unix(),
		},
		Image{
			Label:     "Mailchimp\nStart targeting valuable customers today ↗",
			Url:       "https://images.unsplash.com/photo-1661956600655-e772b2b97db4?ixlib=rb-1.2.1&ixid=MnwxMjA3fDF8MHxlZGl0b3JpYWwtZmVlZHwxNjd8fHxlbnwwfHx8fA%3D%3D&w=1000&q=80",
			CreatedAt: time.Now().Unix(),
		},
		Image{
			Label:     "Vladislav Nahorny\nAvailable for hire",
			Url:       "https://images.unsplash.com/photo-1663947010000-3096c03d5ef5?ixlib=rb-1.2.1&ixid=MnwxMjA3fDB8MHxlZGl0b3JpYWwtZmVlZHwxNzB8fHxlbnwwfHx8fA%3D%3D&w=1000&q=80",
			CreatedAt: time.Now().Unix(),
		},
		Image{
			Label:     "Isaac Mitchell\nAvailable for hire",
			Url:       "https://images.unsplash.com/photo-1663895832076-5715b1f7c59b?ixlib=rb-1.2.1&ixid=MnwxMjA3fDB8MHxlZGl0b3JpYWwtZmVlZHwxNzR8fHxlbnwwfHx8fA%3D%3D&w=1000&q=80",
			CreatedAt: time.Now().Unix(),
		},
		Image{
			Label:     "Gantas Vaičiulėnas\nAvailable for hire",
			Url:       "https://images.unsplash.com/photo-1663873881284-6dd4aeacdac9?ixlib=rb-1.2.1&ixid=MnwxMjA3fDB8MHxlZGl0b3JpYWwtZmVlZHwxNzZ8fHxlbnwwfHx8fA%3D%3D&w=1000&q=80",
			CreatedAt: time.Now().Unix(),
		},
		Image{
			Label:     "op23\nAvailable for hire",
			Url:       "https://images.unsplash.com/photo-1663930981910-7a4c077feb85?ixlib=rb-1.2.1&ixid=MnwxMjA3fDB8MHxlZGl0b3JpYWwtZmVlZHwxODB8fHxlbnwwfHx8fA%3D%3D&w=1000&q=80",
			CreatedAt: time.Now().Unix(),
		},
		Image{
			Label:     "Lauren Pelesky\nAvailable for hire",
			Url:       "https://images.unsplash.com/photo-1663711920061-6ad463a32456?ixlib=rb-1.2.1&ixid=MnwxMjA3fDB8MHxlZGl0b3JpYWwtZmVlZHwxODN8fHxlbnwwfHx8fA%3D%3D&w=1000&q=80",
			CreatedAt: time.Now().Unix(),
		},
		Image{
			Label:     "Viktor Mindt",
			Url:       "https://images.unsplash.com/photo-1663942535328-4adb74bb0380?ixlib=rb-1.2.1&ixid=MnwxMjA3fDB8MHxlZGl0b3JpYWwtZmVlZHwxODV8fHxlbnwwfHx8fA%3D%3D&w=1000&q=80",
			CreatedAt: time.Now().Unix(),
		},
		Image{
			Label:     "Boxed Water Is Better\nPlant-based. Build a better planet. ↗",
			Url:       "https://images.unsplash.com/photo-1553531384-cc64ac80f931?ixlib=rb-1.2.1&ixid=MnwxMjA3fDF8MHxlZGl0b3JpYWwtZmVlZHwxODd8fHxlbnwwfHx8fA%3D%3D&w=1000&q=80",
			CreatedAt: time.Now().Unix(),
		},
		Image{
			Label:     "Jeremy Stewardson",
			Url:       "https://images.unsplash.com/photo-1663924369654-f0070587b6be?ixlib=rb-1.2.1&ixid=MnwxMjA3fDB8MHxlZGl0b3JpYWwtZmVlZHwxOTB8fHxlbnwwfHx8fA%3D%3D&w=1000&q=80",
			CreatedAt: time.Now().Unix(),
		},
		Image{
			Label:     "Samsung Memory\nMemory storage made for everyone ↗",
			Url:       "https://images.unsplash.com/photo-1659535973636-6cef468d093b?ixlib=rb-1.2.1&ixid=MnwxMjA3fDF8MHxlZGl0b3JpYWwtZmVlZHwxOTJ8fHxlbnwwfHx8fA%3D%3D&w=1000&q=80",
			CreatedAt: time.Now().Unix(),
		},
		Image{
			Label:     "martin bennie\nAvailable for hire",
			Url:       "https://images.unsplash.com/photo-1663776211814-ffdd8553e00d?ixlib=rb-1.2.1&ixid=MnwxMjA3fDB8MHxlZGl0b3JpYWwtZmVlZHwxOTZ8fHxlbnwwfHx8fA%3D%3D&w=1000&q=80",
			CreatedAt: time.Now().Unix(),
		},
		Image{
			Label:     "Mailchimp\nStart creating better email content today ↗",
			Url:       "https://images.unsplash.com/photo-1661956600654-edac218fea67?ixlib=rb-1.2.1&ixid=MnwxMjA3fDF8MHxlZGl0b3JpYWwtZmVlZHwxOTd8fHxlbnwwfHx8fA%3D%3D&w=1000&q=80",
			CreatedAt: time.Now().Unix(),
		},
		Image{
			Label:     "Boxed Water Is Better\nPlant-based. Build a better planet. ↗",
			Url:       "https://images.unsplash.com/photo-1553531768-88af16561c0f?ixlib=rb-1.2.1&ixid=MnwxMjA3fDF8MHxlZGl0b3JpYWwtZmVlZHwyMDJ8fHxlbnwwfHx8fA%3D%3D&w=1000&q=80",
			CreatedAt: time.Now().Unix(),
		},
		Image{
			Label:     "Matt Drenth\nAvailable for hire",
			Url:       "https://images.unsplash.com/photo-1663875972135-bb6654ed702d?ixlib=rb-1.2.1&ixid=MnwxMjA3fDB8MHxlZGl0b3JpYWwtZmVlZHwyMDV8fHxlbnwwfHx8fA%3D%3D&w=1000&q=80",
			CreatedAt: time.Now().Unix(),
		},
		Image{
			Label:     "Samsung Memory\nMemory storage made for everyone ↗",
			Url:       "https://images.unsplash.com/photo-1661347332466-9b6897c93a2c?ixlib=rb-1.2.1&ixid=MnwxMjA3fDF8MHxlZGl0b3JpYWwtZmVlZHwyMDd8fHxlbnwwfHx8fA%3D%3D&w=1000&q=80",
			CreatedAt: time.Now().Unix(),
		},
		Image{
			Label:     "Lauren Pelesky\nAvailable for hire",
			Url:       "https://images.unsplash.com/photo-1663711935287-4a7323fea555?ixlib=rb-1.2.1&ixid=MnwxMjA3fDB8MHxlZGl0b3JpYWwtZmVlZHwyMTB8fHxlbnwwfHx8fA%3D%3D&w=1000&q=80",
			CreatedAt: time.Now().Unix(),
		},
		Image{
			Label:     "Michiel Annaert\nAvailable for hire",
			Url:       "https://images.unsplash.com/photo-1663832669528-df3f45bce473?ixlib=rb-1.2.1&ixid=MnwxMjA3fDB8MHxlZGl0b3JpYWwtZmVlZHwyMTR8fHxlbnwwfHx8fA%3D%3D&w=1000&q=80",
			CreatedAt: time.Now().Unix(),
		},
		Image{
			Label:     "Boxed Water Is Better\nPlant-based. Build a better planet. ↗",
			Url:       "https://images.unsplash.com/photo-1570654639102-bdd95efeca7a?ixlib=rb-1.2.1&ixid=MnwxMjA3fDF8MHxlZGl0b3JpYWwtZmVlZHwyMTd8fHxlbnwwfHx8fA%3D%3D&w=1000&q=80",
			CreatedAt: time.Now().Unix(),
		},
		Image{
			Label:     "Michiel Annaert\nAvailable for hire",
			Url:       "https://images.unsplash.com/photo-1663832670362-5c91a48f0b44?ixlib=rb-1.2.1&ixid=MnwxMjA3fDB8MHxlZGl0b3JpYWwtZmVlZHwyMjF8fHxlbnwwfHx8fA%3D%3D&w=1000&q=80",
			CreatedAt: time.Now().Unix(),
		},
		Image{
			Label:     "Natali Hordiiuk\nAvailable for hire",
			Url:       "https://images.unsplash.com/photo-1663188646682-7bdc8a7738d9?ixlib=rb-1.2.1&ixid=MnwxMjA3fDB8MHxlZGl0b3JpYWwtZmVlZHwyMjR8fHxlbnwwfHx8fA%3D%3D&w=1000&q=80",
			CreatedAt: time.Now().Unix(),
		},
	}

	result, err := storage.ImageCollection().InsertMany(context.TODO(), data)
	if err != nil {
		return errors.New(fmt.Sprintf("insert error: %s", err.Error()))
	}

	fmt.Print(result.InsertedIDs...)

	return nil
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
	ctx := context.Background()
	images := make([]Image, 0)

	err := CreateIndex(storage)
	if err != nil {
		return nil, err
	}

	opts := options.FindOptions{Limit: &p.Limit}
	filter := bson.M{
		"$text": bson.D{{Key: "$search", Value: p.Search}},
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
