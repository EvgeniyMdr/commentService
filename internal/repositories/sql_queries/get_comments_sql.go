package sql_queries

const GetComments = `
	SELECT *
			FROM comments
			WHERE parent_id IS NULL AND post_id = $1
			LIMIT $2 OFFSET $3;
`
const GetTotalCommentsCount = `
    SELECT COUNT(*)
    FROM comments
    WHERE parent_id IS NULL AND post_id = $1;
`
