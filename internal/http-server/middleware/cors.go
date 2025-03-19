package middleware

import (
	"github.com/ayayaakasvin/restapigolang/internal/config"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func CorsWithConfig(addresses config.ServiceAddresses) gin.HandlerFunc {
	var CorsDefaultConfig cors.Config = cors.Config{
		AllowOrigins:     addresses.Addresses,
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
	}

	return cors.New(CorsDefaultConfig)
}
