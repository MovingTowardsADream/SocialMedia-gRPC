package service

import (
	"context"
	"crypto/sha1"
	"fmt"
	"log/slog"
	"test-gRPC/entity"
	"time"
)

const (
	salt = "f9uuruefje3"
)

type Auth struct {
	log      *slog.Logger
	usrProc  UserProcedure
	tokenTLL time.Duration
}

type UserProcedure interface {
	CreateUser(ctx context.Context, user entity.User) (int64, error)
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
	user.Password = generatePasswordHash(user.Password)
	return a.usrProc.CreateUser(ctx, user)
}

func (a *Auth) GenerateToken(ctx context.Context, email, password string) (string, error) {
	return "", nil
}

func generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))
	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
