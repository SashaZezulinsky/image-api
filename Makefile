local:
	docker build -t image-api .
	docker run -d --network host --name image-api-mongo mongo:latest
	docker run -d --network host --name image-api image-api image-api
clean:
	- docker rm -f image-api image-api-mongo
test:
	go test -v ./...
build:
	go build -o bin/image-api cmd/image-api/main.go
run:
	go run cmd/image-api/main.go
