package rest

import (
	"errors"

	"github.com/HackathonUC2025-Hackfest/Hackfest-BE/internal/domain/local"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func (h *LocalHandler) GetTourGuides(ctx *fiber.Ctx) error {
	city := ctx.Query("city", "")

	response, err := h.service.GetTourGuides(ctx.Context(), city)
	if err != nil {
		return err
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "get all tour guides successful",
		"payload": response,
	})
}

func (h *LocalHandler) GetSpcificTA(ctx *fiber.Ctx) error {
	taIDStr := ctx.Params("taID", "")

	taID, err := uuid.Parse(taIDStr)
	if err != nil {
		return errors.New("invalid uuid")
	}

	response, err := h.service.GetSpecificTourGuide(ctx.Context(), taID)
	if err != nil {
		return err
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "get specific local business successful",
		"payload": response,
	})
}

func (h *LocalHandler) GetDateFullBooked(ctx *fiber.Ctx) error {
	taIDStr := ctx.Params("taID", "")

	_, err := uuid.Parse(taIDStr)
	if err != nil {
		return err
	}

	response, err := h.service.GetFullBook(ctx.Context(), taIDStr)
	if err != nil {
		return err
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "get booking details successful",
		"payload": response,
	})
}

func (h *LocalHandler) GenerateSnapPayment(ctx *fiber.Ctx) error {
	var request local.RequestGenerateSnapLink
	if err := ctx.BodyParser(&request); err != nil {
		return err
	}

	if err := h.validator.Struct(request); err != nil {
		return err
	}

	taIDStr := ctx.Params("taID", "")

	_, err := uuid.Parse(taIDStr)
	if err != nil {
		return errors.New("invalid uuid")
	}

	userIDRaw, ok := ctx.Locals("user_id").(string)
	if !ok {
		return errors.New("failed to get user_id")
	}

	_, err = uuid.Parse(userIDRaw)
	if err != nil {
		return err
	}

	request.TAID = taIDStr
	request.UserID = userIDRaw
	response, err := h.service.GenerateSnapPayment(ctx.Context(), request)
	if err != nil {
		return err
	}

	return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "generate snap url successful",
		"payload": response,
	})
}
