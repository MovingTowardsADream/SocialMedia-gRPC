package service

import (
	"context"
	"crypto/sha1"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"log/slog"
	"test-gRPC/entity"
	"time"
)

const (
	salt       = "f9uuruefje3"
	signingKey = "opofpajdskvisvieorfd"
	tokenTTL   = 12 * time.Hour
)

type tokenClaims struct {
	jwt.StandardClaims
	UserId int `json:"user_id"`
}

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
	user, err := a.usrProc.GetUser(ctx, email, generatePasswordHash(password))
	if err != nil {
		return "", err
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(tokenTTL).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		int(user.Id),
	})
	return token.SignedString([]byte(signingKey))
}

func generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))
	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
