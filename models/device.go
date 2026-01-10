package models

import (
	"time"
)

type Device struct {
	ID       uint   `gorm:"primaryKey"`
	DeviceId string `gorm:"uniqueIndex"`
	LastSeen time.Time
}
