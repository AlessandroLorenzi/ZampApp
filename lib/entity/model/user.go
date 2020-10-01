package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Picture  string `json:"picture"`
	Email    string `json:"email"`
	NickName string `json:"nick_name"`
}
