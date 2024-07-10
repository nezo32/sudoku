include .env

run:
	go run .

migrate_create:
	migrate create -ext=sql -dir=internal/migrations -seq init

migrate_up:
	migrate -path internal/migrations -database "postgresql://${POSTGRES_USER}:${POSTGRES_PASSWORD}@${POSTGRES_HOST}:${POSTGRES_PORT}/${POSTGRES_DB}?sslmode=disable" -verbose up ${VERSION}

migrate_down:
	migrate -path internal/migrations -database "postgresql://${POSTGRES_USER}:${POSTGRES_PASSWORD}@${POSTGRES_HOST}:${POSTGRES_PORT}/${POSTGRES_DB}?sslmode=disable" -verbose down ${VERSION}

migrate_force:
	migrate -path internal/migrations -database "postgresql://${POSTGRES_USER}:${POSTGRES_PASSWORD}@${POSTGRES_HOST}:${POSTGRES_PORT}/${POSTGRES_DB}?sslmode=disable" -verbose force ${VERSION}

migration_version:
	migrate -path internal/migrations -database "postgresql://${POSTGRES_USER}:${POSTGRES_PASSWORD}@${POSTGRES_HOST}:${POSTGRES_PORT}/${POSTGRES_DB}?sslmode=disable" -verbose version

generate_postgres:
	cd ./generator && go run .

generate_proto:
	protoc --proto_path=. --go_out=./generated --go-grpc_out=./generated ./internal/proto/*

.PHONY: create_migration migrate_up migrate_down