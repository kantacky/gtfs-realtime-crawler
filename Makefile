.PHONY: install build-for-docker-hub push-to-docker-hub

install:
	go get -u
	go mod tidy

build-for-docker-hub:
	docker build -t 'kantacky/gtfs-realtime-crawler:latest' .

push-to-docker-hub:
	docker push 'kantacky/gtfs-realtime-crawler'
