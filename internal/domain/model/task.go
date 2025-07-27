package model

import (
	"github.com/MarrRoss/go-links-to-zip/internal/domain/exception"
	"time"
)

type Task struct {
	ID          ID
	Status      string
	Files       []*TaskFile
	ArchivePath *string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	EndedAt     *time.Time
}

// по status определяется, можно ли пользовтаелю скачать архив

func NewTask(files []*TaskFile) (*Task, error) {
	id := NewID()
	now := time.Now()
	newTask := Task{
		ID:          id,
		Status:      StatusCreated,
		Files:       files,
		ArchivePath: nil,
		CreatedAt:   now,
		UpdatedAt:   now,
		EndedAt:     nil,
	}
	return &newTask, nil
}

func (task *Task) SetArchivePath(path string) error {
	//if task.EndedAt != nil {
	//	return domain.ErrTaskIsDeleted
	//}
	if path == "" {
		return exception.ErrInvalidArchivePath
	}
	task.ArchivePath = &path
	task.Status = StatusFinished
	task.UpdatedAt = time.Now()
	timeNow := time.Now()
	task.EndedAt = &timeNow
	return nil
}
