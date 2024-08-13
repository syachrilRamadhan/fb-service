package models

import "time"

type User struct {
	Id          int64     `gorm:"primaryKey" json:"id"`
	NamaLengkap string    `gorm:"varchar(100)" json:"nama_lengkap"`
	Username    string    `gorm:"varchar(100),unique" json:"username"`
	Password    string    `gorm:"varchar(100)" json:"password"`
	CreatedAt   time.Time `gorm:"column:created_at"`
	UpdatedAt   time.Time `gorm:"column:updated_at"`
}
