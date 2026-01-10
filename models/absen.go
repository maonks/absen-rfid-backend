package models

import "time"

type Absen struct {
	ID       uint `gorm:"primaryKey"`
	UID      string
	DeviceId string
	Waktu    time.Time `gorm:"type:timestamp"`
}
