# Build Stage
FROM golang:{{ .GoVersion }}-alpine AS builder

RUN apk update && apk add --no-cache git upx

WORKDIR /build

COPY . .

RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-w -s" -o ./bin/api ./cmd/api/main.go

RUN upx --best --lzma ./bin/api

# Run Stage
FROM alpine:latest

WORKDIR /app

RUN apk update && apk add ca-certificates && rm -rf /var/cache/apk*

COPY --from=builder /build/bin/api ./api

COPY .env /app

EXPOSE 3000

Cmd [ "/app/api" ]