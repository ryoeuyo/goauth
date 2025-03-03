package main

import (
	"context"
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	"github.com/ryoeuyo/goauth/internal/app"
	"github.com/ryoeuyo/goauth/internal/config"
	"github.com/ryoeuyo/goauth/pkg/logging"
)

func main() {
	cfg := config.MustLoad()
	ll := logging.Setup(cfg.Env)

	ll.Debug("config initialized", slog.Any("config", cfg))

	application := app.New(cfg, ll)

	go application.Start()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	<-quit

	if err := application.Shutdown(context.Background()); err != nil {
		ll.Error("failed to graceful shutdown", slog.Any("error", err))
		os.Exit(1)
	}

	ll.Info("service shutdowned")
}
