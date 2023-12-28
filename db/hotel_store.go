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
	GetHotels(ctx context.Context, filter bson.M) ([]*types.Hotel, error)
	GetHotelById(ctx context.Context, id primitive.ObjectID) (*types.Hotel, error)
}

type MongoHotelStore struct{
	client *mongo.Client
	collection *mongo.Collection
}

const hotelCollection = "hotels"

func NewMongoHotelStore(client *mongo.Client) *MongoHotelStore {
	return &MongoHotelStore{
		client: client,
		collection: client.Database(DBNAME).Collection(hotelCollection),
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

func (s *MongoHotelStore) GetHotels(ctx context.Context, filter bson.M) ([]*types.Hotel, error) {
	var hotels []*types.Hotel

	cursor, err := s.collection.Find(ctx, filter)

	if err != nil {
		return nil, err
	}

	if err = cursor.All(ctx, &hotels); err != nil {
		return nil, err
	}

	return hotels, nil
}

func (s *MongoHotelStore) GetHotelById(ctx context.Context, id primitive.ObjectID) (*types.Hotel, error) {
	var hotel types.Hotel

	if err := s.collection.FindOne(ctx, bson.M{"_id": id}).Decode(&hotel); err != nil {
		return nil, err
	}

	return &hotel, nil
}

func (s *MongoHotelStore) UpdateHotel(ctx context.Context, filter bson.M, update bson.M) error {
	_, err := s.collection.UpdateOne(ctx, filter, update)

	return err
}

