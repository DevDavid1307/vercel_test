package api_proto

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type DemoHTTPServer interface {
	Greeter(string) (string, error)
}

func RegisterDemoService(g *gin.Engine, srv DemoHTTPServer) {
	group := g.Group("/api/demo")

	greeter_HTTP_Handler(group, srv)
}

func greeter_HTTP_Handler(g *gin.RouterGroup, srv DemoHTTPServer) {
	g.GET("/index", func(c *gin.Context) {
		resp, err := srv.Greeter("david")
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"err": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": resp})
	})
}
