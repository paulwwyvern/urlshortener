package userstorage

import (
	"context"
	"math/rand"
)

type Storage struct {
}

func NewStorage() *Storage {
	return &Storage{}
}

func (s *Storage) CreateUser(_ context.Context) (int32, error) {
	userId := rand.Int31()
	for userId == 0 {
		userId = rand.Int31()
	}

	return userId, nil
}
