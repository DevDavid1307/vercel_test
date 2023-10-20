package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Demo(w http.ResponseWriter, r *http.Request) {
	serve := gin.New()
	group := serve.Group("/api/demo")

	group.GET("/index", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Hello from demo endpoint!"})
	})

	serve.ServeHTTP(w, r)
}
