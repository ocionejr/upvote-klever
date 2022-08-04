package servers

import (
	"context"
	"fmt"

	pb "github.com/ocionejr/upvote-klever/pb"
	"github.com/ocionejr/upvote-klever/server/models"
	"github.com/ocionejr/upvote-klever/server/repositories"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
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

func (s *TweetServer) ListTweets(in *emptypb.Empty, stream pb.TweetService_ListTweetsServer) error {
	cur, err := s.tweetRepository.ListAll()
	if err != nil {
		return err
	}
	defer cur.Close(context.Background())

	for cur.Next(context.Background()) {
		data := &models.Tweet{}
		err := cur.Decode(data)

		if err != nil {
			return status.Errorf(
				codes.Internal,
				fmt.Sprintf("Error while decoding data from MongoDB: %v", err),
			)
		}

		stream.Send(models.TweetToTweetResponse(data))
	}

	if err = cur.Err(); err != nil {
		return status.Errorf(
			codes.Internal,
			fmt.Sprintf("Unknown internal error: %v", err),
		)
	}

	return nil
}
