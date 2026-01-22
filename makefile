start:
	@clear && \
	go build -o ./tmp/godploy ./cmd/main.go
	./tmp/godploy
