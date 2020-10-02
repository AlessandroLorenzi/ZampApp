package model

import (
	"gorm.io/gorm"
)

type User struct {
	ID        int `gorm:"primary key"`
	CreatedAt []uint8
	UpdatedAt []uint8
	DeletedAt gorm.DeletedAt `gorm:"index"`

	Picture     string `gorm:"NOT NULL" json:"picture"`
	Email       string `gorm:"NOT NULL" json:"email"`
	NickName    string `gorm:"NOT NULL" json:"nick_name"`
	Description string `gorm:"NOT NULL" json:"description"`
}
