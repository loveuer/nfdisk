package handler

import (
	"github.com/gin-gonic/gin"
)

func NewHandler() *gin.Engine {
	engine := gin.Default()

	app := engine.Group("/api/nfdisk/")

	{
		api := app.Group("/object")
		api.GET("/download", func(c *gin.Context) {
			id := c.Query("id")
			bucket := c.Query("bucket")
			path := c.Query("path")

			_ = id
			_ = bucket
			_ = path
		})
	}

	return engine
}
