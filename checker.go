package main

import (
	"log"
	"time"

	"github.com/kantacky/gtfs-realtime-crawler/realtime"
)

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
