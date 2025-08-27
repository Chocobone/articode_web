package user

import (
	"context"
	"log"

	"github.com/Chocobone/articode_web/v2/user/repository"
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