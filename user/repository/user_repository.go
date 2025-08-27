package repository

import (
	"context"
)

type UserRepository interface {
	//User methods
	GetUserInfo(ctx context.Context, userID int) (*UserInfoResponse, error)

	//add user's model
}