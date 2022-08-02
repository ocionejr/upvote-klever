package repositories

import "go.mongodb.org/mongo-driver/mongo"

type UserRepository struct {
	users *mongo.Collection
}

func NewUserRepository(database *mongo.Database) *UserRepository {
	return &UserRepository{
		users: database.Collection("users"),
	}
}