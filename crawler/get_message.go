package crawler

import "github.com/kantacky/apis-go/transit_realtime"

func GetMessage(url string) (*transit_realtime.FeedMessage, error) {
	data, err := FetchData(url)
	if err != nil {
		return nil, err
	}

	message, err := Unmarshal(data)
	if err != nil {
		return nil, err
	}

	return message, nil
}
