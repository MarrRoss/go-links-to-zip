package adapter

import (
	"github.com/MarrRoss/go-links-to-zip/internal/observability"
	"sync"
)

type FileRepositoryImpl struct {
	observer *observability.Observability
	files    sync.Map
}

func NewFileRepositoryImpl(
	observer *observability.Observability,
) (*FileRepositoryImpl, error) {
	return &FileRepositoryImpl{
		observer: observer,
	}, nil
}
