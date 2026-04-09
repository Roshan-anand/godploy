.PHONY: start reset restart build install-web install-server install check build-web build-bin generate test img-build setup dev services-rm web-logs server-logs traefik-logs cloud-tunnel clean clean-web clean-server clean-cache clean-all

install-web:
	cd frontend && \
	bun install && \
	bun run prepare

install-server:
	cd backend && go mod tidy

install: install-web install-server

check:
	cd frontend && bun check

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

img-build: install build-web
	docker compose -f ./docker/compose.dev.yaml build

setup:
	@cd backend && \
	go run cmd/setup/main.go setup

dev:
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

clean-web:
	rm -rf ./frontend/node_modules ./frontend/.svelte-kit ./frontend/build

clean-server:
	rm -rf ./backend/bin ./backend/frontend/dist ./bin/godploy

clean: clean-web clean-server
