package model

import (
	"context"
	"fmt"

	"github.com/google/uuid"

	"github.com/sirupsen/logrus"

	"github.com/paulmach/orb"

	"github.com/paulmach/orb/encoding/wkb"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Animal struct {
	ID            string   `gorm:"primary key" json:"id"`
	Name          string   `gorm:"NOT NULL" json:"name"`
	Breed         string   `gorm:"NOT NULL" json:"breed"`
	Size          int      `gorm:"NOT NULL" json:"size"`
	Sex           bool     `gorm:"NOT NULL" json:"sex"`
	OwnerID       int      `gorm:"NOT NULL" json:"owner_id"`
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
	OwnerID int,
	Owner User,
	Picture string,
	Wormed bool,
	ChildFriendly bool,
	Position Location,
	PositionDesc string,
	Description string,
) Animal {
	return Animal{
		ID:            uuid.New().String(),
		Name:          Name,
		Breed:         Breed,
		Size:          Size,
		Sex:           Sex,
		OwnerID:       OwnerID,
		Owner:         Owner,
		Picture:       Picture,
		Wormed:        Wormed,
		ChildFriendly: ChildFriendly,
		Position:      Position,
		PositionDesc:  PositionDesc,
		Description:   Description,
	}
}

type Location struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
}

// Scan implements the sql.Scanner interface
func (loc *Location) Scan(v interface{}) error {
	// Scan a value into struct from database driver
	//err := json.Unmarshal(v, loc)
	if v == nil {
		return nil
	}

	mysqlEncoding, ok := v.([]byte)
	if !ok {
		logrus.Errorf("did not scan: expected []byte but was %T", v)
		return fmt.Errorf("did not scan: expected []byte but was %T", v)
	}

	var g wkb.GeometryScanner
	err := g.Scan(mysqlEncoding[4:])
	if !ok {
		logrus.Errorf("did not scan: %v", err)
		return fmt.Errorf("did not scan: %w", err)
	}

	loc.X = g.Geometry.(orb.Point).X()
	loc.Y = g.Geometry.(orb.Point).Y()

	return err
}

func (loc Location) GormDataType() string {
	return "geometry"
}

func (loc Location) GormValue(ctx context.Context, db *gorm.DB) clause.Expr {
	return clause.Expr{
		SQL:  "ST_PointFromText(?, 4326)",
		Vars: []interface{}{fmt.Sprintf("POINT(%f %f)", loc.X, loc.Y)},
	}
}
