.PHONY: build
build:
	go build -v ./cmd/calc-bot

.DEFAULT_GOAL := build