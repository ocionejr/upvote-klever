package models

import (
	pb "github.com/ocionejr/upvote-klever/pb"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"gopkg.in/validator.v2"
)

type User struct {
	ID primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Username string `bson:"username" json:"username" validate:"regexp=^[A-Za-z][A-Za-z0-9_.-@#]+$,min=5"`
	Password string `bson:"password" json:"password" validate:"regexp=^[A-Za-z][A-Za-z0-9_.-@#]+$,min=5"`
}

func UserToUserRequest(data *User) *pb.UserRequest {
	return &pb.UserRequest{
		Username: data.Username,
		Password: data.Password,
	}
}

func UserRequestToUser(user *pb.UserRequest) *User{
	return &User{
		Username: user.Username,
		Password: user.Password,
	}
}

func (u *User) Validate() error {
	if err := validator.Validate(u); err != nil {
		return err
	}
	return nil
}