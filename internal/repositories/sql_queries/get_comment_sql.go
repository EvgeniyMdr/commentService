package sql_queries

const GetComment = `
	SELECT id, post_id, author_id, content, created_at, updated_at, parent_id, child_count 
	FROM comments
	WHERE id = $1
`
