package models

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"gorm.io/gorm"
)

type SeatPlan [][]string

func (s *SeatPlan) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}
	return json.Unmarshal(bytes, &s)
}

func (s SeatPlan) Value() (driver.Value, error) {
	return json.Marshal(s)
}

type Bus struct {
	gorm.Model
	LicensePlate     string   `json:"license_plate" gorm:"unique;not null"`
	SeatPlan         SeatPlan `json:"seat_plan" gorm:"type:jsonb;not null"`
	OwnerID          uint     `json:"owner_id"`
	Owner            User     `json:"owner" gorm:"foreignKey:OwnerID"`
	BusModel         string   `json:"bus_model" gorm:"not null"`
	Capacity         int      `json:"capacity" gorm:"not null"`
	Manufacturer     string   `json:"manufacturer"`
	YearOfProduction int      `json:"year_of_production"`
}
