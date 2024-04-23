package model

import (
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Coordinate struct {
	Latitude  float32
	Longitude float32
}

func (coordinate Coordinate) GormDataType() string {
	return "POINT"
}

func (coordinate Coordinate) GormValue(ctx context.Context, db *gorm.DB) clause.Expr {
	return clause.Expr{
		SQL:  "POINT(?,?)",
		Vars: []interface{}{coordinate.Latitude, coordinate.Longitude},
	}
}

// Scan implements the sql.Scanner interface
func (coordinate *Coordinate) Scan(v interface{}) error {
	// Scan a value into struct from database driver
	return nil
}
