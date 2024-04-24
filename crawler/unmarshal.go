package crawler

import "google.golang.org/protobuf/proto"

func Unmarshal(data []byte) (*FeedMessage, error) {
	message := &FeedMessage{}
	err := proto.Unmarshal(data, message)
	if err != nil {
		return nil, err
	}
	return message, nil
}
