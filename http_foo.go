package main // import "github.com/aphexddb/lightshow"

import (
	"io"
	"net/http"
)

// FooHandler is a placeholder api call
func FooHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "text/plain")
	io.WriteString(w, "ok")
}
