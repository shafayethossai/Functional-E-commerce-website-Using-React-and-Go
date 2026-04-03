package product

import (
	middleware "first-program/rest/middlewares"
	"net/http"
)

func (h *Handler) RegisterRoutes(mux *http.ServeMux, manager *middleware.Manager) {
	mux.Handle("GET /products", manager.With( // End points to get all products
		http.HandlerFunc(h.GetProducts),
	))

	mux.Handle("POST /products", manager.With( // End points to create a new product
		http.HandlerFunc(h.CreateProduct),
		h.middlewares.AuthenticateJWT,
	))

	mux.Handle("GET /products/{id}", manager.With( // End points to get a product by its ID
		http.HandlerFunc(h.GetProduct),
	))

	mux.Handle("PUT /products/{id}", manager.With( // End points to update a product by its ID
		http.HandlerFunc(h.UpdateProducts),
		h.middlewares.AuthenticateJWT,
	))

	mux.Handle("DELETE /products/{id}", manager.With( // End points to delete a product by its ID
		http.HandlerFunc(h.DeleteProducts),
		h.middlewares.AuthenticateJWT,
	))
}
