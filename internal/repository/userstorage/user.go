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
	userID := rand.Int31()
	for userID == 0 {
		userID = rand.Int31()
	}

	return userID, nil
}
