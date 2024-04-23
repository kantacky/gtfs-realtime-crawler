package main

import (
	"flag"
	"log"
	"strings"
	"time"

	"github.com/joho/godotenv"
	"github.com/kantacky/gtfs-realtime-crawler/realtime"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("godotenv.Load error:", err)
	}

	var (
		url             = flag.String("u", "", "GTFS Realtime feed URL")
		agencyID        = flag.String("a", "", "Agency ID (UUID)")
		intervalSeconds = flag.Int("i", 15, "Interval seconds")
	)
	flag.Parse()

	if *url == "" {
		panic("url is required")
	}
	if *agencyID == "" {
		panic("agency is required")
	}

	schemaName := strings.ReplaceAll(*agencyID, "-", "")

	var lastChangedAt *time.Time

	for {
		message, err := realtime.GetMessage(*url)
		if err != nil {
			log.Println("realtime.GetMessage: ", err)
		}

		timestamp := time.Unix(int64(*message.Header.Timestamp), 0)
		if lastChangedAt == nil || *lastChangedAt != timestamp {
			lastChangedAt = &timestamp
			realtime.RecordVehiclePositions(message, "a"+schemaName)
			log.Println("Recorded")
		}

		time.Sleep(time.Duration(*intervalSeconds) * time.Second)
	}
}
