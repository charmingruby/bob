# STEP 1: Build
FROM golang:{{ .GoVersion }}-alpine AS builder
RUN apk update && apk add --no-cache git upx
WORKDIR /build
COPY . .
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux GOARCH="arm64" go build -ldflags="-w -s" -o ./bin/api ./cmd/api/main.go

# STEP 2: Optimize
RUN upx --best --lzma ./bin/api

# STEP 3: Run
FROM alpine:latest
WORKDIR /app
RUN apk update && apk add ca-certificates && rm -rf /var/cache/apk*
COPY --from=builder /build/bin/api ./api
CMD [ "/app/api" ]