package handler

import (
	"fmt"
	"github.com/MarrRoss/go-links-to-zip/internal/presentation/response"
	"github.com/gofiber/fiber/v2"
)

func (h *PresentHandler) CreateTask(ctx *fiber.Ctx) error {
	id, err := h.appHandler.CreateTask(ctx.UserContext())
	if err != nil {
		h.observer.Logger.Error().Err(err).Msg("failed to create task")
		return fmt.Errorf("failed to create task")
	}
	return ctx.Status(fiber.StatusOK).JSON(response.NewAddTaskResponse(id))
}
