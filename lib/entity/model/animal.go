package model

import (
	"github.com/google/uuid"
)

type Animal struct {
	ID            string   `gorm:"type:varchar(36);primaryKey" json:"id"`
	Name          string   `gorm:"type:varchar(15) NOT NULL" json:"name"`
	Breed         string   `gorm:"type:varchar(15) NOT NULL" json:"breed"`
	Size          int      `gorm:"type:int NOT NULL" json:"size"`
	Sex           bool     `gorm:"type:bool NOT NULL" json:"sex"`
	OwnerID       string   `gorm:"type:varchar(36) NOT NULL" json:"owner_id"`
	Owner         User     `gorm:"foreignKey:OwnerID" json:"owner" `
	Picture       string   `gorm:"type:varchar(125) NOT NULL" json:"picture"`
	Wormed        bool     `gorm:"type:bool NOT NULL" json:"wormed"`
	ChildFriendly bool     `gorm:"type:bool NOT NULL" json:"child_friendly"`
	Position      Location `gorm:"type:geometry NOT NULL" json:"position"`
	PositionDesc  string   `gorm:"type:text NOT NULL" json:"position_desc"`
	Description   string   `gorm:"type:text NOT NULL" json:"description"`
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
