package main

import (
	"fmt"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /", func(w http.ResponseWriter, _ *http.Request) {
		w.Write([]byte("hellow server"))
	})

	fmt.Println("server listening on server 8080")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		panic(err)
	}
}
