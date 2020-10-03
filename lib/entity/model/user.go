package model

type User struct {
	ID          int    `gorm:"primary key"`
	Picture     string `gorm:"NOT NULL" json:"picture"`
	Email       string `gorm:"NOT NULL" json:"email"`
	NickName    string `gorm:"NOT NULL" json:"nick_name"`
	Description string `gorm:"NOT NULL" json:"description"`
}
