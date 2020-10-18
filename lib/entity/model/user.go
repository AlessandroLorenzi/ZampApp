package model

import (
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID             string `gorm:"type:varchar(36);primaryKey" json:"id"`
	Picture        string `gorm:"type:varchar(125) NOT NULL" json:"picture"`
	Email          string `gorm:"type:varchar(36) NOT NULL" json:"email"`
	NickName       string `gorm:"type:varchar(36) NOT NULL" json:"nick_name"`
	Description    string `gorm:"type:NOT NULL" json:"description"`
	HashedPassword string `gorm:"type:varchar(36) NOT NULL" json:"-"`
}

func NewUser(picture, email, nickname, description, plainPassword string) (User, error) {
	u := User{
		ID:          uuid.New().String(),
		Picture:     picture,
		Email:       email,
		NickName:    nickname,
		Description: description,
	}
	u.SetPassword(plainPassword)

	return u, nil
}

func (u *User) SetPassword(password string) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		logrus.Error("")
	}
	u.HashedPassword = string(hashedPassword)
}

func (u User) ValidatePassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.HashedPassword), []byte(password))

	return !(err != nil)
}
