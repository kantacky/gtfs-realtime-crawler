package model

import (
	"time"
)

type VehiclePosition struct {
	TripID               *string     `gorm:"column:trip_id"`
	RouteID              *string     `gorm:"column:route_id"`
	DirectionID          *uint32     `gorm:"column:direction_id"`
	StartDatetime        *time.Time  `gorm:"column:start_datetime"`
	ScheduleRelationship *string     `gorm:"column:schedule_relationship"`
	VehicleID            *string     `gorm:"column:vehicle_id"`
	VehicleLabel         *string     `gorm:"column:vehicle_label"`
	VehiclePosition      *Coordinate `gorm:"column:vehicle_position"`
	CurrentStopSequence  *uint32     `gorm:"column:current_stop_sequence"`
	StopID               *string     `gorm:"column:stop_id"`
	Timestamp            *time.Time  `gorm:"column:timestamp"`
}
