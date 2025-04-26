package session

import (
	"errors"

	"github.com/HackathonUC2025-Hackfest/Hackfest-BE/pkg/cerr"
	"github.com/gofiber/fiber/v2"
)

var (
	ErrSessionNotFound = cerr.New(fiber.ErrNotFound.Code, "session not found", errors.New("session not found"))
)
