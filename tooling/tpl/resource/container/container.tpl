# Build Stage
FROM golang:{{ .GoVersion }}-alpine AS builder
WORKDIR /build
COPY . .
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -o ./bin/api ./cmd/api/main.go

# Run Stage
FROM alpine:latest
WORKDIR /app
RUN apk update && apk add ca-certificates && rm -rf /var/cache/apk*
COPY --from=builder /build/bin/api ./api
COPY .env /app
EXPOSE 3000
Cmd [ "/app/api" ]