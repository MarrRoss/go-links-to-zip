package port

import (
	"context"
	"github.com/MarrRoss/go-links-to-zip/internal/domain/model"
)

type TaskRepository interface {
	CreateTask(ctx context.Context, task *model.Task) error
	CheckActiveTasksLimit(ctx context.Context) error
	PrintAllTasks()
}
