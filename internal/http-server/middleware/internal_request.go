package middleware

import (
	"log/slog"
	"net/http"
	"strings"

	"restapi/internal/lib/sl"
	"restapi/internal/models/response"
	"github.com/gin-gonic/gin"
)

func AllowInternalRequests (log *slog.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		logger := log.With(slog.String("middleware", "AllowInternalRequests"))

		allowedOrigins := []string{"http://0.0.0.0:8088"}

		origin := c.GetHeader("Origin")
		if origin == "" {
			origin = c.Request.RemoteAddr // Fallback to remote address
		}

		logger.Info("Origin obtained", sl.Any("origin", origin))

		var allowed bool
		for _, allowedOrigin := range allowedOrigins {
			if strings.Contains(allowedOrigin, allowedOrigin) {
				allowed = true
				break
			}
		}

		if !allowed {
			response.Error(c, http.StatusForbidden, "Acces denied")
			c.Abort()
			return
		}

		c.Next()
	}
}
