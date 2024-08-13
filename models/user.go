package models

import "time"

type User struct {
	Id          int64  `gorm:"primaryKey" json:"id"`
	NamaLengkap string `gorm:"varchar(100)" json:"nama_lengkap"`
	Username    string `gorm:"varchar(100)" json:"username"`
	Password    string `gorm:"varchar(100)" json:"password"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
