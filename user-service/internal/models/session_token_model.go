package models

import "time"

type SessionToken struct {
	ID         string    `json:"id" gorm:"primaryKey"`
	UserID     string    `json:"user_id"`
	Token      string    `json:"token"`
	CreatedAt  time.Time `json:"created_at"`
	ExpiredAt  time.Time `json:"expire_at"`
	DeviceInfo string    `json:"device_info"`
	User       User      `json:"user" gorm:"foreignKey:UserID;references:ID"`
}
