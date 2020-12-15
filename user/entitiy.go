package user

import "time"

type User struct {
	ID           int `gorm:"primaryKey"`
	Username     string
	Password     string
	Nama_Lengkap string
	Foto         string
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
