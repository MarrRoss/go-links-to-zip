package port

import (
	"context"
	"github.com/MarrRoss/go-links-to-zip/internal/domain/model"
)

type TaskRepository interface {
	AddTask(ctx context.Context, task *model.Task) error
}
