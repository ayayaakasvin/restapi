package middleware

import (
	"net/http"
	"restapi/internal/errorset"
	helper "restapi/internal/lib/helperfunctions"
	"restapi/internal/lib/jwtutil"
	"restapi/internal/models/response"

	"github.com/gin-gonic/gin"
)

const (
	AuthorizationHeader = "Authorization"
)

func JWNAuthMiddleware () gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString, err := helper.FetchTokenFromContext(c)
		if err != nil {
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