package routes

import (
	"antar-jemput/models"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"net/http"
)

type CreateBusRouteRequest struct {
	LicensePlate     string          `json:"license_plate" validate:"required"`
	SeatPlan         models.SeatPlan `json:"seat_plan" validate:"required"`
	OwnerID          uint            `json:"owner_id" validate:"required"`
	BusModel         string          `json:"model" validate:"required"`
	Capacity         int             `json:"capacity" validate:"required"`
	Manufacturer     string          `json:"manufacturer"`
	YearOfProduction int             `json:"year_of_production"`
}

func CreateBusRoute(c echo.Context) error {
	req := new(CreateBusRouteRequest)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid request"})
	}

	if err := c.Validate(req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Validation failed", "error": err.Error()})
	}

	db := c.Get("db").(*gorm.DB)

	// Check if the owner exists
	var owner models.User
	if err := db.Where("id = ?", req.OwnerID).First(&owner).Error; err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Owner does not exist"})
	}

	bus := models.Bus{
		LicensePlate: req.LicensePlate,
		SeatPlan:     req.SeatPlan,
		OwnerID:      req.OwnerID,
		//Owner:            owner,
		BusModel:         req.BusModel,
		Capacity:         req.Capacity,
		Manufacturer:     req.Manufacturer,
		YearOfProduction: req.YearOfProduction,
	}

	if err := db.Create(&bus).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Failed to create bus"})
	}

	return c.JSON(http.StatusCreated, map[string]string{"message": "Bus created successfully"})
}
