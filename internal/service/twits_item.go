package service

import (
	"context"
	"log/slog"
	"test-gRPC/entity"
)

type ListTwit struct {
	log     *slog.Logger
	twtProc TwitProcedure
}

type TwitProcedure interface {
	CreateTwit(ctx context.Context, twit entity.Twit) (int64, error)
	GetTwit(ctx context.Context, email, password string) (string, error)
	DeleteTwit(ctx context.Context, email, password string) (string, error)
}

func NewListTwit(log *slog.Logger, twtProc TwitProcedure) *ListTwit {
	return &ListTwit{
		log:     log,
		twtProc: twtProc,
	}
}

func (a *ListTwit) CreateTwit(ctx context.Context, twit entity.Twit) (int64, error) {
	return 1, nil
}
func (a *ListTwit) GetTwit(ctx context.Context, email, password string) (string, error) {
	return "", nil
}
func (a *ListTwit) DeleteTwit(ctx context.Context, email, password string) (string, error) {
	return "", nil
}
