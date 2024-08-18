package routes

import (
	"antar-jemput/models"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"net/http"
)

func DeleteBusRoute(c echo.Context) error {
	id := c.Param("id")
	db := c.Get("db").(*gorm.DB)

	var bus models.Bus
	if err := db.Where("id = ?", id).First(&bus).Error; err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"message": "Bus not found"})
	}

	if err := db.Delete(&bus).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Failed to delete bus"})
	}
	return c.JSON(http.StatusOK, map[string]string{"message": "Bus deleted successfully"})
}
