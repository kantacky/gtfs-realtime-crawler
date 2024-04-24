package crawler

func GetMessage(url string) (*FeedMessage, error) {
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
