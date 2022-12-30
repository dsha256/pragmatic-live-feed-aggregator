package server

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"time"
)

// CORS adds cors policy. From https://github.com/gin-contrib/cors/issues/29.
func CORS() gin.HandlerFunc {
	const maxAge = 12 * time.Hour
	return cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET"},
		AllowHeaders:     []string{"Origin", "Authorization", "Content-Type", "Content-Length"},
		ExposeHeaders:    nil,
		MaxAge:           maxAge,
		AllowCredentials: false,
		AllowWebSockets:  false,
		AllowFiles:       false,
	})
}
