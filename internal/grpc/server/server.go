package server

import (
	"context"
	"github.com/EvgeniyMdr/commentService/internal/services"
	commentsv1 "github.com/EvgeniyMdr/protos/gen/go/comments"
	"google.golang.org/grpc"
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
	comment, err := s.service.CreateComment(ctx, req)

	if err != nil {
		return nil, err
	}

	return comment, nil
}

func (s *serverAPI) GetComments(ctx context.Context, req *commentsv1.GetCommentsDto) (*commentsv1.CommentsResponse, error) {
	comments, err := s.service.GetComments(ctx, req)

	if err != nil {
		return nil, err
	}

	return comments, nil
}
func (s *serverAPI) GetChildComments(ctx context.Context, req *commentsv1.GetChildCommentsDto) (*commentsv1.CommentsResponse, error) {
	comments, err := s.service.GetChildComments(ctx, req)

	if err != nil {
		return nil, err
	}

	return comments, nil
}
func (s *serverAPI) DeleteComment(ctx context.Context, req *commentsv1.DeleteCommentDto) (*commentsv1.DeleteCommentResponse, error) {
	comment, err := s.service.DeleteComment(ctx, req)

	if err != nil {
		return nil, err
	}

	return comment, nil
}
func (s *serverAPI) UpdateComment(ctx context.Context, req *commentsv1.UpdateCommentDto) (*commentsv1.CommentDto, error) {
	comment, err := s.service.UpdateComment(ctx, req)

	if err != nil {
		return nil, err
	}

	return comment, nil
}
