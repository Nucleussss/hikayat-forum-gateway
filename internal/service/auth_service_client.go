package service

import (
	"context"
	"net"
	"time"

	authpb "github.com/Nucleussss/hikayat-proto/gen/go/auth/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type AuthServiceClient struct {
	client authpb.AuthServiceClient
}

func NewAuthServiceClient(addr string) (*AuthServiceClient, error) {
	conn, err := grpc.NewClient(addr,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithContextDialer(func(ctx context.Context, addr string) (net.Conn, error) {
			d := net.Dialer{}
			return d.DialContext(ctx, "tcp", addr)
		}),
	)

	if err != nil {
		return nil, err
	}

	return &AuthServiceClient{client: authpb.NewAuthServiceClient(conn)}, nil
}

func (s *AuthServiceClient) Login(ctx context.Context, req *authpb.LoginRequest) (*authpb.LoginResponse, error) {
	ctxTO, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()
	return s.client.Login(ctxTO, req)
}

func (s *AuthServiceClient) Register(ctx context.Context, req *authpb.RegisterRequest) (*authpb.RegisterResponse, error) {
	ctxTO, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()
	return s.client.Register(ctxTO, req)
}

func (s *AuthServiceClient) GetUser(ctx context.Context, req *authpb.GetUserRequest) (*authpb.User, error) {
	ctxTO, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()
	return s.client.GetUser(ctxTO, req)
}

func (s *AuthServiceClient) UpdateUserProfile(ctx context.Context, req *authpb.UpdateUserProfileRequest) (*authpb.UpdateUserProfileResponse, error) {
	ctxTO, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()
	return s.client.UpdateUserProfile(ctxTO, req)
}

func (s *AuthServiceClient) ChangeUserEmail(ctx context.Context, req *authpb.ChangeUserEmailRequest) (*authpb.ChangeUserEmailResponse, error) {
	ctxTO, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()
	return s.client.ChangeUserEmail(ctxTO, req)
}

func (s *AuthServiceClient) ChangeUserPassword(ctx context.Context, req *authpb.ChangeUserPasswordRequest) (*authpb.ChangeUserPasswordResponse, error) {
	ctxTO, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()
	return s.client.ChangeUserPassword(ctxTO, req)
}

func (s *AuthServiceClient) DeleteUser(ctx context.Context, req *authpb.DeleteUserRequest) (*authpb.DeleteUserResponse, error) {
	ctxTO, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()
	return s.client.DeleteUser(ctxTO, req)
}
