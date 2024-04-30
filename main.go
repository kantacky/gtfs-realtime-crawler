package main

import (
	"flag"
	"log"
	"os"
	"strings"
	"time"

	"github.com/joho/godotenv"
	"github.com/kantacky/gtfs-realtime-crawler/crawler"
)

func main() {
	godotenv.Load()

	var (
		url             = flag.String("url", "", "GTFS Realtime feed URL")
		agencyID        = flag.String("agency", "", "Agency ID (UUID)")
		intervalSeconds = flag.Int("interval", 15, "Interval seconds")
	)
	flag.Parse()

	if *url == "" {
		if os.Getenv("FEED_URL") != "" {
			*url = os.Getenv("FEED_URL")
		} else {
			panic("url is required")
		}
	}
	if *agencyID == "" {
		if os.Getenv("AGENCY_ID") != "" {
			*agencyID = os.Getenv("AGENCY_ID")
		} else {
			panic("agency is required")
		}
	}

	schemaName := strings.ReplaceAll(*agencyID, "-", "")

	var lastChangedAt *time.Time

	for {
		go func() {
			message, err := crawler.GetMessage(*url)
			if err != nil {
				log.Println("crawler.GetMessage:", err)
			}

			timestamp := unixTime(message.Header.Timestamp)
			if lastChangedAt == nil || (timestamp != nil && *lastChangedAt != *timestamp) {
				lastChangedAt = timestamp
				crawler.RecordVehiclePositions(message, "a"+schemaName)
			}
		}()

		time.Sleep(time.Duration(*intervalSeconds) * time.Second)
	}
}

func unixTime(unixTime *uint64) *time.Time {
	if unixTime == nil {
		return nil
	}
	timestamp := &time.Time{}
	*timestamp = time.Unix(int64(*unixTime), 0).Local()
	return timestamp
}
