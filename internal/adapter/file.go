package adapter

import (
	"github.com/MarrRoss/go-links-to-zip/internal/domain/model"
	"github.com/MarrRoss/go-links-to-zip/internal/observability"
	"sync"
)

type FileRepositoryImpl struct {
	observer *observability.Observability
	files    map[model.ID]*model.TaskFile
	mu       sync.RWMutex
}

func NewFileRepositoryImpl(
	observer *observability.Observability,
) (*FileRepositoryImpl, error) {
	return &FileRepositoryImpl{
		observer: observer,
		files:    make(map[model.ID]*model.TaskFile),
	}, nil
}
