all: build lint

build:
	@go build -o bin/repair_queue cmd/main.go

run: build
	@./bin/repair_queue

lint:
	@golangci-lint run
	
lint-fix:
	@golangci-lint run --fix

migration:
	@migrate create -ext sql -dir cmd/migrate/migrations $(filter-out $@,$(MAKECMDGOALS))

migrate-up:
	@go run cmd/migrate/main.go up

migrate-down:
	@go run cmd/migrate/main.go down