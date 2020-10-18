package model

import (
	"github.com/google/uuid"
)

type Animal struct {
	ID            string   `gorm:"primary key" json:"id"`
	Name          string   `gorm:"NOT NULL" json:"name"`
	Breed         string   `gorm:"NOT NULL" json:"breed"`
	Size          int      `gorm:"NOT NULL" json:"size"`
	Sex           bool     `gorm:"NOT NULL" json:"sex"`
	OwnerID       string   `gorm:"NOT NULL" json:"owner_id"`
	Owner         User     `gorm:"foreignKey:OwnerID" json:"owner" `
	Picture       string   `gorm:"NOT NULL" json:"picture"`
	Wormed        bool     `gorm:"NOT NULL" json:"wormed"`
	ChildFriendly bool     `gorm:"NOT NULL" json:"child_friendly"`
	Position      Location `gorm:"NOT NULL" json:"position"`
	PositionDesc  string   `gorm:"NOT NULL" json:"position_desc"`
	Description   string   `gorm:"NOT NULL" json:"description"`
}

func NewAnimal(
	Name string,
	Breed string,
	Size int,
	Sex bool,
	OwnerID string,
	Picture string,
	Wormed bool,
	ChildFriendly bool,
	Position Location,
	PositionDesc string,
	Description string,
) (Animal, error) {
	return Animal{
		ID:            uuid.New().String(),
		Name:          Name,
		Breed:         Breed,
		Size:          Size,
		Sex:           Sex,
		OwnerID:       OwnerID,
		Picture:       Picture,
		Wormed:        Wormed,
		ChildFriendly: ChildFriendly,
		Position:      Position,
		PositionDesc:  PositionDesc,
		Description:   Description,
	}, nil
}
