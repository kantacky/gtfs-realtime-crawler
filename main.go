package main

import (
	"flag"
	"log"
	"strings"
	"time"

	"github.com/kantacky/gtfs-realtime-crawler/lib"
	"github.com/kantacky/gtfs-realtime-crawler/realtime"
)

func main() {
	var (
		url             = flag.String("url", "", "GTFS Realtime feed URL")
		agencyID        = flag.String("agency", "", "Agency ID (UUID)")
		intervalSeconds = flag.Int("interval", 15, "Interval seconds")
	)
	flag.Parse()

	if *url == "" {
		if lib.GetenvFromSecretfile("FEED_URL") != "" {
			*url = lib.GetenvFromSecretfile("FEED_URL")
		} else {
			panic("url is required")
		}
	}
	if *agencyID == "" {
		if lib.GetenvFromSecretfile("AGENCY_ID") != "" {
			*agencyID = lib.GetenvFromSecretfile("AGENCY_ID")
		} else {
			panic("agency is required")
		}
	}

	schemaName := strings.ReplaceAll(*agencyID, "-", "")

	var lastChangedAt *time.Time

	for {
		go func() {
			message, err := realtime.GetMessage(*url)
			if err != nil {
				log.Println("realtime.GetMessage:", err)
			}

			timestamp := unixTime(message.Header.Timestamp)
			if lastChangedAt == nil || (timestamp != nil && *lastChangedAt != *timestamp) {
				lastChangedAt = timestamp
				realtime.RecordVehiclePositions(message, "a"+schemaName)
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
