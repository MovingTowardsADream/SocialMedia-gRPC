package authorization

import (
	"context"
	"google.golang.org/grpc"
	ssov1 "test-gRPC/protobuf"
)

type serverAPI struct {
	ssov1.UnimplementedAuthorizationServer
}

func Register(gRPC *grpc.Server) {
	ssov1.RegisterAuthorizationServer(gRPC, &serverAPI{})
}

func (s *serverAPI) SignUp(ctx context.Context, req *ssov1.SignUpRequest) (*ssov1.SignUpResponse, error) {

	return &ssov1.SignUpResponse{
		UserId: 1,
	}, nil
}

func (s *serverAPI) SignIn(ctx context.Context, req *ssov1.SignInRequest) (*ssov1.SignInResponse, error) {
	return &ssov1.SignInResponse{
		Token: "dsd",
	}, nil
}
