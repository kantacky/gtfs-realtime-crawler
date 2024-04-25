package crawler

import (
	"github.com/kantacky/apis-go/transit_realtime"
	"google.golang.org/protobuf/proto"
)

func Unmarshal(data []byte) (*transit_realtime.FeedMessage, error) {
	message := &transit_realtime.FeedMessage{}
	err := proto.Unmarshal(data, message)
	if err != nil {
		return nil, err
	}
	return message, nil
}
