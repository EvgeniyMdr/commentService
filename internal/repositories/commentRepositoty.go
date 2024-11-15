package repositories

import (
	"context"
	"github.com/google/uuid"
)

type Repo interface {
	CreateComment(ctx context.Context) error
	GetComments(ctx context.Context, id uuid.UUID) error
	GetChildComments(ctx context.Context) error
	DeleteComment(ctx context.Context, id uuid.UUID) (*bool, error)
	UpdateComment(ctx context.Context, id uuid.UUID) error
}
