start:
	@clear && \
	cd frontend && bun run build && \
    cd .. && go mod tidy && \
	go build -o ./bin/godploy cmd/main.go && \
    ./bin/godploy

build:
	@clear && \
	cd frontend && bun run build && \
	cd .. && go build -o ./bin/godploy cmd/main.go

install:
	@clear && \
	cd frontend && bun install && \
	cd .. && go mod tidy