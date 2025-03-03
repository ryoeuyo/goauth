package middleware

import (
	"log/slog"

	"github.com/gin-gonic/gin"
)

func Logger(logger *slog.Logger) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		logger.Info("received request", slog.Any("information", map[string]interface{}{
			"method": ctx.Request.Method,
			"uri":    ctx.Request.RequestURI,
		}))
	}
}
