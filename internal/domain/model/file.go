package model

import (
	"github.com/MarrRoss/go-links-to-zip/internal/domain/exception"
	"time"
)

type TaskFile struct {
	ID        ID
	Name      string
	Link      string
	Status    string
	Error     string
	CreatedAt time.Time
	UpdatedAt time.Time
	EndedAt   *time.Time
}

// TODO: добавить имя файла?
// TODO: добавить поле с ошибкой?

func NewFiles(name, link string) (*TaskFile, error) {
	if link == "" {
		return nil, exception.ErrInvalidFileLink
	}
	id := NewID()
	now := time.Now()
	newFileLink := TaskFile{
		ID:        id,
		Name:      name,
		Link:      link,
		Status:    StatusCreated,
		Error:     "",
		CreatedAt: now,
		UpdatedAt: now,
		EndedAt:   nil,
	}
	return &newFileLink, nil
}

func (file *TaskFile) SetError(err string) error {
	//if file.EndedAt != nil {
	//	return domain.ErrFileIsDeleted
	//}
	if err == "" {
		return exception.ErrInvalidFileError
	}
	file.Error = err
	file.Status = StatusError
	file.UpdatedAt = time.Now()
	return nil
}
