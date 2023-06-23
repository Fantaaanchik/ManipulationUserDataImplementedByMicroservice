package models

import "time"

type User struct {
	ID        int64     `json:"id" gorm:"column:id;primary_key;autoIncrement"`
	Fio       string    `json:"fio" gorm:"column:fio"`
	Number    string    `json:"number" gorm:"column:number"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoCreateTime"`
}
