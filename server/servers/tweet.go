package servers

import (
	"context"
	"fmt"

	pb "github.com/ocionejr/upvote-klever/pb"
	"github.com/ocionejr/upvote-klever/server/models"
	"github.com/ocionejr/upvote-klever/server/repositories"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type TweetServer struct {
	pb.UnimplementedTweetServiceServer
	tweetRepository *repositories.TweetRepository
}

func NewTweetServer(tweetRepository *repositories.TweetRepository) *TweetServer {
	return &TweetServer{
		tweetRepository: tweetRepository,
	}
}

func (s *TweetServer) CreateTweet(ctx context.Context, in *pb.TweetRequest) (*pb.TweetResponse, error) {
	tweet := models.TweetRequestToTweet(in)

	if err := tweet.Validate(); err != nil {
		return nil, status.Errorf(
			codes.InvalidArgument,
			fmt.Sprintf("Invalid tweet: %v\n", err),
		)
	}

	if err := s.tweetRepository.InsertTweet(ctx, tweet); err != nil {
		return nil, err
	}

	return models.TweetToTweetResponse(tweet), nil
}

func (s *TweetServer) FindTweetById(ctx context.Context, in *pb.TweetId) (*pb.TweetResponse, error) {
	tweet, err := s.tweetRepository.FindById(ctx, in.Id)

	if err != nil {
		return nil, err
	}

	return models.TweetToTweetResponse(tweet), nil
}
