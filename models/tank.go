package models

import (
	"time"
)

type Tank struct {
	ID           uint      `json:"id"`
	Name         string    `json:"name"`
	Volume       uint      `json:"volume"`
	Location     string    `json:"location"`
	Description  string    `json:"description"`
	Manufacturer string    `json:"manufacturer"`
	Model        string    `json:"model"`
	CreatedAt    time.Time `json:"_"`
	UpdatedAt    time.Time `json:"_"`
}
