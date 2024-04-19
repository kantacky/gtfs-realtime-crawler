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

	timestamp := time.Unix(int64(message.Header.GetTimestamp()), 0)
	timestampString := timestamp.Format("2006-01-02T15:04:05+09:00")
	fmt.Printf("Timestamp: %s  \n", timestampString)
	fmt.Printf("Incrementality: %s  \n", message.Header.GetIncrementality())
	fmt.Printf("Number of entities: %d  \n", len(message.Entity))

	fmt.Println("| ID | Timestamp | Vehicle ID | Position |")
	fmt.Println("| :-- | :-- | :-- | :-- |")

	for _, entity := range message.Entity {
		fmt.Printf(
			"| %s | %s | %s | %f, %f |\n",
			entity.GetId(),
			time.Unix(int64(entity.Vehicle.GetTimestamp()), 0).Format("2006-01-02T15:04:05+09:00"),
			entity.Vehicle.Vehicle.GetId(),
			entity.Vehicle.Position.GetLatitude(),
			entity.Vehicle.Position.GetLongitude(),
		)
	}
}
