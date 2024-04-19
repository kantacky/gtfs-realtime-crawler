package main

import (
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

	log.Println("Start fetching data from", url)

	data, err := realtime.FetchData(url)
	if err != nil {
		log.Fatal(err)
	}

	message, err := realtime.Deserialize(data)
	if err != nil {
		log.Fatal(err)
	}

	lastChangedAt := time.Unix(int64(*message.Header.Timestamp), 0)

	for {
		data, err := realtime.FetchData(url)
		if err != nil {
			log.Fatal(err)
		}

		message, err := realtime.Deserialize(data)
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
