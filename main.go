package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/RichDom2185/glam-mailer-api/handlers"
	"github.com/RichDom2185/glam-mailer-api/models"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"https://*", "http://*"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
	}))
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello from glam-mailer backend!")
	})

	r.Post("/convert", func(w http.ResponseWriter, r *http.Request) {
		var params models.InlineCSSPostParams
		err := json.NewDecoder(r.Body).Decode(&params)
		if err != nil {
			log.Fatalln(err)
		}

		// TODO: Handle data using models
		inlined := handlers.InlineCSS(params.HTML)
		substituted := handlers.SubstituteCSSVariables(inlined)
		cleaned := handlers.RemoveStyleTags(substituted)
		fmt.Fprintf(w, cleaned)
	})

	fmt.Println("Server started...")
	err := http.ListenAndServe("127.0.0.1:8080", r)
	if err != nil {
		log.Fatalln(err)
	}
}
