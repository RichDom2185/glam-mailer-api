package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

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

		syntaxHighlight, err := os.ReadFile("assets/css/github.css")
		if err != nil {
			log.Fatalln(err)
		}

		// TODO: Handle data using models
		tailwind := handlers.GenerateRequiredTailwindStyles(params.HTML)
		inlined := handlers.InlineCSS("<style>\n" + string(syntaxHighlight) + tailwind + "\n</style>\n" + params.HTML)
		substitutedCSS := handlers.SubstituteCSSVariables(inlined)
		substitutedSVG := handlers.ConvertSVGToImg(substitutedCSS)
		cleaned := handlers.RemoveStyleTags(substitutedSVG)
		fmt.Fprintf(w, cleaned)
	})

	appMode := os.Getenv("GO_ENV")
	if appMode == "" {
		appMode = "production"
	}
	fmt.Printf("Server started in %s mode!", appMode)

	host := ""
	if appMode == "development" {
		host = "127.0.0.1"
	}

	err := http.ListenAndServe(host+":8080", r)
	if err != nil {
		log.Fatalln(err)
	}
}
