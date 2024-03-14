package db

import (
	"github.com/arjunmalhotra1/hotel-reservation/types"
	"go.mongodb.org/mongo-driver/mongo"
)

const DBNAME = "hotel-reservation"
const UserColl = "users"

type UserStore interface {
	GetUserByID(string) (*types.User, error)
}

type MongoUserStore struct {
	client *mongo.Client
	coll   *mongo.Collection
}

func NewMongoUserStore(client *mongo.Client) *MongoUserStore {

	return &MongoUserStore{
		client: client,
		coll:   client.Database(DBNAME).Collection(UserColl),
	}
}

func (s *MongoUserStore) GetUserByID(string) (*types.User, error) {
	return nil, nil
}
