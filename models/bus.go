package models

import (
	"gorm.io/gorm"
)

type SeatPlan struct {
	Rows [][]string `json:"rows"`
}

type Bus struct {
	gorm.Model
	LicensePlate     string   `json:"license_plate" gorm:"unique;not null"`
	SeatPlan         SeatPlan `json:"seat_plan" gorm:"type:jsonb;not null"`
	OwnerID          uint     `json:"owner_id"`
	Owner            User     `json:"owner" gorm:"foreignKey:OwnerID"`
	BusModel         string   `json:"model" gorm:"not null"`
	Capacity         int      `json:"capacity" gorm:"not null"`
	Manufacturer     string   `json:"manufacturer"`
	YearOfProduction int      `json:"year_of_production"`
}
