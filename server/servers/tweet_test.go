package servers

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	pb "github.com/ocionejr/upvote-klever/pb"
	"github.com/ocionejr/upvote-klever/server/interfaces/mocks"
	"github.com/ocionejr/upvote-klever/server/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
)

func SetupTest(t *testing.T) (*TweetServer, *mocks.MockTweetRepositoryInterface) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	tweetRepositoryMock := mocks.NewMockTweetRepositoryInterface(mockCtrl)
	return &TweetServer{
		tweetRepository: tweetRepositoryMock,
	}, tweetRepositoryMock
}

type TweetService_ListTweetsServerMock struct {
	grpc.ServerStream
}

func (t *TweetService_ListTweetsServerMock) Send(*pb.TweetResponse) error { return nil }

var (
	validTweetMock = &models.Tweet{
		ID:        primitive.NewObjectID(),
		AuthorId:  "12345",
		Message:   "Message",
		Upvotes:   []string{"6789"},
		CreatedAt: primitive.NewDateTimeFromTime(time.Now()),
		UpdatedAt: primitive.NewDateTimeFromTime(time.Now()),
	}

	validTweetMockWithUpvoteOfDefinedUser = &models.Tweet{
		ID:        primitive.NewObjectID(),
		AuthorId:  "12345",
		Message:   "Message",
		Upvotes:   []string{"12345"},
		CreatedAt: primitive.NewDateTimeFromTime(time.Now()),
		UpdatedAt: primitive.NewDateTimeFromTime(time.Now()),
	}

	invalidTweetRequestMock = &pb.TweetRequest{
		AuthorId: "12345",
		Message: `Lorem ipsum dolor sit amet, consectetur adipiscing elit. 
    Maecenas at suscipit nibh. Phasellus quis nulla sed magna feugiat eleifend eget id nisi. 
    Pellentesque id dolor a mauris bibendum rutrum. Ut luctus nisi ac dictum scelerisque. 
    Interdum et malesuada fames ac ante ipsum primis in faucibus turpis.`,
	}

	validTweetRequestMock = &pb.TweetRequest{
		AuthorId: "12345",
		Message:  "Mensagem",
	}

	tweetId = &pb.TweetId{
		Id: primitive.NewObjectID().Hex(),
	}

	invalidUpdateTweetRequestMock = &pb.UpdateTweetRequest{
		Id: primitive.NewObjectID().Hex(),
		Message: `Lorem ipsum dolor sit amet, consectetur adipiscing elit. 
    Maecenas at suscipit nibh. Phasellus quis nulla sed magna feugiat eleifend eget id nisi. 
    Pellentesque id dolor a mauris bibendum rutrum. Ut luctus nisi ac dictum scelerisque. 
    Interdum et malesuada fames ac ante ipsum primis in faucibus turpis.`,
	}

	validUpdateTweetRequestMock = &pb.UpdateTweetRequest{
		Id:      primitive.NewObjectID().Hex(),
		Message: `Message`,
	}

	toggleUpvoteRequestMock = &pb.ToggleUpvoteRequest{
		TweetId: primitive.NewObjectID().Hex(),
		UserId:  "12345",
	}

	listStreamMock = &TweetService_ListTweetsServerMock{}
)

func Test_CreateTweet(t *testing.T) {
	s, r := SetupTest(t)

	t.Run("Should return error when tweet is invalid", func(t *testing.T) {
		res, err := s.CreateTweet(context.Background(), invalidTweetRequestMock)

		if res != nil || err == nil {
			t.Error("Invalid tweet should not be created")
		}
	})

	t.Run("Should return error when repository returns error", func(t *testing.T) {
		r.EXPECT().InsertTweet(gomock.Any(), gomock.Any()).Return(errors.New("erro"))

		res, err := s.CreateTweet(context.Background(), validTweetRequestMock)
		if res != nil || err == nil {
			t.Error("Should not create tweet when repository returns error")
		}
	})

	t.Run("Should create tweet on success", func(t *testing.T) {
		r.EXPECT().InsertTweet(gomock.Any(), gomock.Any()).Return(nil)

		res, err := s.CreateTweet(context.Background(), validTweetRequestMock)
		if res == nil || err != nil {
			t.Error("Should create tweet")
		}
	})
}

func Test_FindTweetById(t *testing.T) {
	s, r := SetupTest(t)

	t.Run("Should return error when repository returns error", func(t *testing.T) {
		r.EXPECT().FindById(gomock.Any(), gomock.Any()).Return(nil, errors.New("erro"))

		res, err := s.FindTweetById(context.Background(), tweetId)
		if res != nil || err == nil {
			t.Error("Should not find tweet when repository returns error")
		}
	})

	t.Run("Should return tweet on success", func(t *testing.T) {
		r.EXPECT().FindById(gomock.Any(), gomock.Any()).Return(validTweetMock, nil)

		res, err := s.FindTweetById(context.Background(), tweetId)
		if res == nil || err != nil {
			t.Error("Should return tweet")
		}
	})
}

func Test_ListTweets(t *testing.T) {
	s, r := SetupTest(t)

	t.Run("Should return error when repository returns error", func(t *testing.T) {
		r.EXPECT().ListAll().Return(nil, errors.New("erro"))

		err := s.ListTweets(&emptypb.Empty{}, listStreamMock)
		if err == nil {
			t.Error("Should not return tweets when repository returns error")
		}
	})

	t.Run("Should return nil when sucess", func(t *testing.T) {
		r.EXPECT().ListAll().Return([]models.Tweet{*validTweetMock}, nil)

		err := s.ListTweets(&emptypb.Empty{}, listStreamMock)
		if err != nil {
			t.Error("Should return nil")
		}
	})
}

func Test_UpdateTweet(t *testing.T) {
	s, r := SetupTest(t)

	t.Run("Should return error when tweet is invalid", func(t *testing.T) {
		res, err := s.UpdateTweet(context.Background(), invalidUpdateTweetRequestMock)

		if res != nil || err == nil {
			t.Error("Invalid tweet should not be updated")
		}
	})

	t.Run("Should return error when repository returns error", func(t *testing.T) {
		r.EXPECT().Update(gomock.Any(), gomock.Any(), gomock.Any()).Return(errors.New("erro"))

		res, err := s.UpdateTweet(context.Background(), validUpdateTweetRequestMock)
		if res != nil || err == nil {
			t.Error("Should not update tweet when repository returns error")
		}
	})

	t.Run("Should return empty when sucess", func(t *testing.T) {
		r.EXPECT().Update(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil)

		res, err := s.UpdateTweet(context.Background(), validUpdateTweetRequestMock)
		if res == nil || err != nil {
			t.Error("Should return empty")
		}
	})
}

func Test_DeleteTweet(t *testing.T) {
	s, r := SetupTest(t)

	t.Run("Should return error when repository returns error", func(t *testing.T) {
		r.EXPECT().DeleteTweet(gomock.Any(), gomock.Any()).Return(errors.New("erro"))

		res, err := s.DeleteTweet(context.Background(), tweetId)
		if res != nil || err == nil {
			t.Error("Should return error when repository returns error")
		}
	})

	t.Run("Should return empty when sucess", func(t *testing.T) {
		r.EXPECT().DeleteTweet(gomock.Any(), gomock.Any()).Return(nil)

		res, err := s.DeleteTweet(context.Background(), tweetId)
		if res == nil || err != nil {
			t.Error("Should return empty")
		}
	})
}

func Test_ToggleUpvote(t *testing.T) {
	s, r := SetupTest(t)

	t.Run("Should return error when repository returns error", func(t *testing.T) {
		r.EXPECT().FindById(gomock.Any(), gomock.Any()).Return(nil, errors.New("erro"))

		res, err := s.ToggleUpvote(context.Background(), toggleUpvoteRequestMock)
		if res != nil || err == nil {
			t.Error("Should return error when repository returns error")
		}
	})

	t.Run("Should return error when remove upvote returns error", func(t *testing.T) {
		r.EXPECT().FindById(gomock.Any(), gomock.Any()).Return(validTweetMockWithUpvoteOfDefinedUser, nil)
		r.EXPECT().RemoveUpvote(gomock.Any(), gomock.Any(), gomock.Any()).Return(errors.New("erro"))

		res, err := s.ToggleUpvote(context.Background(), toggleUpvoteRequestMock)
		if res != nil || err == nil {
			t.Error("Should return error when remove upvote returns error")
		}
	})

	t.Run("Should return error when add upvote returns error", func(t *testing.T) {
		r.EXPECT().FindById(gomock.Any(), gomock.Any()).Return(validTweetMock, nil)
		r.EXPECT().AddUpvote(gomock.Any(), gomock.Any(), gomock.Any()).Return(errors.New("erro"))

		res, err := s.ToggleUpvote(context.Background(), toggleUpvoteRequestMock)
		if res != nil || err == nil {
			t.Error("Should return error when add upvote returns error")
		}
	})

	t.Run("Should return empty when remove upvote returns nil", func(t *testing.T) {
		r.EXPECT().FindById(gomock.Any(), gomock.Any()).Return(validTweetMockWithUpvoteOfDefinedUser, nil)
		r.EXPECT().RemoveUpvote(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil)

		res, err := s.ToggleUpvote(context.Background(), toggleUpvoteRequestMock)
		if res == nil || err != nil {
			t.Error("Should return empty when remove upvote returns nil")
		}
	})

	t.Run("Should return empty when add upvote returns nil", func(t *testing.T) {
		r.EXPECT().FindById(gomock.Any(), gomock.Any()).Return(validTweetMock, nil)
		r.EXPECT().AddUpvote(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil)

		res, err := s.ToggleUpvote(context.Background(), toggleUpvoteRequestMock)
		if res == nil || err != nil {
			t.Error("Should return empty when add upvote returns nil")
		}
	})
}
