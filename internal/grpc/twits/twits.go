package twits

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
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
	return &ssov1.Message{
		Message: "dsds1",
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
