package repositories

import (
	"context"
	commentsv1 "github.com/EvgeniyMdr/protos/gen/go/comments"
)

type Repo interface {
	CreateComment(ctx context.Context, comDto *commentsv1.CommentDto) (*commentsv1.CommentDto, error)
	GetComments(ctx context.Context, req *commentsv1.GetCommentsDto) (*commentsv1.CommentsResponse, error)
	GetChildComments(ctx context.Context, req *commentsv1.GetChildCommentsDto) (*commentsv1.CommentsResponse, error)
	DeleteComment(ctx context.Context, dto *commentsv1.DeleteCommentDto) (*commentsv1.DeleteCommentResponse, error)
	UpdateComment(ctx context.Context, updDto *commentsv1.UpdateCommentDto) (*commentsv1.CommentDto, error)
}
