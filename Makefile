run:
	go run cmd/tustel/main.go

migrate-up:
	go run cmd/migrate/main.go -run=up

migrate-down:
	go run cmd/migrate/main.go -run=down
