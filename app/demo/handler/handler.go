package handler

import (
	"fmt"

	"david.test.vercel_test/api_proto"
)

type demoHandler struct{}

func (d *demoHandler) Greeter(params *api_proto.DemoIndexRequest) (*api_proto.DemoIndexResponse, error) {
	return &api_proto.DemoIndexResponse{
		Message: fmt.Sprintf("Hello %s", params.Name),
	}, nil
}

func NewDemoHandler() api_proto.DemoHTTPServer {
	return new(demoHandler)
}
