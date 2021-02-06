package main

import "github.com/go-chi/chi"

func newRouter() *chi.Mux {
	return chi.NewRouter()
}

func (h *Handler) Register(r *chi.Mux) {
	r.Use(JWTParserMiddleware)
	r.Get("/health-check", h.HealthCheck)
}
