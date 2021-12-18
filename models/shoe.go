package models

import (
	"time"

	"gorm.io/gorm"
)

type Shoe struct {
	ID        uint           `json:"id,string"`
	Name      string         `json:"name"`
	Style     string         `json:"style"`
	Colour    string         `json:"colour"`
	Material  string         `json:"material"`
	Price     float64        `json:"price,string"`
	CreatedAt time.Time      `json:"created"`
	UpdatedAt time.Time      `json:"updated"`
	DeletedAt gorm.DeletedAt `json:"deleted" gorm:"index"`
}
