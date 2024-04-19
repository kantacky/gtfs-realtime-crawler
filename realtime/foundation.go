package realtime

import (
	"io"
	"net/http"

	"google.golang.org/protobuf/proto"
)

func GetMessage(url string) (*FeedMessage, error) {
	data, err := fetchData(url)
	if err != nil {
		return nil, err
	}

	message, err := unmarshal(data)
	if err != nil {
		return nil, err
	}

	return message, nil
}

func fetchData(url string) ([]byte, error) {
	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	res, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

func unmarshal(data []byte) (*FeedMessage, error) {
	message := &FeedMessage{}
	err := proto.Unmarshal(data, message)
	if err != nil {
		return nil, err
	}
	return message, nil
}
