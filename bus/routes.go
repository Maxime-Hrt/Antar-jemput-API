package bus

import (
	"antar-jemput/bus/routes"
	"github.com/labstack/echo/v4"
)

func RoutesBus(e *echo.Echo) {
	busGroup := e.Group("/bus")
	busGroup.POST("/create", routes.CreateBusRoute)
	//busGroup.GET("/get/:id", routes.GetBusRoute)
	//busGroup.GET("/get", routes.GetBusRoutes)
	//busGroup.PUT("/update/:id", routes.UpdateBusRoute)
	//busGroup.DELETE("/delete/:id", routes.DeleteBusRoute)
}
