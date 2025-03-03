package controller

import (
	"errors"
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/ryoeuyo/goauth/internal/usecase"
)

func (c *Controller) Register() gin.HandlerFunc {
	const fn = "controller.Register"
	ll := c.ll.With(slog.String("fn", fn))

	return func(ctx *gin.Context) {
		var request RegisterRequest
		if err := ctx.Bind(&request); err != nil {
			ll.Warn("failed to parse request", slog.String("error", err.Error()))

			ctx.JSON(http.StatusBadRequest, ErrorResponse{
				Message: ErrBadRequest.Error(),
			})

			return
		}

		if err := c.validate.Struct(request); err != nil {
			var error validator.ValidationErrors

			if errors.As(err, &error) {
				ll.Warn("validaton error", slog.String("error", err.Error()))

				ctx.JSON(http.StatusBadRequest, ErrorResponse{
					Message: ErrFailedValidate.Error(),
				})

				return
			}

			ll.Error("failed to validate", slog.String("error", err.Error()))

			ctx.JSON(http.StatusInternalServerError, ErrorResponse{
				Message: ErrInternalError.Error(),
			})

			return
		}

		id, err := c.us.Register(ctx.Request.Context(), request.Email, request.Password)
		if err != nil {
			if errors.Is(err, usecase.ErrUserIsExists) {
				ll.Warn("try register with exist email", slog.Any("email", request.Email))

				ctx.JSON(http.StatusConflict, ErrorResponse{
					Message: ErrInvalidCredentials.Error(),
				})

				return
			}

			ll.Error("failed to regitser user", slog.Any("error", err))

			ctx.JSON(http.StatusInternalServerError, ErrorResponse{
				Message: ErrInternalError.Error(),
			})

			return
		}

		ll.Info("registered user", slog.String("id", id))
	}
}
