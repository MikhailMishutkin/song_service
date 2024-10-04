.PHONY: build
db:
	docker-compose up -d --build

migrateup:
	migrate -path migrations -database "postgres://root:root@localhost:5444/song_service?sslmode=disable" -verbose up

migratedown:
	migrate -path migrations -database "postgres://root:root@localhost:5444/song_service?sslmode=disable" -verbose down -1

http:
	go run ./cmd/service/main.go

base:
	go run ./cmd/base/main.go

build: db migrateup http base


.DEFAULT_GOAL := build