package controller

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func (c *Controller) Health() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, HealthResponse{
			Message:     "ok",
			CurrentTime: time.Now().Format(time.RFC1123),
		})
	}
}
