package api

import (
	"net/http"

	"david.test.vercel_test/api_proto"
	"david.test.vercel_test/app/demo/service"

	"github.com/gin-gonic/gin"
)

func Demo(w http.ResponseWriter, r *http.Request) {
	serve := gin.New()

	api_proto.RegisterDemoService(serve, service.NewDemoHandler())

	serve.ServeHTTP(w, r)
}
