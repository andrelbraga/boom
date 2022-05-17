package boom

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// Healthz ...
func Healthz() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, time.Now())
	}
}
