package services

import (
	"context"
	"github.com/EvgeniyMdr/commentService/internal/repositories"
	"github.com/google/uuid"
	"sync"
)

var once sync.Once
var instance Service

type commentsService struct {
	repo repositories.Repo
}

func (c commentsService) CreateComment(ctx context.Context, id uuid.UUID) error {
	//TODO implement me
	panic("implement me")
}

func (c commentsService) GetComments(ctx context.Context) error {
	//TODO implement me
	panic("implement me")
}

func (c commentsService) GetChildComments(ctx context.Context, id uuid.UUID) (*bool, error) {
	//TODO implement me
	panic("implement me")
}

func (c commentsService) DeleteComment(ctx context.Context, id uuid.UUID) (*bool, error) {
	//TODO implement me
	panic("implement me")
}

func (c commentsService) UpdateComment(ctx context.Context, id uuid.UUID) (*bool, error) {
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
