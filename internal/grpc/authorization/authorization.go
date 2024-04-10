package authorization

import (
	"context"
	"fmt"
	"github.com/golang/protobuf/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"test-gRPC/entity"
	ssov1 "test-gRPC/protobuf"
)

type Authorization interface {
	CreateUser(ctx context.Context, user entity.User) (int64, error)
	GenerateToken(ctx context.Context, email, password string) (string, error)
}

type serverAPI struct {
	ssov1.UnimplementedAuthorizationServer
	auth Authorization
}

func Register(gRPC *grpc.Server, auth Authorization) {
	ssov1.RegisterAuthorizationServer(gRPC, &serverAPI{auth: auth})
}

func (s *serverAPI) SignUp(ctx context.Context, req *ssov1.SignUpRequest) (*ssov1.SignUpResponse, error) {
	var input entity.User

	data, err := proto.Marshal(req)
	if err != nil {
		return nil, err
	}
	err = proto.Unmarshal(data, &input)
	if err != nil {
		return nil, err
	}

	id, err := s.auth.CreateUser(ctx, input)
	if err != nil {
		return nil, status.Error(codes.Internal, "internal error")
	}

	return &ssov1.SignUpResponse{
		UserId: id,
	}, nil
}

type signInInput struct {
	Email    string `protobuf:"bytes,1,opt,name=email"`
	Password string `protobuf:"bytes,1,opt,name=password"`
}

func (i *signInInput) Reset()         { *i = signInInput{} }
func (i *signInInput) String() string { return fmt.Sprintf("%+v", *i) }
func (i *signInInput) ProtoMessage()  {}

func (s *serverAPI) SignIn(ctx context.Context, req *ssov1.SignInRequest) (*ssov1.SignInResponse, error) {
	var input signInInput

	data, err := proto.Marshal(req)
	if err != nil {
		return nil, err
	}
	err = proto.Unmarshal(data, &input)
	if err != nil {
		return nil, err
	}

	token, err := s.auth.GenerateToken(ctx, input.Email, input.Password)
	if err != nil {
		return nil, status.Error(codes.Internal, "internal error")
	}

	return &ssov1.SignInResponse{
		Token: token,
	}, nil
}
