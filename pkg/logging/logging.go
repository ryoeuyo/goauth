package logging

import (
	"log/slog"
	"os"

	"github.com/golang-cz/devslog"
	"github.com/ryoeuyo/slogdiscard"
)

const (
	envLocal = "local"
	envProd  = "prod"
	envTest  = "test"
)

func Setup(env string) *slog.Logger {
	switch env {
	case envLocal:
		return slog.New(devslog.NewHandler(
			os.Stdout,
			&devslog.Options{
				HandlerOptions:    &slog.HandlerOptions{Level: slog.LevelDebug},
				MaxSlicePrintSize: 10,
				SortKeys:          false,
				NewLineAfterLog:   true,
				StringerFormatter: true,
			},
		))
	case envProd:
		return slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
			Level: slog.LevelInfo,
		}))
	case envTest:
		return slog.New(slogdiscard.NewDiscardHandler())
	default:
		return nil
	}
}
