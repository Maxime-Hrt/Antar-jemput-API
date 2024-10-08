package auth

import (
	"antar-jemput/auth/routes"
	"github.com/labstack/echo/v4"
)

func RoutesAuth(e *echo.Echo) {
	authGroup := e.Group("/auth")
	authGroup.POST("/register", routes.RegisterUser)
	authGroup.POST("/login", routes.LoginUser)
	authGroup.DELETE("/delete/:id", routes.DeleteUser)
}
