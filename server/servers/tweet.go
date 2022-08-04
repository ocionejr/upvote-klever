package servers

import (
	pb "github.com/ocionejr/upvote-klever/pb"
	"github.com/ocionejr/upvote-klever/server/repositories"
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
