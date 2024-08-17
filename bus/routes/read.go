package routes

import (
	"antar-jemput/models"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"net/http"
)

func GetBusRoute(c echo.Context) error {
	id := c.Param("id")
	db := c.Get("db").(*gorm.DB)

	var bus models.Bus
	if err := db.Preload("Owner").Where("id = ?", id).First(&bus).Error; err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"message": "Bus not found"})
	}
	return c.JSON(http.StatusOK, bus)
}

func GetBusRoutes(c echo.Context) error {
	db := c.Get("db").(*gorm.DB)

	var buses []models.Bus
	if err := db.Preload("Owner").Find(&buses).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Failed to get buses"})
	}
	return c.JSON(http.StatusOK, buses)
}

func GetBusRoutesByOwner(c echo.Context) error {
	ownerID := c.Param("owner_id")

	if ownerID == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Owner ID is required"})
	}

	db := c.Get("db").(*gorm.DB)

	var buses []models.Bus
	if err := db.Preload("Owner").Where("owner_id = ?", ownerID).Find(&buses).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Failed to get buses"})
	}

	if len(buses) == 0 {
		return c.JSON(http.StatusNotFound, map[string]string{"message": "No bus found"})
	}

	return c.JSON(http.StatusOK, buses)
}
