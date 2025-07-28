package port

import (
	"context"
	"github.com/MarrRoss/go-links-to-zip/internal/domain/model"
)

type FileRepository interface {
	CreateFile(ctx context.Context, file *model.TaskFile) error
	PrintAllFiles()
}
