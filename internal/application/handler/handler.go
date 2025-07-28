package handler

import (
	"github.com/MarrRoss/go-links-to-zip/internal/observability"
	"github.com/MarrRoss/go-links-to-zip/internal/port"
)

type AppHandler struct {
	observer    *observability.Observability
	fileStorage port.FileRepository
	taskStorage port.TaskRepository
}

func NewAppHandler(
	observer *observability.Observability,
	fileStorage port.FileRepository,
	taskStorage port.TaskRepository,
) *AppHandler {
	return &AppHandler{
		observer:    observer,
		fileStorage: fileStorage,
		taskStorage: taskStorage,
	}
}
