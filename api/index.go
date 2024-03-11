package handler

import (
	"fmt"
	"net/http"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/hello" {

		fmt.Fprintf(w, "<p>hello from hello route</p>")
		return
	}
	fmt.Fprintf(w, "<h1>Hello from vercel</h1>")
}
