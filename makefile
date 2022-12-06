.PHONY: build cli

build:
	go build -o bin/repostats .

cli:
	go run . --mode cli