package routes

import (
	"antar-jemput/auth/models"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"net/http"
)

func DeleteUser(c echo.Context) error {
	id := c.Param("id")
	db := c.Get("db").(*gorm.DB)

	var user models.User
	if err := db.Where("id = ?", id).First(&user).Error; err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"message": "User not found"})
	}

	if err := db.Delete(&user).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Failed to delete user"})
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "User deleted successfully"})
}
