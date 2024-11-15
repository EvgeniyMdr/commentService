package services

import (
	"context"
	"github.com/google/uuid"
)

type Service interface {
	CreateComment(ctx context.Context, id uuid.UUID) error
	GetComments(ctx context.Context) error
	GetChildComments(ctx context.Context, id uuid.UUID) (*bool, error)
	DeleteComment(ctx context.Context, id uuid.UUID) (*bool, error)
	UpdateComment(ctx context.Context, id uuid.UUID) (*bool, error)
}
