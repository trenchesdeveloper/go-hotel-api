package db

import (
	"context"

	"github.com/trenchesdeveloper/go-hotel/types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type BookingStore interface {
	CreateBooking(ctx context.Context, room *types.Booking) (*types.Booking, error)
	GetBookings(ctx context.Context, filter bson.M) ([]*types.Booking, error)
}

type MongoBookingStore struct {
	client     *mongo.Client
	collection *mongo.Collection
}

const bookingCollection = "booking"

func NewMongoBookingStore(client *mongo.Client) *MongoBookingStore {
	return &MongoBookingStore{
		client:     client,
		collection: client.Database(DBNAME).Collection(roomCollection),
	}
}

func (s *MongoBookingStore) CreateBooking(ctx context.Context, booking *types.Booking) (*types.Booking, error) {
	resp, err := s.collection.InsertOne(ctx, booking)
	if err != nil {
		return nil, err
	}

	booking.ID = resp.InsertedID.(primitive.ObjectID)

	return booking, nil
}

func (s *MongoBookingStore) GetBookings(ctx context.Context, filter bson.M) ([]*types.Booking, error) {
	var bookings []*types.Booking

	cursor, err := s.collection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}

	if err := cursor.All(ctx, &bookings); err != nil {
		return nil, err
	}

	return bookings, nil
}
