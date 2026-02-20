package router

import (
	"github.com/gin-gonic/gin"
	"gin-app/internal/handlers"
	// "gin-app/internal/middleware"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	// public routes
	// r.GET("/health", handlers.HealthCheck)
	r.GET("/login", handlers.Login)

	// // protected routes
	// protected := r.Group("/api")
	// protected.Use(middleware.AuthMiddleware())
	// protected.GET("/profile", handlers.Profile)

	return r
} 