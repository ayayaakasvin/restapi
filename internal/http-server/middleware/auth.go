package middleware

import (
	"net/http"
	"restapi/internal/errorset"
	"restapi/internal/lib/jwtutil"
	"restapi/internal/models/response"
	"strings"

	"github.com/gin-gonic/gin"
)

const (
	AuthorizationHeader = "Authorization"
)

func JWNAuthMiddleware () gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader(AuthorizationHeader)
		if authHeader == "" {
			response.Error(c, http.StatusUnauthorized, errorset.ErrAuthorizationMissing)
			c.Abort()
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		if tokenString == authHeader {
			response.Error(c, http.StatusUnauthorized, errorset.ErrAuthorizationMissing)
			c.Abort()
			return
		}

		claims, err := jwtutil.ValidateJWT(tokenString)
		if err != nil {
			response.Error(c, http.StatusUnauthorized, err.Error())
			c.Abort()
			return
		}

		c.Set("userId", claims["userId"])
	}
}