package service

import (
	"context"
	"net"
	"time"

	postpb "github.com/Nucleussss/hikayat-proto/gen/go/post/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/types/known/emptypb"
)

type PostServiceClient struct {
	client postpb.PostServiceClient
}

func NewPostServiceClient(addr string) (*PostServiceClient, error) {
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

	return &PostServiceClient{client: postpb.NewPostServiceClient(conn)}, nil
}

func (s *PostServiceClient) CreatePost(ctx context.Context, post *postpb.CreatePostRequest) (*postpb.Post, error) {
	ctxTO, cencel := context.WithTimeout(ctx, 3*time.Second)
	defer cencel()
	return s.client.CreatePost(ctxTO, post)
}

func (s *PostServiceClient) GetPost(ctx context.Context, req *postpb.GetPostRequest) (*postpb.Post, error) {
	ctxTO, cencel := context.WithTimeout(ctx, 3*time.Second)
	defer cencel()
	return s.client.GetPost(ctxTO, req)
}

func (s *PostServiceClient) ListPost(ctx context.Context, req *postpb.ListPostsRequest) (*postpb.ListPostsResponse, error) {
	ctxTO, cencel := context.WithTimeout(ctx, 3*time.Second)
	defer cencel()
	return s.client.ListPosts(ctxTO, req)
}

func (s *PostServiceClient) UpdatePost(ctx context.Context, req *postpb.UpdatePostRequest) (*postpb.Post, error) {
	ctxTO, cencel := context.WithTimeout(ctx, 3*time.Second)
	defer cencel()
	return s.client.UpdatePost(ctxTO, req)
}

func (s *PostServiceClient) DeletePost(ctx context.Context, req *postpb.DeletePostRequest) (*emptypb.Empty, error) {
	ctxTO, cencel := context.WithTimeout(ctx, 3*time.Second)
	defer cencel()
	return s.client.DeletePost(ctxTO, req)
}
