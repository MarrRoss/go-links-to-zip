package handler

import (
	"fmt"
	"github.com/MarrRoss/go-links-to-zip/internal/application/handler"
	"github.com/MarrRoss/go-links-to-zip/internal/presentation/request"
	"github.com/MarrRoss/go-links-to-zip/internal/presentation/response"
	"github.com/gofiber/fiber/v2"
)

func (h *PresentHandler) AddFilesToTask(ctx *fiber.Ctx) error {
	var pathReq request.GetTaskByIDRequest
	if err := ctx.ParamsParser(&pathReq); err != nil {
		h.observer.Logger.Trace().Err(err).Msg("failed to parse id request")
		return fmt.Errorf("failed to parse id request")
	}
	var bodyReq request.AddFilesToTask
	if err := ctx.BodyParser(&bodyReq); err != nil {
		h.observer.Logger.Trace().Err(err).Msg("failed to parse body request")
		return fmt.Errorf("failed to parse body request")
	}
	cmd := handler.AddFilesToTaskCommand{
		TaskID:    pathReq.TaskID,
		FileLinks: bodyReq.FileLinks,
	}
	id, err := h.appHandler.AddFilesToTask(ctx.UserContext(), cmd)
	if err != nil {
		h.observer.Logger.Error().Err(err).Msg("failed to add files to task")
		return fmt.Errorf("failed to add files to task")
	}
	return ctx.Status(fiber.StatusOK).JSON(response.NewAddTaskResponse(id))
}
