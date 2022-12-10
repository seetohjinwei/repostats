.PHONY: build cli web

build:
	go build -o bin/repostats .

cli:
	go run . --mode cli

web:
	go run . --mode web