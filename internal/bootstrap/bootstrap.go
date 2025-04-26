package bootstrap

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/HackathonUC2025-Hackfest/Hackfest-BE/internal/infra/config"
	"github.com/HackathonUC2025-Hackfest/Hackfest-BE/internal/infra/db"
	"github.com/HackathonUC2025-Hackfest/Hackfest-BE/internal/infra/http"
	"github.com/HackathonUC2025-Hackfest/Hackfest-BE/internal/infra/logger"
	"github.com/HackathonUC2025-Hackfest/Hackfest-BE/internal/infra/storage"
	"github.com/HackathonUC2025-Hackfest/Hackfest-BE/pkg/jwt"
	_validator "github.com/HackathonUC2025-Hackfest/Hackfest-BE/pkg/validator"
	supabasestorageuploader "github.com/adityarizkyramadhan/supabase-storage-uploader"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog/log"
)

var (
	app *App
)

type Handler interface {
	Mount(router fiber.Router)
}

type App struct {
	http      *fiber.App
	config    *config.Env
	postgres  *sqlx.DB
	validator *validator.Validate
	handlers  []Handler
	jwt       *jwt.JWTStruct
	storage   *supabasestorageuploader.Client
}

func Initialize() error {
	env, err := config.LoadEnv()
	if err != nil {
		return err
	}

	postgres, err := db.NewPostgres(env)
	if err != nil {
		return err
	}

	jwt := jwt.New(env.JWTSecret)
	val := _validator.New()
	http := http.NewFiber()
	storage := storage.New(env.StorageURL, env.StorageToken, env.StorageBucket)

	app = &App{
		http:      http,
		config:    env,
		postgres:  postgres,
		validator: val,
		jwt:       jwt,
		storage:   storage,
	}

	logger.New()
	app.InitHandlers()

	go shutdown()

	return app.http.Listen(fmt.Sprintf(":%d", env.AppPort))
}

func shutdown() {
	signalCh := make(chan os.Signal, 1)
	signal.Notify(signalCh, syscall.SIGTERM, syscall.SIGINT, syscall.SIGHUP)
	<-signalCh

	log.Info().Msg("Received shutdown signal")
	log.Info().Msg("Shutting down...")

	_ = app.postgres.Close()
	_ = app.http.Shutdown()
}
