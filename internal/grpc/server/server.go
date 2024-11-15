package server

import (
	"context"
	"github.com/EvgeniyMdr/commentService/internal/services"
	commentsv1 "github.com/EvgeniyMdr/protos/gen/go/comments"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type serverAPI struct {
	commentsv1.UnimplementedCommentsServer
	service services.Service
}

func Register(gRPC *grpc.Server, service services.Service) {
	commentsv1.RegisterCommentsServer(gRPC, &serverAPI{
		service: service,
	})
}
func (s *serverAPI) CreateComment(ctx context.Context, req *commentsv1.CreateCommentDto) (*commentsv1.CommentDto, error) {
	comment := &commentsv1.CommentDto{
		Id:        "123123",
		PostId:    "123123",
		AuthorId:  "123123",
		Content:   "123123",
		CreatedAt: timestamppb.Now(),
		UpdatedAt: timestamppb.Now(),
	}

	return comment, nil
}

func (s *serverAPI) GetComments(ctx context.Context, req *commentsv1.GetCommentsDto) (*commentsv1.CommentsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetComments not implemented")
}
func (s *serverAPI) GetChildComments(ctx context.Context, req *commentsv1.GetChildCommentsDto) (*commentsv1.CommentsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetChildComments not implemented")
}
func (s *serverAPI) DeleteComment(ctx context.Context, req *commentsv1.DeleteCommentDto) (*commentsv1.DeleteCommentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteComment not implemented")
}
func (s *serverAPI) UpdateComment(ctx context.Context, req *commentsv1.UpdateCommentDto) (*commentsv1.CommentDto, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateComment not implemented")
}
