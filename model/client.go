package model

import (
	"time"
)

type Client struct {
	Id       uint      `gorm:"primary_key"`
	Email    string    `gorm:"not null"`
	Password string    `gorm:"not null"`
	NickName string    `gorm:"default: null"`
	CreateAt time.Time `gorm:"default: NOW()"`
	UpdateAt time.Time `gorm:"default: null"`
}

// Set Client's table name to be `client`
func (Client) TableName() string {
	return "client"
}
