package crawler

import (
	"fmt"
	"log"
	"time"

	"github.com/kantacky/apis-go/transit_realtime"
	"github.com/kantacky/gtfs-realtime-crawler/lib"
	"github.com/kantacky/gtfs-realtime-crawler/model"
)

const rangeMinutes = 15

func RecordVehiclePositions(message *transit_realtime.FeedMessage, schemaName string) {
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
		log.Println("lib.GetSQLDB:", err)
	}
	defer sqldb.Close()

	gormdb, err := lib.GetGORMDB(sqldb)
	if err != nil {
		log.Println("lib.GetGORMDB:", err)
	}

	var records []model.VehiclePosition
	if err := gormdb.Table(fmt.Sprintf("%s.vehicle_positions", schemaName)).Where("timestamp > ?", time.Now().Add(-rangeMinutes*time.Minute)).Find(&records).Error; err != nil {
		log.Println("gormdb.Where:", err)
	}

	filteredVehiclePositions := []model.VehiclePosition{}
	flag := false
	for _, vehiclePosition := range vehiclePositions {
		if vehiclePosition.VehicleID == nil || vehiclePosition.Timestamp == nil {
			continue
		}
		for _, record := range records {
			if *vehiclePosition.VehicleID == *record.VehicleID && *vehiclePosition.Timestamp == *record.Timestamp {
				flag = true
				break
			}
		}
		if !flag {
			filteredVehiclePositions = append(filteredVehiclePositions, vehiclePosition)
		}
		flag = false
	}

	if err := gormdb.Table(fmt.Sprintf("%s.vehicle_positions", schemaName)).Create(filteredVehiclePositions).Error; err != nil {
		log.Println("db.Create:", err)
	}

	log.Printf(
		"Recorded: %d records added, %d duplicated records are in %d records recorded in %d minutes\n",
		len(filteredVehiclePositions),
		len(vehiclePositions)-len(filteredVehiclePositions),
		len(records),
		rangeMinutes,
	)
}
