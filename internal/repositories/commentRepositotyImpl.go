package repositories

import (
	"context"
	"database/sql"
	"errors"
	"github.com/EvgeniyMdr/commentService/internal/repositories/sql_queries"
	commentsv1 "github.com/EvgeniyMdr/protos/gen/go/comments"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"strconv"
	"sync"
	"time"
)

var once sync.Once
var instance Repo

type commentsRepository struct {
	db *sql.DB
}

func (c commentsRepository) CreateComment(ctx context.Context, comDto *commentsv1.CommentDto) (*commentsv1.CommentDto, error) {
	ctx, cancel := context.WithTimeout(ctx, 4*time.Second)
	defer cancel()

	childCountStr := strconv.Itoa(int(comDto.ChildCount))

	_, err := c.db.ExecContext(ctx, sql_queries.CreateComment, comDto.Id, comDto.PostId, comDto.AuthorId, comDto.Content, comDto.ParentId, childCountStr)

	if err != nil {
		if errors.Is(ctx.Err(), context.DeadlineExceeded) {
			return nil, status.Errorf(codes.DeadlineExceeded, "таймаут при создании комментария: %v", err)
		}

		return nil, status.Errorf(codes.Internal, "ошибка при создании комментария: %v", err)
	}

	return comDto, nil
}

func (c commentsRepository) GetComments(ctx context.Context, req *commentsv1.GetCommentsDto) (*commentsv1.CommentsResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (c commentsRepository) GetChildComments(ctx context.Context, req *commentsv1.GetChildCommentsDto) (*commentsv1.CommentsResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (c commentsRepository) DeleteComment(ctx context.Context, dto *commentsv1.DeleteCommentDto) (*commentsv1.DeleteCommentResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (c commentsRepository) UpdateComment(ctx context.Context, updDto *commentsv1.UpdateCommentDto) (*commentsv1.CommentDto, error) {
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
