package main

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"
	"net/http"
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

	http.ListenAndServe(":8080", r)
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
