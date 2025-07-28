package handler

import (
	"github.com/MarrRoss/go-links-to-zip/internal/application/handler"
	"github.com/MarrRoss/go-links-to-zip/internal/observability"
)

type PresentHandler struct {
	observer   *observability.Observability
	appHandler *handler.AppHandler
}

func NewPresentHandler(
	observer *observability.Observability,
	appHandler *handler.AppHandler,
) *PresentHandler {
	return &PresentHandler{
		observer:   observer,
		appHandler: appHandler,
	}
}
