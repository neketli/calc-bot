.PHONY: build
build:
	go build -v ./cmd/calc-bot

.PHONY: install
install:
	go mod download

.PHONY: docker-build
docker:
	docker build -t "calc-bot" .

.PHONY: docker-run
docker:
	docker run --env-file .env calc-bot

.DEFAULT_GOAL := build