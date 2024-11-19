package services

import (
	"context"
	"github.com/EvgeniyMdr/commentService/internal/repositories"
	commentsv1 "github.com/EvgeniyMdr/protos/gen/go/comments"
	"github.com/google/uuid"
	"sync"
)

var once sync.Once
var instance Service

type commentsService struct {
	repo repositories.Repo
}

func (c commentsService) CreateComment(ctx context.Context, req *commentsv1.CreateCommentDto) (*commentsv1.CommentDto, error) {
	comUuid := uuid.NewString()

	comment := &commentsv1.CommentDto{
		Id:         comUuid,
		PostId:     req.PostId,
		AuthorId:   req.AuthorId,
		Content:    req.Content,
		ParentId:   req.ParentId,
		ChildCount: 0,
	}
	comm, err := c.repo.CreateComment(ctx, comment)

	if err != nil {
		return nil, err
	}

	return comm, err
}

func (c commentsService) GetComments(ctx context.Context, req *commentsv1.GetCommentsDto) (*commentsv1.CommentsResponse, error) {
	comments, err := c.repo.GetComments(ctx, req)

	if err != nil {
		return nil, err
	}

	return comments, err
}

func (c commentsService) GetChildComments(ctx context.Context, req *commentsv1.GetChildCommentsDto) (*commentsv1.CommentsResponse, error) {
	comments, err := c.repo.GetChildComments(ctx, req)

	if err != nil {
		return nil, err
	}

	return comments, err
}

func (c commentsService) DeleteComment(ctx context.Context, req *commentsv1.DeleteCommentDto) (*commentsv1.DeleteCommentResponse, error) {
	res, err := c.repo.DeleteComment(ctx, req)

	if err != nil {
		return nil, err
	}

	return res, err
}

func (c commentsService) UpdateComment(ctx context.Context, req *commentsv1.UpdateCommentDto) (*commentsv1.CommentDto, error) {
	comm, err := c.repo.UpdateComment(ctx, req)

	if err != nil {
		return nil, err
	}

	return comm, err
}

func NewCommentsService(repo repositories.Repo) Service {
	once.Do(func() {
		instance = &commentsService{
			repo: repo,
		}
	})

	return instance
}
