package models

import (
	pb "github.com/ocionejr/upvote-klever/pb"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"gopkg.in/validator.v2"
)

type User struct {
	ID primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Username string `bson:"username" json:"username" validate:"regexp=^(?![_.])(?!.*[_.]{2})[a-zA-Z0-9._]+(?<![_.])$,min=5"`
	Password string `bson:"password" json:"password" validate:"regexp=^(?=.*?[A-Z])(?=.*?[a-z])(?=.*?[0-9])(?=.*?[#?!@$%^&*-])$,min=5"`
}

func documentToUserRequest(data *User) *pb.UserRequest {
	return &pb.UserRequest{
		Id: data.ID.Hex(),
		Username: data.Username,
		Password: data.Password,
	}
}

func (u *User) Validate() error {
	if err := validator.Validate(u); err != nil {
		return err
	}
	return nil
}