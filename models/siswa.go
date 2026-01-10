package models

import "time"

type Siswa struct {
	ID           uint   `gorm:"primaryKey"`
	NIS          string `gorm:"unique;not null"`
	Nama         string `gorm:"size:100;not null"`
	JenisKelamin string `gorm:"size:1"` // L / P
	TempatLahir  string
	TanggalLahir time.Time
	Kelas        string
	Jurusan      string
	Alamat       string
	NamaWali     string
	NoHP         string
	Status       string `gorm:"default:aktif"`
	CreatedAt    time.Time
	UpdatedAt    time.Time

	// RELATION
	Kartu *Kartu `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
