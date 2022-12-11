package web

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/seetohjinwei/repostats/data"
)

// Returns username, repo, ok from ctx.
func checkQueries(ctx *gin.Context) (string, string, bool) {
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

func GetRepo(pool *pgxpool.Pool) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		username, repo, ok := checkQueries(ctx)
		if !ok {
			return
		}

		typeData, err := data.QueryRepository(pool, username, repo)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{"data": typeData})
	}
}

func ForceGetRepo(pool *pgxpool.Pool) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		username, repo, ok := checkQueries(ctx)
		if !ok {
			return
		}

		typeData, err := data.ForceQueryRepository(pool, username, repo)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{"data": typeData})
	}
}
