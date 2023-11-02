.SILENT:

# default: build

build:
	go mod tidy -v
	go mod verify
	go test ./... -race
	go build -o .bin/app ./cmd/app

run: build
	./.bin/app

stop:
	docker compose down -v

up: stop
	docker compose up --build
