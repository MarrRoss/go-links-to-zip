package model

import (
	"fmt"
	"github.com/MarrRoss/go-links-to-zip/internal/domain/exception"
	"net/url"
	"path"
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

// TODO: добавить поле с ошибкой?

func NewFile(link string, parsedLink url.URL) (*TaskFile, error) {
	if link == "" {
		return nil, exception.ErrInvalidFileLink
	}
	id := NewID()
	name := CreateFileName(link, parsedLink, id)
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

func CreateFileName(link string, parsedLink url.URL, id ID) string {
	defaultName := fmt.Sprintf("file-%s%s", id, path.Ext(link))
	base := path.Base(parsedLink.Path)
	if base == "/" || base == "." || base == "" {
		return defaultName
	}
	return base
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
