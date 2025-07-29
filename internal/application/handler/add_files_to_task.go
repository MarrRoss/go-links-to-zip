package handler

import (
	"context"
	"github.com/MarrRoss/go-links-to-zip/internal/domain/exception"
	"github.com/MarrRoss/go-links-to-zip/internal/domain/model"
	"github.com/google/uuid"
)

type AddFilesToTaskCommand struct {
	TaskID    uuid.UUID
	FileLinks []string
}

func (h *AppHandler) AddFilesToTask(ctx context.Context, cmd AddFilesToTaskCommand) (model.ID, error) {
	if len(cmd.FileLinks) > 3 {
		h.observer.Logger.Error().Msg("invalid links count")
		return model.ID{}, exception.ErrInvalidLinksCount
	}
	task, err := h.taskStorage.GetTaskByID(ctx, model.UUIDtoID(cmd.TaskID)) // добавить ошибки репо
	if err != nil {
		h.observer.Logger.Error().Msg("failed to get task by id")
		return model.ID{}, err
	}
	taskFilesCount := len(task.Files)
	if taskFilesCount+len(cmd.FileLinks) > 3 {
		h.observer.Logger.Error().Msg("invalid links count")
		return model.ID{}, exception.ErrInvalidLinksCount
	}

	filesIDs := [3]model.ID{}
	for key, value := range cmd.FileLinks {
		id, err := h.CreateFile(ctx, value)
		if err != nil {
			h.observer.Logger.Error().Err(err).Msg("failed to create file")
			return model.ID{}, err
		}
		filesIDs[key] = id
	}

	// сборка архива если файлов 3
	// проверить доступность файлов по ссылкам
	return newTask.ID, nil
}

func (h *AppHandler) CreateFile(ctx context.Context, link string) (model.ID, error) {
	valid, parsedLink := model.IsValidURL(link)
	if !valid {
		h.observer.Logger.Trace().Msg("invalid url format")
		return model.ID{}, exception.ErrInvalidURL
	}
	valid = model.IsValidExtension(link)
	if !valid {
		h.observer.Logger.Trace().Msg("invalid file extension")
		return model.ID{}, exception.ErrInvalidFileExtension
	}
	newFile, err := model.NewFile(link, parsedLink)
	if err != nil {
		h.observer.Logger.Trace().Err(err).Msg("failed to create domain file")
		return model.ID{}, err
	}
	err = h.fileStorage.CreateFile(ctx, newFile)
	if err != nil {
		h.observer.Logger.Trace().Err(err).Msg("failed to add file to storage")
		return model.ID{}, err
	}
	h.fileStorage.PrintAllFiles()
	return newFile.ID, nil
}
