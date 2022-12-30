package web

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/seetohjinwei/repostats/data"
	"github.com/seetohjinwei/repostats/image"
)

// Returns username, ok from ctx.
func checkUserQueries(ctx *gin.Context) (string, bool) {
	username := ctx.DefaultQuery("username", "")

	if username == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "username missing"})
		return username, false
	}

	return username, true
}

// Returns username, repo, ok from ctx.
func checkRepoQueries(ctx *gin.Context) (string, string, bool) {
	username := ctx.DefaultQuery("username", "")
	repo := ctx.DefaultQuery("repo", "")

	if username == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "username missing"})
		return username, repo, false
	}
	if repo == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "repo missing"})
		return username, repo, false
	}

	return username, repo, true
}

var ErrBadRequest = errors.New("something went wrong!")

func GetUser(pool *pgxpool.Pool) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		username, ok := checkUserQueries(ctx)
		if !ok {
			return
		}

		repos, err := data.QueryUser(pool, username)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": ErrBadRequest})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{"data": repos})
	}
}

func ForceGetUser(pool *pgxpool.Pool) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		username, ok := checkUserQueries(ctx)
		if !ok {
			return
		}

		repos, err := data.ForceQueryUser(pool, username)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": ErrBadRequest})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{"data": repos})
	}
}

func GetRepo(pool *pgxpool.Pool) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		username, repo, ok := checkRepoQueries(ctx)
		if !ok {
			return
		}

		typeData, err := data.QueryRepository(pool, username, repo)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": ErrBadRequest})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{"data": typeData})
	}
}

func ForceGetRepo(pool *pgxpool.Pool) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		username, repo, ok := checkRepoQueries(ctx)
		if !ok {
			return
		}

		typeData, err := data.ForceQueryRepository(pool, username, repo)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": ErrBadRequest})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{"data": typeData})
	}
}

func GetUserImage(pool *pgxpool.Pool) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		username, repo, ok := checkRepoQueries(ctx)
		if !ok {
			return
		}

		typeData, err := data.QueryRepository(pool, username, repo)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": ErrBadRequest})
			return
		}

		w := ctx.Writer
		err = image.CreateUserSvg(w, username, typeData)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": ErrBadRequest})
			return
		}

		ctx.Header("Content-Type", "image/svg+xml")
		ctx.Status(http.StatusOK)
	}
}
