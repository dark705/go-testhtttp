.PHONY: build
build:
	go build -o ./bin/app ./cmd/main.go
	chmod +x ./bin/app

lint:
	docker run --rm -e GOFLAGS='-buildvcs=false' -v $(shell pwd):/app -w /app 'golangci/golangci-lint:v1.59.1' sh -c \
    '$(git_config)\
     && golangci-lint run -v'

build_docker:
	buildx build . --tag dark705/go-testhttp:1.1 --tag dark705/go-testhttp:latest