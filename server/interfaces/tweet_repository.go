package interfaces

import (
	"context"

	"github.com/ocionejr/upvote-klever/server/models"
)

type TweetRepositoryInterface interface {
	InsertTweet(ctx context.Context, tweet *models.Tweet) error
	FindById(ctx context.Context, id string) (*models.Tweet, error)
	ListAll() ([]models.Tweet, error)
	Update(ctx context.Context, tweet *models.Tweet, id string) error
	DeleteTweet(ctx context.Context, id string) error
	AddUpvote(ctx context.Context, tweet *models.Tweet, userId string) error
	RemoveUpvote(ctx context.Context, tweet *models.Tweet, userId string) error
}
