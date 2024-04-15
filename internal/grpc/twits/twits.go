package twits

import (
	"context"
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"strings"
	"test-gRPC/entity"
	ssov1 "test-gRPC/protobuf"
)

type ListTwit interface {
	CreateTwit(ctx context.Context, twit entity.Twit) (int64, error)
	GetTwit(ctx context.Context, email, password string) (string, error)
	DeleteTwit(ctx context.Context, email, password string) (string, error)
}

type serverAPI struct {
	ssov1.UnimplementedTwitListServer
	listTwit ListTwit
}

func TwitList(gRPC *grpc.Server, listTwit ListTwit) {
	ssov1.RegisterTwitListServer(gRPC, &serverAPI{listTwit: listTwit})
}

func (s *serverAPI) CreateTwit(ctx context.Context, req *ssov1.CreateTwitRequest) (*ssov1.Message, error) {
	userId, err := userIdentity(ctx)
	if err != nil {
		return nil, err
	}
	fmt.Println(userId)
	return &ssov1.Message{
		Message: string(userId),
	}, nil
}

func (s *serverAPI) GetTwit(ctx context.Context, req *ssov1.GetTwitRequest) (*ssov1.Message, error) {
	return &ssov1.Message{
		Message: "dsds2",
	}, nil
}

func (s *serverAPI) DeleteTwit(ctx context.Context, req *ssov1.DeleteTwitRequest) (*ssov1.Message, error) {
	fmt.Println(int(req.TwitId))
	return &ssov1.Message{
		Message: "dsds3",
	}, nil
}

func userIdentity(ctx context.Context) (int, error) {
	md, _ := metadata.FromIncomingContext(ctx)

	tokenString := md.Get("authorization")
	tokenBearer := strings.Split(tokenString[0], " ")

	if len(tokenBearer) != 2 {
		return 0, errors.New("token empty")
	}
	userId, err := ParseToken(tokenBearer[1])
	if err != nil {
		return 0, errors.New("token empty")
	}
	return userId, nil
}

const signingKey = "opofpajdskvisvieorfd"

type tokenClaims struct {
	jwt.StandardClaims
	UserId int `json:"user_id"`
}

func ParseToken(accessToken string) (int, error) {
	token, err := jwt.ParseWithClaims(accessToken, &tokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}

		return []byte(signingKey), nil
	})
	if err != nil {
		return 0, err
	}

	claims, ok := token.Claims.(*tokenClaims)
	if !ok {
		return 0, errors.New("token claims are not of type *tokenClaims")
	}

	return claims.UserId, nil
}
