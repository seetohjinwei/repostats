package web

import (
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/robfig/cron/v3"
)

func Start(pool *pgxpool.Pool) {
	c := cron.New()
	c.AddFunc("@every 30m", func() {
		CleanDatabase(pool)
	})
	// Asynchronously invoked in goroutines. So, won't affect Gin.
	c.Start()
	defer c.Stop()

	gin.SetMode(gin.ReleaseMode)
	router := gin.New()
	router.Use(gin.Recovery(), gin.Logger())

	router.GET("/user", GetUser(pool))
	router.GET("/user_force", ForceGetUser(pool))
	router.GET("/repo", GetRepo(pool))
	router.GET("/repo_force", ForceGetRepo(pool))

	router.GET("/repo_image", GetUserImage(pool))

	router.POST("/clean", PostCleanDatabase(pool))

	router.Run(":8083")
}
