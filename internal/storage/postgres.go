package postgres

import (
	"context"
	"database/sql"
	"fmt"
	"test-gRPC/entity"
)

type Storage struct {
	db *sql.DB
}

func New(storagePath string) (*Storage, error) {
	const op = "storage.sqlite.New"

	db, err := sql.Open("sqlite3", storagePath)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return &Storage{db: db}, nil
}

func (s *Storage) CreateUser(ctx context.Context, email string, passHash []byte) (int64, error) {
	return 1, nil
}
func (s *Storage) GetUser(ctx context.Context, email, password string) (entity.User, error) {
	return entity.User{}, nil
}

func (s *Storage) CreateTwit(ctx context.Context, twit entity.Twit) (int64, error) {
	return 1, nil
}
func (s *Storage) GetTwit(ctx context.Context, email, password string) (string, error) {
	return "", nil
}
func (s *Storage) DeleteTwit(ctx context.Context, email, password string) (string, error) {
	return "", nil
}
