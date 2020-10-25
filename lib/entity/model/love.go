package model

type Love struct {
	PersonID string `gorm:"primaryKey;autoincrement:false;type:varchar(36) NOT NULL" json:"person_id"`
	Person   User   `gorm:"foreignKey:PersonID" json:"person" `
	AnimalID string `gorm:"primaryKey;autoincrement:false;type:varchar(36) NOT NULL" json:"animal_id"`
	Animal   Animal `gorm:"foreignKey:AnimalID" json:"animal" `
}
