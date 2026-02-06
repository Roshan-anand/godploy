install:
	@clear && \
	cd frontend && bun install && \
	cd .. && go mod tidy

build:
	@clear && \
	cd frontend && bun run build && \
	cd .. && go build -o ./bin/godploy cmd/main.go

test:
	@clear && go test -v ./testing/...

start:
	@clear && \
	cd frontend && bun run build && \
    cd .. && \
    sqlc generate && \
	go mod tidy && \
	go build -o ./bin/godploy cmd/main.go && \
    ./bin/godploy

reset:
	@clear && \
	rm -rf ./test.db && \
	go build -o ./bin/godploy cmd/main.go

generate:
	sqlc generate
