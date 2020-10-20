package model

import (
	"context"
	"fmt"

	"github.com/paulmach/orb"
	"github.com/paulmach/orb/encoding/wkb"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Location struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
}

// Scan implements the sql.Scanner interface
func (loc *Location) Scan(v interface{}) error {
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

func (loc Location) GormValue(_ context.Context, _ *gorm.DB) clause.Expr {
	return clause.Expr{
		SQL:  "ST_PointFromText(?, 4326)",
		Vars: []interface{}{fmt.Sprintf("POINT(%f %f)", loc.X, loc.Y)},
	}
}
