package db

import (
	"context"

	"github.com/trenchesdeveloper/go-hotel/types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type HotelStore interface {
	CreateHotel(ctx context.Context, hotel *types.Hotel) (*types.Hotel, error)
	UpdateHotel(ctx context.Context, filter bson.M, update bson.M)  error

}

type MongoHotelStore struct{
	client *mongo.Client
	collection *mongo.Collection
}

const hotelCollection = "hotels"

func NewMongoHotelStore(client *mongo.Client, dbName string) *MongoHotelStore {
	return &MongoHotelStore{
		client: client,
		collection: client.Database(dbName).Collection(hotelCollection),
	}
}

func (s *MongoHotelStore) CreateHotel(ctx context.Context, hotel *types.Hotel) (*types.Hotel, error) {
	resp, err := s.collection.InsertOne(ctx, hotel)

	if err != nil {
		return nil, err
	}

	hotel.ID = resp.InsertedID.(primitive.ObjectID)

	return hotel, nil

}

func (s *MongoHotelStore) UpdateHotel(ctx context.Context, filter bson.M, update bson.M) error {
	_, err := s.collection.UpdateOne(ctx, filter, update)

	return err
}