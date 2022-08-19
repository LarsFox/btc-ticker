include .env
export $(shell sed 's/=.*//' .env)

default:
	@go run cmd/*.go

test:
	@go test ./...

lint:
	@golangci-lint run
