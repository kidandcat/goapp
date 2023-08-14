package model

import (
	"time"

	"gorm.io/gorm"
)

type Reserva struct {
	gorm.Model
	Name         string
	Address      string
	Date         time.Time
	Duration     time.Duration
	Participants string
}
