package service

import (
	"context"
	"log/slog"
	ssov1 "test-gRPC/protobuf"
)

type ListTwit struct {
	log     *slog.Logger
	twtProc TwitProcedure
}

type TwitProcedure interface {
	CreateTwit(ctx context.Context, twit ssov1.CreateTwitRequest, userId int) (int64, error)
	GetTwit(ctx context.Context, twitId int64, userId int) (string, error)
	DeleteTwit(ctx context.Context, twitId int64, userId int) error
}

func NewListTwit(log *slog.Logger, twtProc TwitProcedure) *ListTwit {
	return &ListTwit{
		log:     log,
		twtProc: twtProc,
	}
}

func (a *ListTwit) CreateTwit(ctx context.Context, twit ssov1.CreateTwitRequest, userId int) (int64, error) {
	return a.twtProc.CreateTwit(ctx, twit, userId)
}
func (a *ListTwit) GetTwit(ctx context.Context, twitId int64, userId int) (string, error) {
	return a.twtProc.GetTwit(ctx, twitId, userId)
}
func (a *ListTwit) DeleteTwit(ctx context.Context, twitId int64, userId int) error {
	return a.twtProc.DeleteTwit(ctx, twitId, userId)
}
