DOCKER_IMAGE_NAME := bob

###################
# Docker          #
###################
.PHONY: docker-build
docker-build:
	docker build -t ${DOCKER_IMAGE_NAME} .

.PHONY: docker-run
docker-run:
	docker run -p 3000:3000 ${DOCKER_IMAGE_NAME}

###################
# App             #
###################
.PHONY: run
run:
	go run cmd/api/main.go

.PHONY: build
build:
	CGO_ENABLED=0 GOOS=linux go build -o ./bin/api ./cmd/api/main.go