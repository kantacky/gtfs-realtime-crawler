package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/kantacky/gtfs-realtime-crawler/realtime"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("URL is not found. Input URL as an argument.")
	}
	url := os.Args[1]
	checkFeed(url)
}

func checkRenewalInterval(url string) {
	log.Println("Start fetching data from", url)

	message, err := realtime.GetMessage(url)
	if err != nil {
		log.Fatal(err)
	}

	lastChangedAt := time.Unix(int64(*message.Header.Timestamp), 0)

	for {
		message, err := realtime.GetMessage(url)
		if err != nil {
			log.Fatal(err)
		}

		timestamp := time.Unix(int64(*message.Header.Timestamp), 0)
		if lastChangedAt != timestamp {
			log.Printf(
				"%s; Last change was %d seconds ago.\n",
				timestamp.Format("2006-01-02T15:04:05+09:00"),
				int(timestamp.Sub(lastChangedAt).Abs().Seconds()),
			)
			lastChangedAt = timestamp
		}

		time.Sleep(1 * time.Second)
	}
}

func checkFeed(url string) {
	message, err := realtime.GetMessage(url)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("GTFS Realtime Version: %s  \n", message.Header.GetGtfsRealtimeVersion())
	fmt.Printf("Incrementality: %s  \n", message.Header.GetIncrementality())
	timestamp := time.Unix(int64(message.Header.GetTimestamp()), 0)
	timestampString := timestamp.Format("2006-01-02T15:04:05+09:00")
	fmt.Printf("Timestamp: %s  \n", timestampString)
	fmt.Printf("Number of entities: %d  \n", len(message.Entity))

	fmt.Println("| ID | Trip ID | Route ID | Direction ID | Start Time | Start Date | Schedule Relationship | Vehicle ID | Vehicle Label | Latitude | Longitude | Current Stop Sequence | Stop ID | Timestamp |")
	fmt.Println("| :-- | :-- | :-- | --: | :-- | :-- | :-- | :-- | :-- | :-- | :-- | --: | :-- | :-- |")

	for _, entity := range message.Entity {
		if entity.Vehicle != nil {
			fmt.Printf(
				"| %s | %s | %s | %d | %s | %s | %s | %s | %s | %f | %f | %d | %s | %s |\n",
				entity.GetId(),
				entity.Vehicle.Trip.GetTripId(),
				entity.Vehicle.Trip.GetRouteId(),
				entity.Vehicle.Trip.GetDirectionId(),
				entity.Vehicle.Trip.GetStartTime(),
				entity.Vehicle.Trip.GetStartDate(),
				entity.Vehicle.Trip.GetScheduleRelationship(),
				entity.Vehicle.Vehicle.GetId(),
				entity.Vehicle.Vehicle.GetLabel(),
				entity.Vehicle.Position.GetLatitude(),
				entity.Vehicle.Position.GetLongitude(),
				entity.Vehicle.GetCurrentStopSequence(),
				entity.Vehicle.GetStopId(),
				time.Unix(int64(entity.Vehicle.GetTimestamp()), 0).Format("2006-01-02T15:04:05+09:00"),
			)
		}
	}
}
