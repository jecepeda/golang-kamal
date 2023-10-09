package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("welcome from golang-kamal!"))
	})
	r.Get("/up", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("up"))
	})
	port := ":3000"
	fmt.Println("Server running on port", port)
	http.ListenAndServe(port, r)
}
