package api

import (
	"fmt"
	"net/http"
)

func Demo(w http.ResponseWriter, r *http.Request) {
	_, _ = fmt.Fprintf(w, "%s", r.URL.Path)
	//serve := gin.New()
	//
	//serve.GET("/index", func(c *gin.Context) {
	//	c.JSON(http.StatusOK, gin.H{"message": "Hello from demo endpoint!"})
	//})
	//
	//serve.ServeHTTP(w, r)
}
