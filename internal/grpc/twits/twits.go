package twits

import (
	"context"
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/golang/protobuf/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"strconv"
	"strings"
	ssov1 "test-gRPC/protobuf"
)

type ListTwit interface {
	CreateTwit(ctx context.Context, twit ssov1.CreateTwitRequest, userId int) (int64, error)
	GetTwit(ctx context.Context, twitId int64, userId int) (string, error)
	DeleteTwit(ctx context.Context, twitId int64, userId int) error
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
	var input ssov1.CreateTwitRequest

	data, err := proto.Marshal(req)
	if err != nil {
		return nil, err
	}
	err = proto.Unmarshal(data, &input)
	if err != nil {
		return nil, err
	}

	id, err := s.listTwit.CreateTwit(ctx, input, userId)
	if err != nil {
		return nil, status.Error(codes.Internal, "internal error")
	}
	return &ssov1.Message{
		Message: "Successfully added. Id: " + strconv.Itoa(int(id)),
	}, nil
}

func (s *serverAPI) GetTwit(ctx context.Context, req *ssov1.GetTwitRequest) (*ssov1.Message, error) {
	userId, err := userIdentity(ctx)
	if err != nil {
		return nil, err
	}

	l, err := s.listTwit.GetTwit(ctx, req.TwitId, userId)
	if err != nil {
		return nil, status.Error(codes.Internal, "internal error")
	}

	return &ssov1.Message{
		Message: l,
	}, nil
}

func (s *serverAPI) DeleteTwit(ctx context.Context, req *ssov1.DeleteTwitRequest) (*ssov1.Message, error) {
	userId, err := userIdentity(ctx)
	if err != nil {
		return nil, err
	}

	err = s.listTwit.DeleteTwit(ctx, req.TwitId, userId)
	if err != nil {
		return nil, status.Error(codes.Internal, "internal error")
	}

	return &ssov1.Message{
		Message: "Successfully delete",
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
		return 0, err
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
