package repositories

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/EvgeniyMdr/commentService/internal/repositories/sql_queries"
	commentsv1 "github.com/EvgeniyMdr/protos/gen/go/comments"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
	"log"
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
	parentID := sql.NullString{}
	if comDto.ParentId != nil && *comDto.ParentId != "" {
		parentID = sql.NullString{
			String: *comDto.ParentId,
			Valid:  true,
		}
	}

	_, err := c.db.ExecContext(ctx, sql_queries.CreateComment, comDto.Id, comDto.PostId, comDto.AuthorId, comDto.Content, parentID, childCountStr)

	if err != nil {
		if errors.Is(ctx.Err(), context.DeadlineExceeded) {
			return nil, status.Errorf(codes.DeadlineExceeded, "таймаут при создании комментария: %v", err)
		}

		return nil, status.Errorf(codes.Internal, "ошибка при создании комментария: %v", err)
	}

	row := c.db.QueryRowContext(ctx, sql_queries.GetComment, comDto.Id)

	var comment commentsv1.CommentDto
	var createdAt time.Time
	var updatedAt time.Time

	err = row.Scan(
		&comment.Id,
		&comment.PostId,
		&comment.AuthorId,
		&comment.Content,
		&createdAt,
		&updatedAt,
		&comment.ParentId,
		&comment.ChildCount,
	)

	if err != nil {
		return nil, fmt.Errorf("scanning comment: %w", err)
	}

	comment.CreatedAt = timestamppb.New(createdAt)
	comment.UpdatedAt = timestamppb.New(updatedAt)

	return &comment, nil
}

func (c commentsRepository) GetComments(ctx context.Context, req *commentsv1.GetCommentsDto) (*commentsv1.CommentsResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, 4*time.Second)
	defer cancel()

	rows, err := c.db.QueryContext(ctx, sql_queries.GetComments, req.PostId, req.Take, req.Skip)

	var comments []*commentsv1.CommentDto
	for rows.Next() {
		var comment commentsv1.CommentDto

		var createdAt time.Time
		var updatedAt time.Time

		err := rows.Scan(
			&comment.Id,
			&comment.PostId,
			&comment.AuthorId,
			&comment.Content,
			&createdAt,
			&updatedAt,
			&comment.ParentId,
			&comment.ChildCount,
		)

		if err != nil {
			return nil, fmt.Errorf("scanning comment: %w", err)
		}

		comment.CreatedAt = timestamppb.New(createdAt)
		comment.UpdatedAt = timestamppb.New(updatedAt)

		comments = append(comments, &comment)
	}

	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			log.Fatalf("cant clone rows clone")
		}
	}(rows)

	if err != nil {
		if errors.Is(ctx.Err(), context.DeadlineExceeded) {
			return nil, status.Errorf(codes.DeadlineExceeded, "таймаут при полученнии комментариев: %v", err)
		}

		return nil, status.Errorf(codes.Internal, "ошибка при получении комментариев: %v", err)
	}

	// Выполняем запрос на подсчет общего количества
	var totalCount int
	err = c.db.QueryRowContext(ctx, sql_queries.GetTotalCommentsCount, req.PostId).Scan(&totalCount)
	if err != nil {
		return nil, fmt.Errorf("counting comments: %w", err)
	}

	return &commentsv1.CommentsResponse{
		Comments: comments,
		Total:    int32(totalCount),
		Skip:     req.Skip,
	}, nil
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
