package model

import (
	"time"
)

type VehiclePosition struct {
	TripID               *string     `gorm:"type:varchar(255)"`
	RouteID              *string     `gorm:"type:varchar(255)"`
	DirectionID          *uint32     `gorm:"type:integer"`
	StartDatetime        *time.Time  `gorm:"type:timestamp with time zone"`
	ScheduleRelationship *string     `gorm:"type:varchar(255)"`
	VehicleID            *string     `gorm:"type:varchar(255)"`
	VehicleLabel         *string     `gorm:"type:varchar(255)"`
	VehiclePosition      *Coordinate `gorm:"type:point"`
	CurrentStopSequence  *uint32     `gorm:"type:integer"`
	StopID               *string     `gorm:"type:varchar(255)"`
	Timestamp            *time.Time  `gorm:"type:timestamp with time zone"`
}
