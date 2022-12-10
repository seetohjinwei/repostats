package postgres

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/seetohjinwei/repostats/models"
)

// These functions all silently fail.

func AddUser(pool *pgxpool.Pool, username string) {
	sql := "INSERT INTO Users VALUES ($1);"
	pool.Exec(context.Background(), sql, username)
}

func AddRepository(pool *pgxpool.Pool, username, repo, default_branch string) {
	sql := "CALL add_repo($1, $2, $3);"
	pool.Exec(context.Background(), sql, username, repo, default_branch)
}

// Make sure to add the repository first.
func AddTypeData(pool *pgxpool.Pool, username, repo string, typeData []models.TypeData) int64 {
	rows := make([][]interface{}, len(typeData))

	for i, d := range typeData {
		rows[i] = []interface{}{username, repo, d.Type, d.FileCount, d.Bytes}
	}

	count, _ := pool.CopyFrom(
		context.Background(),
		pgx.Identifier{"typedata"},
		[]string{"username", "repo", "language", "file_count", "bytes"},
		pgx.CopyFromRows(rows),
	)

	return count
}
