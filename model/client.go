package model

import (
	"time"
)

type Client struct {
	Id       uint   `gorm:"primary_key"`
	Email    string `gorm:"not null"`
	Password string `gorm:"not null"`
	NickName string
	CreateAt time.Time `gorm:"default: NOW()"`
	UpdateAt time.Time
}

// Set Client's table name to be `client`
func (Client) TableName() string {
	return "client"
}
