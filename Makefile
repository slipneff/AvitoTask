build:
	docker-compose build avitotask

run:
	docker-compose up avitotask

test:
	go test -v ./...

migrate-up:
	migrate -path ./schema -database 'postgres://postgres:admin@0.0.0.0:5432/avitodb?sslmode=disable' up

migrate-down:
	migrate -path ./schema -database 'postgres://postgres:admin@0.0.0.0:5432/avitodb?sslmode=disable' down


swag:
	swag init -g cmd/app/main.go