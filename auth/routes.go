package auth

import (
	"antar-jemput/auth/routes"
	"github.com/labstack/echo/v4"
)

func RegisterAuthRoutes(e *echo.Echo) {
	authGroup := e.Group("/auth")
	authGroup.POST("/register", routes.RegisterUser)
}
