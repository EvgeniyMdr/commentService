package repositories

import (
	"context"
	"database/sql"
	"github.com/google/uuid"
	"sync"
)

var once sync.Once
var instance Repo

type commentsRepository struct {
	db *sql.DB
}

func (c commentsRepository) CreateComment(ctx context.Context) error {
	//TODO implement me
	panic("implement me")
}

func (c commentsRepository) GetComments(ctx context.Context, id uuid.UUID) error {
	//TODO implement me
	panic("implement me")
}

func (c commentsRepository) GetChildComments(ctx context.Context) error {
	//TODO implement me
	panic("implement me")
}

func (c commentsRepository) DeleteComment(ctx context.Context, id uuid.UUID) (*bool, error) {
	//TODO implement me
	panic("implement me")
}

func (c commentsRepository) UpdateComment(ctx context.Context, id uuid.UUID) error {
	//TODO implement me
	panic("implement me")
}

func NewCommentsRepository(db *sql.DB) Repo {
	once.Do(func() {
		instance = &commentsRepository{
			db: db,
		}
	})

	return instance
}
