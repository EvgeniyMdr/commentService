package migrations

import (
	"context"
	"database/sql"
	"github.com/pressly/goose/v3"
)

func init() {
	goose.AddMigrationContext(upCommentTable, downCommentTable)
}

func upCommentTable(ctx context.Context, tx *sql.Tx) error {
	_, err := tx.ExecContext(ctx, `
		CREATE TABLE comments (
			id VARCHAR(36) PRIMARY KEY,
			post_id VARCHAR(36) NOT NULL,
			author_id VARCHAR(36) NOT NULL,
			content TEXT NOT NULL,
			created_at TIMESTAMP NOT NULL DEFAULT NOW(),
			updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
			parent_id VARCHAR(36),
			child_count INT DEFAULT 0
		)
	`)
	return err
}

func downCommentTable(ctx context.Context, tx *sql.Tx) error {
	_, err := tx.ExecContext(ctx, `DROP TABLE IF EXISTS comments`)
	return err
}
