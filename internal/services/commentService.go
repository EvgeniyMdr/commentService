package services

import (
	"context"
	commentsv1 "github.com/EvgeniyMdr/protos/gen/go/comments"
)

type Service interface {
	CreateComment(ctx context.Context, req *commentsv1.CreateCommentDto) (*commentsv1.CommentDto, error)
	GetComments(ctx context.Context, req *commentsv1.GetCommentsDto) (*commentsv1.CommentsResponse, error)
	GetChildComments(ctx context.Context, req *commentsv1.GetChildCommentsDto) (*commentsv1.CommentsResponse, error)
	DeleteComment(ctx context.Context, req *commentsv1.DeleteCommentDto) (*commentsv1.DeleteCommentResponse, error)
	UpdateComment(ctx context.Context, req *commentsv1.UpdateCommentDto) (*commentsv1.CommentDto, error)
}
