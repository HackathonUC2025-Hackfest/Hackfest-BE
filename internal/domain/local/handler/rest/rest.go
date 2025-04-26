package rest

import (
	"github.com/HackathonUC2025-Hackfest/Hackfest-BE/internal/domain/local/service"
	"github.com/HackathonUC2025-Hackfest/Hackfest-BE/internal/middleware"
	"github.com/HackathonUC2025-Hackfest/Hackfest-BE/pkg/jwt"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type LocalHandler struct {
	service   service.LocalServiceItf
	validator *validator.Validate
	jwt       *jwt.JWTStruct
}

func New(service service.LocalServiceItf, validator *validator.Validate, jwt *jwt.JWTStruct) *LocalHandler {
	return &LocalHandler{service: service, validator: validator, jwt: jwt}
}

func (h *LocalHandler) Mount(router fiber.Router) {
	local := router.Group("/locals")
	local.Use(middleware.Authentication(h.jwt))

	local.Get("/", h.GetLocalBusiness)
	local.Get("/:localBusinessID", h.GetSpecificLocalBusiness)

	ta := router.Group("/tourist-attractions")
	ta.Use(middleware.Authentication(h.jwt))

	ta.Get("/", h.GetTourGuides)
	ta.Get("/:taID", h.GetSpcificTA)
	ta.Get("/:taID/book", h.GetDateFullBooked)
	ta.Post("/:taID/book", h.GenerateSnapPayment)

}
