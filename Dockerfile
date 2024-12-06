FROM golang:1.22 AS builder
WORKDIR /app
COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o ./bin/app ./cmd/main.go

FROM alpine:latest
RUN apk --no-cache add ca-certificates && addgroup -S app && adduser -S app -G app
USER app:app
WORKDIR /app
COPY --from=builder /app/bin ./bin

CMD ["./bin/app"]
EXPOSE 8000/tcp 9000/tcp