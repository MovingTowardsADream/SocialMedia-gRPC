package service

import (
	"context"
	"log/slog"
	"test-gRPC/entity"
	"time"
)

type Auth struct {
	log      *slog.Logger
	usrProc  UserProcedure
	tokenTLL time.Duration
}

type UserProcedure interface {
	CreateUser(ctx context.Context, email string, passHash []byte) (int64, error)
	GetUser(ctx context.Context, email, password string) (entity.User, error)
}

func NewAuth(log *slog.Logger, usrProc UserProcedure, tokenTLL time.Duration) *Auth {
	return &Auth{
		log:      log,
		usrProc:  usrProc,
		tokenTLL: tokenTLL,
	}
}

func (a *Auth) CreateUser(ctx context.Context, user entity.User) (int64, error) {
	return 10, nil
}

func (a *Auth) GenerateToken(ctx context.Context, email, password string) (string, error) {
	return "", nil
}
