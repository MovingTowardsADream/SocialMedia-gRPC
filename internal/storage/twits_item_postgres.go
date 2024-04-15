package postgres

import (
	"context"
	"fmt"
	ssov1 "test-gRPC/protobuf"
)

func (s *Storage) CreateTwit(ctx context.Context, twit ssov1.CreateTwitRequest, userId int) (int64, error) {
	var id int64

	createListQuery := fmt.Sprintf("INSERT INTO %s (twit, user_id) VALUES ($1, $2) RETURNING id", twitsTable)
	row := s.db.QueryRow(createListQuery, twit.Twit, userId)

	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}
func (s *Storage) GetTwit(ctx context.Context, twitId int64, userId int) (string, error) {
	var twit string
	query := fmt.Sprintf("SELECT twit FROM %s WHERE id=$1 AND user_id=$2", twitsTable)
	err := s.db.Get(&twit, query, twitId, userId)

	return twit, err
}
func (s *Storage) DeleteTwit(ctx context.Context, twitId int64, userId int) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE id=$1 AND user_id=$2", twitsTable)

	_, err := s.db.Exec(query, twitId, userId)

	return err
}
