package sql_queries

const CreateComment = `
	INSERT INTO comments (id, post_id, author_id, content, parent_id, child_count, created_at, updated_at)
	VALUES ($1, $2, $3, $4, $5, $6, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP);
`
