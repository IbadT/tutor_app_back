DB_DSN := "postgres://postgres:postgres@localhost:5433/tutor_app_back?sslmode=disable"
MIGRATE := migrate -path ./migrations -database $(DB_DSN)

migrate-new:
	migrate create -ext sql -dir ./migrations $(name)

migrate-up:
	$(MIGRATE) up

migrate-down:
	$(MIGRATE) down

gen:
	oapi-codegen -config openapi/.openapi -include-tags auth -package auth openapi/openapi.yaml > ./internal/web/auth/api.gen.go
	oapi-codegen -config openapi/.openapi -include-tags users -package users openapi/openapi.yaml > ./internal/web/users/api.gen.go
	oapi-codegen -config openapi/.openapi -include-tags courses -package courses openapi/openapi.yaml > ./internal/web/courses/api.gen.go
	oapi-codegen -config openapi/.openapi -include-tags lessons -package lessons openapi/openapi.yaml > ./internal/web/lessons/api.gen.go

lint:
	golangci-lint run --color=always

run:
	@docker compose up --build
	@go run cmd/main.go