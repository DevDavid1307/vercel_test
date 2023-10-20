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

	group.GET("/index", _greeter_HTTP_Handler(srv))
}

func _greeter_HTTP_Handler(srv DemoHTTPServer) gin.HandlerFunc {
	return func(c *gin.Context) {
		resp, err := srv.Greeter("david")
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"err": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": resp})
	}
}
