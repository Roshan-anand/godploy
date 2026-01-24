start:
	@clear && \
	cd frontend && bun run build && \
    cd .. && go build -o godploy cmd/main.go && \
    ./godploy
