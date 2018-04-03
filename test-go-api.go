package main

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"
	"net/http"
	"os"
	"fmt"
	"log"
)

func main() {
	
	r := chi.NewRouter()
	
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.DefaultCompress)
	r.Use(render.SetContentType(render.ContentTypeJSON))

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		render.Render(w, r, NewWelcomeResponse("hello"))
	})

	port := getEnv("PORT", "8080")
	
	log.Printf("Starting on port %v...", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", port), r))
}

func getEnv(key, fallback string) string {
	value, exists := os.LookupEnv(key)
	if (!exists) {
		return fallback
	}

	return value
}

type WelcomeResponse struct {
	Greeting string `json:"greeting"`
}

func (wr *WelcomeResponse) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func NewWelcomeResponse(greeting string) *WelcomeResponse {
	return &WelcomeResponse{Greeting: greeting}
}
