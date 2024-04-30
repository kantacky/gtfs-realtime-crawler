.PHONY: install build-and-push-for-docker-hub

install:
	go get -u
	go mod tidy

build-and-push-for-docker-hub:
	docker buildx build --platform linux/arm64 -t 'kantacky/gtfs-realtime-crawler:arm64' . --no-cache
	docker push 'kantacky/gtfs-realtime-crawler:arm64'
	docker buildx build --platform linux/amd64 -t 'kantacky/gtfs-realtime-crawler:amd64' . --no-cache
	docker push 'kantacky/gtfs-realtime-crawler:amd64'
	docker manifest create 'kantacky/gtfs-realtime-crawler:latest' -a \
		'kantacky/gtfs-realtime-crawler:arm64' \
		'kantacky/gtfs-realtime-crawler:amd64'
	docker manifest push 'kantacky/gtfs-realtime-crawler:latest'
