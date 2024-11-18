package repositories

import (
	"context"
	"database/sql"
	commentsv1 "github.com/EvgeniyMdr/protos/gen/go/comments"
	"sync"
)

var once sync.Once
var instance Repo

type commentsRepository struct {
	db *sql.DB
}

func (c commentsRepository) CreateComment(ctx context.Context, comDto commentsv1.CommentDto) (commentsv1.CommentDto, error) {
	//TODO implement me
	panic("implement me")
}

func (c commentsRepository) GetComments(ctx context.Context, req commentsv1.GetCommentsDto) ([]commentsv1.CommentDto, error) {
	//TODO implement me
	panic("implement me")
}

func (c commentsRepository) GetChildComments(ctx context.Context, req commentsv1.GetChildCommentsDto) ([]commentsv1.CommentDto, error) {
	//TODO implement me
	panic("implement me")
}

func (c commentsRepository) DeleteComment(ctx context.Context, dto commentsv1.DeleteCommentDto) (*bool, error) {
	//TODO implement me
	panic("implement me")
}

func (c commentsRepository) UpdateComment(ctx context.Context, updDto commentsv1.UpdateCommentDto) (commentsv1.CommentDto, error) {
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
