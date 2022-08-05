package interfaces

import (
	"context"

	"github.com/ocionejr/upvote-klever/server/models"
	"go.mongodb.org/mongo-driver/mongo"
)

type TweetRepositoryInterface interface {
	InsertTweet(ctx context.Context, tweet *models.Tweet) error
	FindById(ctx context.Context, id string) (*models.Tweet, error)
	ListAll() (*mongo.Cursor, error)
	Update(ctx context.Context, tweet *models.Tweet, id string) error
	DeleteTweet(ctx context.Context, id string) error
	AddUpvote(ctx context.Context, tweet *models.Tweet, userId string) error
	RemoveUpvote(ctx context.Context, tweet *models.Tweet, userId string) error
}
