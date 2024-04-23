package realtime

import (
	"fmt"
	"log"
	"time"

	"github.com/kantacky/gtfs-realtime-crawler/lib"
	"github.com/kantacky/gtfs-realtime-crawler/model"
)

func RecordVehiclePositions(url string, schemaName string) {
	message, err := GetMessage(url)
	if err != nil {
		log.Println("GetMessage: ", err)
	}

	vehiclePositions := []model.VehiclePosition{}

	for _, entity := range message.Entity {
		if entity.Vehicle != nil {
			startDatetime := &time.Time{}
			*startDatetime, err = time.Parse("2006-01-02T15:04:05-07:00", entity.Vehicle.Trip.GetStartDate()+"T"+entity.Vehicle.Trip.GetStartTime()+"+09:00")
			if err != nil {
				startDatetime = nil
			} else {
				*startDatetime = startDatetime.Local()
			}

			unixTime := entity.Vehicle.Timestamp
			timestamp := &time.Time{}
			if unixTime == nil {
				timestamp = nil
			} else {
				*timestamp = time.Unix(int64(*unixTime), 0).Local()
			}

			coordinate := &model.Coordinate{
				Latitude:  entity.Vehicle.Position.GetLatitude(),
				Longitude: entity.Vehicle.Position.GetLongitude(),
			}
			scheduleRelationship := entity.Vehicle.Trip.GetScheduleRelationship().String()
			vehiclePosition := model.VehiclePosition{
				TripID:               entity.Vehicle.Trip.TripId,
				RouteID:              entity.Vehicle.Trip.RouteId,
				DirectionID:          entity.Vehicle.Trip.DirectionId,
				StartDatetime:        startDatetime,
				ScheduleRelationship: &scheduleRelationship,
				VehicleID:            entity.Vehicle.Vehicle.Id,
				VehicleLabel:         entity.Vehicle.Vehicle.Label,
				VehiclePosition:      coordinate,
				CurrentStopSequence:  entity.Vehicle.CurrentStopSequence,
				StopID:               entity.Vehicle.StopId,
				Timestamp:            timestamp,
			}
			vehiclePositions = append(vehiclePositions, vehiclePosition)
		}
	}

	sqldb, err := lib.GetSQLDB()
	if err != nil {
		log.Println("lib.GetSQLDB: ", err)
	}
	defer sqldb.Close()

	gormdb, err := lib.GetGORMDB(sqldb)
	if err != nil {
		log.Println("lib.GetGORMDB: ", err)
	}

	if err := gormdb.Table(fmt.Sprintf("%s.vehicle_positions", schemaName)).Create(vehiclePositions).Error; err != nil {
		log.Println("db.Create: ", err)
	}
}
