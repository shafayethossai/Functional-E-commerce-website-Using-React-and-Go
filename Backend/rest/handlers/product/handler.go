package product

import (
	middleware "first-program/rest/middlewares"
)

type Handler struct {
	middlewares *middleware.Middlewares
	svc         Service
}

func NewHandler(middlewares *middleware.Middlewares, svc Service) *Handler { // Constructor function to create a new Handler instance
	return &Handler{
		middlewares: middlewares,
		svc:         svc,
	}
}
