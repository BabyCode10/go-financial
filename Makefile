run:
	go run main.go

migrate-create:
	docker compose --profile tools run --rm migrate create -ext sql -dir /migrations $(name)

migrate-up:  
	docker compose --profile tools run --rm migrate up

migrate-down:  
	docker compose --profile tools run --rm migrate down
