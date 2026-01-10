package models

import "time"

type Kartu struct {
	ID        uint   `gorm:"primaryKey"`
	UID       string `gorm:"unique;not null"`
	SiswaID   *uint  `gorm:"unique"` // nullable
	CreatedAt time.Time
	UpdatedAt time.Time

	// RELATION
	Siswa *Siswa
}
