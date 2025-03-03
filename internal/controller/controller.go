package controller

import (
	"log/slog"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/ryoeuyo/goauth/internal/usecase"
)

type Controller struct {
	us       usecase.AuthUseCase
	ll       *slog.Logger
	validate *validator.Validate
}

func New(us usecase.AuthUseCase, logger *slog.Logger, validate *validator.Validate) Controller {
	return Controller{
		us:       us,
		ll:       logger,
		validate: validate,
	}
}

func (c *Controller) InitRouters(r *gin.Engine) {
	apiR := r.Group("/api/v1")
	{
		apiR.GET("/health", c.Health())

		authR := apiR.Group("/auth")
		{
			authR.POST("/login", c.Login())
			authR.POST("register", c.Register())
		}
	}
}
