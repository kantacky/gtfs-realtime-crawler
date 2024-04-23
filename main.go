package main

import (
	"log"
	"os"
	"strings"

	"github.com/joho/godotenv"
	"github.com/kantacky/gtfs-realtime-crawler/realtime"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("godotenv.Load error:", err)
	}

	if len(os.Args) < 3 {
		log.Fatal("URL and AgencyID is not found. Input URL and AgencyID as arguments.")
	}
	url := os.Args[1]
	agencyID := os.Args[2]
	schemaName := strings.ReplaceAll(agencyID, "-", "")
	realtime.RecordVehiclePositions(url, "a"+schemaName)
}
