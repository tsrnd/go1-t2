package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Category struct {
	gorm.Model
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}
