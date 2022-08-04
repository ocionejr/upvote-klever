package repositories

import "go.mongodb.org/mongo-driver/mongo"

type TweetRepository struct {
	tweets *mongo.Collection
}

func NewTweetRepository(database *mongo.Database) *TweetRepository {
	return &TweetRepository{
		tweets: database.Collection("tweets"),
	}
}
