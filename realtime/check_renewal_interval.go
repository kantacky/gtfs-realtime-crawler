package realtime

import (
	"log"
	"time"
)

func CheckRenewalInterval(url string) {
	log.Println("Start fetching data from", url)

	message, err := GetMessage(url)
	if err != nil {
		log.Println("GetMessage: ", err)
	}

	lastChangedAt := time.Unix(int64(*message.Header.Timestamp), 0)

	for {
		message, err := GetMessage(url)
		if err != nil {
			log.Println("GetMessage: ", err)
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
