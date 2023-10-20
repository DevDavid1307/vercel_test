package api_proto

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type DemoHTTPServer interface {
	Greeter(DemoIndexRequest) (DemoIndexResponse, error)
}

type DemoIndexRequest struct {
	Name string `json:"name"`
}

type DemoIndexResponse struct {
	Message string `json:"message"`
}

func RegisterDemoService(g *gin.Engine, srv DemoHTTPServer) {
	group := g.Group("/api/demo")

	group.GET("/index/:name", _greeter_HTTP_Handler(srv))
}

func _greeter_HTTP_Handler(srv DemoHTTPServer) gin.HandlerFunc {
	return func(c *gin.Context) {
		var in DemoIndexRequest
		if err := c.BindQuery(&in); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"err": "bad params"})
		}
		out, err := srv.Greeter(in)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"err": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": out.Message})
	}
}
