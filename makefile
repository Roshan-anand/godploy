.PHONY: start reset restart build

install-web:
	cd frontend && bun install

install-server:
	cd backend && go mod tidy

install: install-web install-server

dev-web:
	cd frontend && bun dev

dev-server:
	cd backend && air
	
check:
	cd frontend && bun check:all

build-web:
	cd frontend && \
	bun install && bun run build

build-bin:
	cd backend && \
	go mod tidy && \
	go build -o ../bin/godploy cmd/main.go

build: build-web build-bin

generate:
	cd backend && \
	sqlc generate

start: generate build
	@./bin/godploy

reset:
	rm -rf ./backend/data/*

restart: reset start

test:
	cd backend && \
	go test -v ./...

service-up:
	docker compose -f ./dynamic/compose.yaml up -d

service-down:
	docker compose -f ./dynamic/compose.yaml down
