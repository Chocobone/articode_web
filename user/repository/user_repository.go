package repository

import (
	"context"
)

type UserRepository interface {
	//User methods
	GetUserInfo(ctx context.Context, userID int) (*UserInfoResponse, error)

	//add user's model
	PostUserInfo(ctx context.Context, user *UserInfoResponse) (*UserInfoResponse, error)

	// delete user info
	DeleteUserInfo(ctx context.Context, userID int) error
}
