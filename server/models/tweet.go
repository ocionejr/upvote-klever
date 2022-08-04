package models

import (
	"time"

	pb "github.com/ocionejr/upvote-klever/pb"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type Tweet struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	AuthorId  string             `bson:"author_id"`
	Message   string             `bson:"message"`
	Upvotes   []string           `bson:"upvotes,omitempty"`
	CreatedAt primitive.DateTime `bson:"createdAt"`
	UpdatedAt primitive.DateTime `bson:"updatedAt"`
}

func TweetRequestToTweet(tweetRequest *pb.TweetRequest) *Tweet {
	return &Tweet{
		AuthorId:  tweetRequest.AuthorId,
		Message:   tweetRequest.Message,
		CreatedAt: primitive.NewDateTimeFromTime(time.Now()),
		UpdatedAt: primitive.NewDateTimeFromTime(time.Now()),
	}
}

func TweetToTweetResponse(tweet *Tweet) *pb.TweetResponse {
	return &pb.TweetResponse{
		Id:        tweet.ID.Hex(),
		AuthorId:  tweet.AuthorId,
		Message:   tweet.Message,
		Upvotes:   tweet.Upvotes,
		CreatedAt: timestamppb.New(tweet.CreatedAt.Time()),
		UpdatedAt: timestamppb.New(tweet.UpdatedAt.Time()),
	}
}
