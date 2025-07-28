package request

import "github.com/google/uuid"

type GetTaskByIDRequest struct {
	TaskID uuid.UUID `params:"task_id" format:"uuid"`
}

type AddFilesToTask struct {
	FileLinks []string `json:"file_links"`
}
