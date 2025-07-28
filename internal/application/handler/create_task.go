package handler

import (
	"context"
	"github.com/MarrRoss/go-links-to-zip/internal/domain/model"
)

func (h *AppHandler) CreateTask(ctx context.Context) (model.ID, error) {
	err := h.taskStorage.CheckActiveTasksLimit(ctx)
	if err != nil {
		h.observer.Logger.Trace().Err(err).Msg("active task limit exceeded")
		return model.ID{}, err
	}
	newTask, err := model.NewTask()
	if err != nil {
		h.observer.Logger.Trace().Err(err).Msg("failed to create domain task")
		return model.ID{}, err
	}
	err = h.taskStorage.CreateTask(ctx, newTask)
	if err != nil {
		h.observer.Logger.Trace().Err(err).Msg("failed to add task to storage")
		return model.ID{}, err
	}
	h.taskStorage.PrintAllTasks()
	return newTask.ID, nil
}
