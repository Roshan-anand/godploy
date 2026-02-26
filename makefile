.PHONY: start reset restart build

install-web:
	cd frontend && bun install
	
install-server:
	go mod tidy

install: install-web install-server

check:
	cd frontend && bun check:all
	
build-web:
	cd frontend && \
	bun install && bun run build

build-bin:
	go mod tidy && \
	go build -o ./bin/godploy cmd/main.go

build: build-web build-bin

generate:
	sqlc generate
	
start: generate build
	@./bin/godploy

reset:
	rm -rf ./data/*

restart: reset start

test:
	go test -v ./...

service-up:
	docker compose -f ./dynamic/compose.yaml up -d

service-down:
	docker compose -f ./dynamic/compose.yaml down
