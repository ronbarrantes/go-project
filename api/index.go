package handler

import (
	"fmt"
	"net/http"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	count := 0

	if r.URL.Path == "/hello" {

		fmt.Fprintf(w, "<p>hello from hello route</p>")
		return
	}

	count += 1

	fmt.Fprintf(w, "<h1>Hello from vercel, clicked %d</h1>", count)
}
