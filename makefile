.PHONY: start reset restart build

install-web:
	cd frontend && bun install

install-server:
	cd backend && go mod tidy

install: install-web install-server

check:
	cd frontend && bun check:all

build-web:
	cd frontend && \
	bun install && bun run build

build-bin:
	cd backend && \
	go mod tidy && \
	go build -o ../bin/godploy cmd/server/main.go

build: build-web build-bin

generate:
	cd backend && \
	sqlc generate

start: generate build
	@./bin/godploy

reset:
	rm -rf ./backend/data/* ./data/*

restart: reset start

test:
	cd backend && \
	go test -v ./...

img-build:
	docker compose -f ./docker/compose.dev.yaml build

setup:install
	@cd backend && \
	go run cmd/setup/main.go setup

dev:install
	@cd backend && \
	go run cmd/setup/main.go dev

services-rm:
	docker service rm godploy_traefik godploy_web godploy_server

web-logs:
	docker service logs -f godploy_web

server-logs:
	docker service logs -f godploy_server

traefik-logs:
	docker service logs -f godploy_traefik

cloud-tunnel:
	docker run --rm -it \
        --network host \
        cloudflare/cloudflared:latest \
        tunnel --no-autoupdate --url http://localhost:8080
