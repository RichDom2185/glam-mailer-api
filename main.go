package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello from glam-mailer backend!")
	})

	fmt.Println("Server started...")
	err := http.ListenAndServe("127.0.0.1:8080", r)
	if err != nil {
		log.Fatalln(err)
	}
}
