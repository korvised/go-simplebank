package gapi

import (
	db "github.com/korvised/go-simplebank/db/sqlc"
	"github.com/korvised/go-simplebank/pb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func convertUser(user db.User) *pb.User {
	return &pb.User{
		FullName:          user.FullName,
		Email:             user.Email,
		Username:          user.Username,
		PasswordChangedAt: timestamppb.New(user.PasswordChangedAt),
		CratedAt:          timestamppb.New(user.CreatedAt),
	}
}
