package main

import (
	"fmt"
	"net/http"
)

type requestHandler struct{}

func (rq *requestHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello World")
}

func main() {
	handler := &requestHandler{}
	server := &http.Server{
		Addr:    ":8080",
		Handler: handler,
	}

	fmt.Println("Server running on http://localhost:8080")
	err := server.ListenAndServe()
	if err != nil {
		fmt.Println("Server failed to start:", err)
	}
}
