package realtime

import (
	"fmt"
	"log"
	"time"

	"github.com/kantacky/gtfs-realtime-crawler/lib"
	"github.com/kantacky/gtfs-realtime-crawler/model"
)

func RecordVehiclePositions(message *FeedMessage, schemaName string) {
	vehiclePositions := []model.VehiclePosition{}

	for _, entity := range message.Entity {
		if entity.Vehicle != nil {
			startDatetime := lib.ParseISO8601(entity.Vehicle.Trip.GetStartDate() + "T" + entity.Vehicle.Trip.GetStartTime() + "+09:00")
			timestamp := unixTime(entity.Vehicle.Timestamp)
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

	writeToDB(vehiclePositions, schemaName)
}

func unixTime(unixTime *uint64) *time.Time {
	if unixTime == nil {
		return nil
	}
	timestamp := &time.Time{}
	*timestamp = time.Unix(int64(*unixTime), 0).Local()
	return timestamp
}

func writeToDB(vehiclePositions []model.VehiclePosition, schemaName string) {
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
