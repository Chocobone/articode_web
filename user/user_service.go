package user

import (
	"context"
	"log"

	"github.com/chocobone/articode_web/user/repository"
)

type UserService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) *UserService {
	if repo == nil {
		log.Fatal("User Repository cannot be nil")
	}

	return &UserService{repo: repo}
}

func (s *UserService) GetUserInfo(ctx context.Context, userID int) (*repository.UserInfoResponse, error) {
	return s.repo.GetUserInfo(ctx, userID)
}

func (s *UserService) PostUserInfo(ctx context.Context, user *repository.UserInfoResponse) (*repository.UserInfoResponse, error) {
	return s.repo.PostUserInfo(ctx, user)
}

func (s *UserService) DeleteUserInfo(ctx context.Context, userID int) error {
	return s.repo.DeleteUserInfo(ctx, userID)
}
