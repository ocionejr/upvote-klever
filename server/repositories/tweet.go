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

type TweetRepository struct {
	tweets *mongo.Collection
}

func NewTweetRepository(database *mongo.Database) *TweetRepository {
	return &TweetRepository{
		tweets: database.Collection("tweets"),
	}
}

func (r *TweetRepository) InsertTweet(ctx context.Context, tweet *models.Tweet) error {
	res, err := r.tweets.InsertOne(ctx, tweet)
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

	tweet.ID = oid
	return nil
}
