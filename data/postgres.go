package data

import (
	"context"
	"errors"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/lib/pq"
	"github.com/seetohjinwei/repostats/models"
)

// These functions all silently fail.

// Adds a user.
func addUser(pool *pgxpool.Pool, username string) {
	pool.Exec(context.Background(), `
	INSERT INTO Users VALUES ($1);
	`, username)
}

// Adds a repository.
func addRepository(pool *pgxpool.Pool, username, repo, default_branch string) {
	pool.Exec(context.Background(), `
	CALL add_repo($1, $2, $3);
	`, username, repo, default_branch)
}

// Updates type data for a repository.
// Creates the repository, if it does not exist.
func updateRepositories(pool *pgxpool.Pool, username string, repos []models.PSQLRepository) {
	pool.Exec(context.Background(), `
	CALL update_repos($1, $2);
	`, username, pq.Array(repos))
}

// Gets the last updated value for a user.
func queryUserLastUpdated(pool *pgxpool.Pool, username string) (*time.Time, error) {
	row := pool.QueryRow(context.Background(), `
	SELECT last_updated FROM Users U
	WHERE U.username = $1;
	`, username)
	last_updated := &time.Time{}

	err := row.Scan(&last_updated)

	if last_updated == nil {
		return nil, errors.New("no entry")
	}

	return last_updated, err
}

// Gets the simple repositories for a user.
func queryCachedUser(pool *pgxpool.Pool, username string) ([]models.PSQLRepository, error) {
	repos := []models.PSQLRepository{}

	rows, err := pool.Query(context.Background(), `
	SELECT * FROM Repositories R
	WHERE R.username = $1;
	`, username)
	if err != nil {
		return repos, err
	}
	defer rows.Close()

	for rows.Next() {
		var name string
		var default_branch string

		// nil => ignore
		err := rows.Scan(nil, &name, nil, &default_branch)
		if err != nil {
			return repos, err
		}

		repo := models.PSQLRepository{
			Repo:          name,
			DefaultBranch: default_branch,
		}

		repos = append(repos, repo)
	}

	return repos, nil
}

// Updates type data for a repository.
// Creates the repository, if it does not exist.
func updateTypeData(pool *pgxpool.Pool, username, repo, default_branch string, typeData []models.TypeData) {
	pool.Exec(context.Background(), `
	CALL update_typedata($1, $2, $3, $4);
	`, username, repo, default_branch, pq.Array(typeData))
}

// Gets the last updated value for a repository.
func queryRepositoryLastUpdated(pool *pgxpool.Pool, username, repo string) (*time.Time, error) {
	row := pool.QueryRow(context.Background(), `
	SELECT last_updated FROM Repositories R
	WHERE R.username = $1
		AND R.repo = $2;
	`, username, repo)
	last_updated := &time.Time{}

	err := row.Scan(&last_updated)

	if last_updated == nil {
		return nil, errors.New("no entry")
	}

	return last_updated, err
}

// Gets the type data for a repository.
func queryCachedRepository(pool *pgxpool.Pool, username, repo string) (map[string]models.TypeData, error) {
	typeData := map[string]models.TypeData{}

	rows, err := pool.Query(context.Background(), `
	SELECT * FROM TypeData TD
	WHERE TD.username = $1
		AND TD.repo = $2;
	`, username, repo)
	if err != nil {
		return typeData, err
	}
	defer rows.Close()

	for rows.Next() {
		var language string
		var file_count int64
		var bytes int64

		// nil => ignore
		err := rows.Scan(nil, nil, &language, &file_count, &bytes)
		if err != nil {
			return typeData, err
		}

		typeData[language] = models.TypeData{
			Type:      language,
			FileCount: file_count,
			Bytes:     bytes,
		}
	}

	return typeData, nil
}
