package postgres

import (
	"context"
	"fmt"
	"test-gRPC/entity"
)

func (s *Storage) CreateUser(ctx context.Context, user entity.User) (int64, error) {
	var id int64
	query := fmt.Sprintf("INSERT INTO %s (email, username, password_hash) values($1, $2, $3) RETURNING id", usersTable)
	row := s.db.QueryRow(query, user.Email, user.Username, user.Password)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}
func (s *Storage) GetUser(ctx context.Context, email, password string) (entity.User, error) {
	return entity.User{}, nil
}