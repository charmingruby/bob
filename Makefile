build:
		go build -o bob cmd/cli/main.go

start: build
		mkdir -p ../dummy
		mv bob ../dummy/
