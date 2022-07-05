package server

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func UseGinCORS() gin.HandlerFunc {
	return cors.Default()
}

func UseGinCORSAdvanced() gin.HandlerFunc {
	return cors.New(cors.Config{
		AllowOrigins:     []string{"https://develop-metrogalaxy.netlify.app"},
		AllowMethods:     []string{"PUT", "PATCH", "GET", "POST", "DELETE"},
		AllowHeaders:     []string{"Origin"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return true
		},
		MaxAge: 12 * time.Hour,
	})
}
