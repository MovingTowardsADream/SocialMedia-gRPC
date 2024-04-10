package postgres

import (
	"context"
	"test-gRPC/entity"
)

func (s *Storage) CreateTwit(ctx context.Context, twit entity.Twit) (int64, error) {
	return 1, nil
}
func (s *Storage) GetTwit(ctx context.Context, email, password string) (string, error) {
	return "", nil
}
func (s *Storage) DeleteTwit(ctx context.Context, email, password string) (string, error) {
	return "", nil
}
