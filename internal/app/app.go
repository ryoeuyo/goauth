package app

import (
	"context"
	"log/slog"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/ryoeuyo/goauth/internal/config"
	"github.com/ryoeuyo/goauth/internal/controller"
	"github.com/ryoeuyo/goauth/internal/middleware"
	"github.com/ryoeuyo/goauth/internal/storage/inmemory"
	"github.com/ryoeuyo/goauth/internal/usecase/auth"
	"github.com/ryoeuyo/goauth/pkg/customvalidate"
)

type App struct {
	httpServer server
}

func New(cfg *config.Config, logger *slog.Logger) App {
	r := gin.New()
	r.Use(gin.Recovery(), middleware.Logger(logger))

	storage := inmemory.NewStorage()
	us := auth.New(storage, cfg.TokenTTL, cfg.SecretKey)

	validate := validator.New()
	validate.RegisterValidation("password", customvalidate.IsValidPassword)

	ctlr := controller.New(us, logger, validate)

	ctlr.InitRouters(r)

	return App{
		httpServer: newServer(cfg.Server, r),
	}
}

func (a *App) Start() {
	go a.httpServer.start()
}

func (a *App) Shutdown(ctx context.Context) error {
	return a.httpServer.shutdown(ctx)
}
