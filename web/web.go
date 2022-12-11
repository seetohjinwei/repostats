package web

import (
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

func Start(pool *pgxpool.Pool) {
	gin.SetMode(gin.ReleaseMode)
	router := gin.New()
	router.Use(gin.Recovery(), gin.Logger())

	router.GET("/repo", GetRepo(pool))
	router.GET("/repo_force", ForceGetRepo(pool))

	router.Run(":8083")
}
