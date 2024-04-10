package postgres

import (
	"context"
	"test-gRPC/entity"
)

func (s *Storage) CreateUser(ctx context.Context, email string, passHash []byte) (int64, error) {
	return 1, nil
}
func (s *Storage) GetUser(ctx context.Context, email, password string) (entity.User, error) {
	return entity.User{}, nil
}
