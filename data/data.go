package data

import (
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/seetohjinwei/repostats/models"
	"golang.org/x/exp/maps"
)

// There is a rate-limit of 5000 with my personal access token.
// Without the token, the rate-limit is only 60.
const CACHE_DURATION time.Duration = time.Hour * 1

// Note: When changing CACHE_DURATION this, be sure to update data.cleanDatabase too.

// Queries user, with potentially cached result.
func QueryUser(pool *pgxpool.Pool, username string) ([]models.PSQLRepository, error) {
	last_updated, err := queryUserLastUpdated(pool, username)

	if err == nil && last_updated.Add(CACHE_DURATION).After(time.Now()) {
		// Use cached result.
		return queryCachedUser(pool, username)
	}

	return ForceQueryUser(pool, username)
}

// Forcefully queries repository from GitHub API.
func ForceQueryUser(pool *pgxpool.Pool, username string) ([]models.PSQLRepository, error) {
	repos, err := getRepos(username)
	if err != nil {
		return nil, err
	}

	updateRepositories(pool, username, repos)

	return repos, nil
}

// Queries repository, with potentially cached result.
func QueryRepository(pool *pgxpool.Pool, username, repo string) (map[string]models.TypeData, error) {
	last_updated, err := queryRepositoryLastUpdated(pool, username, repo)

	if err == nil && last_updated.Add(CACHE_DURATION).After(time.Now()) {
		// Use cached result.
		return queryCachedRepository(pool, username, repo)
	}

	return ForceQueryRepository(pool, username, repo)
}

// Forcefully queries repository from GitHub API.
func ForceQueryRepository(pool *pgxpool.Pool, username, repo string) (map[string]models.TypeData, error) {
	branch, err := getDefaultBranch(username, repo)
	if err != nil {
		return nil, err
	}

	typeData, err := getRepositoryWithData(username, repo, branch)
	if err != nil {
		return nil, err
	}

	updateTypeData(pool, username, repo, branch, maps.Values(typeData))

	return typeData, nil
}

// Cleans outdated data from database.
func CleanDatabase(pool *pgxpool.Pool) (CleanedData, error) {
	cleanedData, err := cleanDatabase(pool)
	if err != nil {
		return cleanedData, err
	}

	return cleanedData, nil
}
