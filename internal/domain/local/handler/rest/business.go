package rest

import (
	"errors"

	"github.com/HackathonUC2025-Hackfest/Hackfest-BE/internal/domain/local"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func (h *LocalHandler) GetLocalBusiness(ctx *fiber.Ctx) error {
	request := local.QueryParamRequestGetLocals{
		City: ctx.Query("city", ""),
		Type: ctx.Query("type", "business"),
	}

	response, err := h.service.GetAllLocalsWithCity(ctx.Context(), request)
	if err != nil {
		return err
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "get all local businesses successful",
		"payload": response,
	})
}

func (h *LocalHandler) GetSpecificLocalBusiness(ctx *fiber.Ctx) error {
	localBusinessIDStr := ctx.Params("localBusinessID", "")

	localBusinessID, err := uuid.Parse(localBusinessIDStr)
	if err != nil {
		return errors.New("invalid uuid")
	}

	response, err := h.service.GetSpecificLocalBusiness(ctx.Context(), localBusinessID)
	if err != nil {
		return err
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "get specific local business successful",
		"payload": response,
	})
}
