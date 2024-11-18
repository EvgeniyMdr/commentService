package services

import (
	"context"
	"github.com/EvgeniyMdr/commentService/internal/repositories"
	commentsv1 "github.com/EvgeniyMdr/protos/gen/go/comments"
	"sync"
)

var once sync.Once
var instance Service

type commentsService struct {
	repo repositories.Repo
}

func (c commentsService) CreateComment(ctx context.Context, req *commentsv1.CreateCommentDto) (*commentsv1.CommentDto, error) {
	//TODO implement me
	panic("implement me")
}

func (c commentsService) GetComments(ctx context.Context, req *commentsv1.GetCommentsDto) (*commentsv1.CommentsResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (c commentsService) GetChildComments(ctx context.Context, req *commentsv1.GetChildCommentsDto) (*commentsv1.CommentsResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (c commentsService) DeleteComment(ctx context.Context, req *commentsv1.DeleteCommentDto) (*commentsv1.DeleteCommentResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (c commentsService) UpdateComment(ctx context.Context, req *commentsv1.UpdateCommentDto) (*commentsv1.CommentDto, error) {
	//TODO implement me
	panic("implement me")
}

func NewCommentsService(repo repositories.Repo) Service {
	once.Do(func() {
		instance = &commentsService{
			repo: repo,
		}
	})

	return instance
}
