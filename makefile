.PHONY: start reset restart build
	
install:
	cd frontend && bun install && \
	cd .. && go mod tidy


build-web:
	cd frontend && bun run build

build-bin:
	go mod tidy && \
	go build -o ./bin/godploy cmd/main.go

build:
	$(MAKE) build-web && \
	cd .. && go build -o ./bin/godploy cmd/main.go

start: 
	@sqlc generate && \
    $(MAKE) build && \
    ./bin/godploy

reset:
	rm -rf ./data/*

restart: reset start

test:
	go test -v ./...

service-up:
	docker compose -f ./dynamic/compose.yaml up -d

service-down:
	docker compose -f ./dynamic/compose.yaml down
