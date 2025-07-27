package handler

import "github.com/MarrRoss/go-links-to-zip/internal/observability"

type Handler struct {
	observer *observability.Observability
}

func New(
	observer *observability.Observability,
) *Handler {
	return &Handler{observer: observer}
}
