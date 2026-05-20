package user

import (
	"context"
	"go.uber.org/zap"
)

type UserRepository interface {
	CreateUser(ctx context.Context) (int32, error)
}

type UserService struct {
	logger *zap.Logger

	userRepo UserRepository
}

func NewService(logger *zap.Logger, userRepo UserRepository) *UserService {

	logger.Info("Created user service")

	return &UserService{
		logger:   logger,
		userRepo: userRepo,
	}
}

func (s *UserService) CreateUser(ctx context.Context) (int32, error) {
	userId, err := s.userRepo.CreateUser(ctx)
	if err != nil {
		return 0, err
	}
	s.logger.Info("Created user", zap.Int32("userId", userId))
	return userId, nil
}
