package api

import (
	"fmt"
	"net/http"
)

func Demo(w http.ResponseWriter, _ *http.Request) {
	_, _ = fmt.Fprintf(w, "Hello World!")
}
