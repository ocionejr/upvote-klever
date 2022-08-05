package models

import (
	"time"

	pb "github.com/ocionejr/upvote-klever/pb"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/protobuf/types/known/timestamppb"
	"gopkg.in/validator.v2"
)

type Tweet struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"_id"`
	AuthorId  string             `bson:"author_id,omitempty" json:"author_id" creating:"nonzero"`
	Message   string             `bson:"message" json:"message" creating:"min=5,max=280" updating:"min=5,max=280"`
	Upvotes   []string           `bson:"upvotes,omitempty" json:"upvotes"`
	CreatedAt primitive.DateTime `bson:"created_at,omitempty" json:"created_at"`
	UpdatedAt primitive.DateTime `bson:"updated_at" json:"updated_at"`
}

var (
	creationValidator = validator.NewValidator()
	updateValidator   = validator.NewValidator()
)

func init() {
	creationValidator.SetTag("creating")
	updateValidator.SetTag("updating")
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

func UpdateTweetRequestToTweet(req *pb.UpdateTweetRequest) *Tweet {
	return &Tweet{
		Message:   req.Message,
		UpdatedAt: primitive.NewDateTimeFromTime(time.Now()),
	}
}

func (tweet *Tweet) Validate(isCreating bool) error {
	var err error

	if isCreating {
		err = creationValidator.Validate(tweet)
	} else {
		err = updateValidator.Validate(tweet)
	}

	if err != nil {
		return err
	}

	return nil
}
