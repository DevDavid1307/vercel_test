package handler

import (
	"fmt"

	"david.test.vercel_test/api_proto"
)

type demoHandler struct{}

func (d *demoHandler) Greeter(s string) (string, error) {
	return fmt.Sprintf("Hello %s", s), nil
}

func NewDemoHandler() api_proto.DemoHTTPServer {
	return new(demoHandler)
}
