package repositories

import (
	"context"
	"fmt"

	"github.com/ocionejr/upvote-klever/server/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type UserRepository struct {
	users *mongo.Collection
}

func NewUserRepository(database *mongo.Database) *UserRepository {
	return &UserRepository{
		users: database.Collection("users"),
	}
}

func (repository *UserRepository) InsertUser(user *models.User, ctx context.Context) error {
	res, err := repository.users.InsertOne(ctx, user)
	if err != nil {
		return status.Errorf(codes.Internal, fmt.Sprintf("Internal Error: %v\n", err))
	}

	oid, ok := res.InsertedID.(primitive.ObjectID)

	if !ok {
		return status.Errorf(
			codes.Internal,
			"Cannot convert to OID",
		)
	}

	user.ID = oid
	return nil
}