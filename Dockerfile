FROM golang:1.23.1-alpine AS builder

RUN apk update && apk add --no-cache git

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download


COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o app ./cmd/app

FROM alpine:latest

WORKDIR /root/

COPY --from=builder /app/app .
COPY --from=builder /app/configs/ ./configs/
COPY --from=builder /app/migrations ./migrations

RUN apk add --no-cache curl bash postgresql-client

RUN curl -L -o /usr/local/bin/goose https://github.com/pressly/goose/releases/download/v3.14.0/goose_linux_arm64 && \
    chmod +x /usr/local/bin/goose

EXPOSE 8080

CMD ["./app"]