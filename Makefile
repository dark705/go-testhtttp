.PHONY: build
build:
	go build -o ./bin/app ./cmd/main.go
	chmod +x ./bin/app
