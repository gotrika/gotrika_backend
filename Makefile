.PHONY:
.SILENT:
.DEFAULT_GOAL := run.http

run.http:
	go run cmd/app/main.go

swag:
	swag init -g internal/bootstrap/app.go