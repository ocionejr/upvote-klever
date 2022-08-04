package models

import (
	"time"

	pb "github.com/ocionejr/upvote-klever/pb"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/protobuf/types/known/timestamppb"
	"gopkg.in/validator.v2"
)

type Tweet struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	AuthorId  string             `bson:"author_id" json:"author_id" validate:"nonzero"`
	Message   string             `bson:"message" json:"message" validate:"max=280"`
	Upvotes   []string           `bson:"upvotes,omitempty" json:"upvotes,omitempty"`
	CreatedAt primitive.DateTime `bson:"created_at" json:"created_at"`
	UpdatedAt primitive.DateTime `bson:"updated_at" json:"updated_at"`
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

func (tweet *Tweet) Validate() error {
	if err := validator.Validate(tweet); err != nil {
		return err
	}
	return nil
}
