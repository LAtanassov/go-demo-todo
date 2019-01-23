package todo

import (
	"io"
	"net/http"
)

// http REST API / JSON
// gRPC, Graph API

func helloHandler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Hello, world!\n")
}
