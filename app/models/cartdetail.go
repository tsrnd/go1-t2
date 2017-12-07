package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

type CartDetail struct {
	gorm.Model
	Price     float64
	Quantity  int64
	CreatedAt time.Time
	UpdatedAt time.Time
}
