package bootstrap

import (
	"github.com/HackathonUC2025-Hackfest/Hackfest-BE/internal/domain/local/handler/rest"
	localRepository "github.com/HackathonUC2025-Hackfest/Hackfest-BE/internal/domain/local/repository"
	localService "github.com/HackathonUC2025-Hackfest/Hackfest-BE/internal/domain/local/service"
	sessionHandler "github.com/HackathonUC2025-Hackfest/Hackfest-BE/internal/domain/session/handler/rest"
	sessionRepository "github.com/HackathonUC2025-Hackfest/Hackfest-BE/internal/domain/session/repository"
	sessionService "github.com/HackathonUC2025-Hackfest/Hackfest-BE/internal/domain/session/service"
	userRepository "github.com/HackathonUC2025-Hackfest/Hackfest-BE/internal/domain/user/repository"
	"github.com/HackathonUC2025-Hackfest/Hackfest-BE/pkg/jwt"
	"github.com/gofiber/fiber/v2"
)

func (app *App) InitHandlers() {
	app.registerRoutes(app.jwt)
	app.MountRoutes()
	app.registerHealthCheck()
}

func (app *App) registerRoutes(jwt *jwt.JWTStruct) {
	userRepository := userRepository.New(app.postgres)
	sessionRepository := sessionRepository.New(app.postgres)
	localRepository := localRepository.New(app.postgres)

	authService := sessionService.New(userRepository, sessionRepository, jwt)
	localService := localService.New(localRepository, app.payment.snap, app.payment.coreapi)

	authHandler := sessionHandler.New(authService, app.validator)
	localHandler := rest.New(localService, app.validator, app.jwt)

	app.handlers = append(app.handlers, authHandler, localHandler)
}

func (app *App) MountRoutes() {
	routerGroup := app.http.Group("/api")
	for _, handler := range app.handlers {
		handler.Mount(routerGroup)
	}
}

func (app *App) registerHealthCheck() {
	app.http.Get("/health", func(ctx *fiber.Ctx) error {
		return ctx.SendString("Everything is good!")
	})
}
