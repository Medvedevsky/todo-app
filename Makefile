.PHONY: build
build:
	go build -v ./cmd

run: build
	./cmd

up: 
	migrate -database "postgres://postgres:andrey5522@localhost/restapi_dev?sslmode=disable" -path migrations up

down: 
	migrate -database "postgres://postgres:andrey5522@localhost/restapi_dev?sslmode=disable" -path migrations down

.DEFAULT_GOAL := run 
