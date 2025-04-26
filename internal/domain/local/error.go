package local

import (
	"errors"

	"github.com/HackathonUC2025-Hackfest/Hackfest-BE/pkg/cerr"
	"github.com/gofiber/fiber/v2"
)

var (
	ErrLBNotFound = cerr.New(fiber.ErrNotFound.Code, "local business not found", errors.New("account not found"))
)
