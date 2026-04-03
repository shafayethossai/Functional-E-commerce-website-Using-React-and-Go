package user

import (
	middleware "first-program/rest/middlewares"
	"net/http"
)

func (h *Handler) RegisterRoutes(mux *http.ServeMux, manager *middleware.Manager) {
	mux.Handle("POST /users", manager.With( // End points to create a new user
		http.HandlerFunc(h.CreateUser),
	))

	mux.Handle("POST /users/login", manager.With( // End points to login a user
		http.HandlerFunc(h.Login),
	))
}
